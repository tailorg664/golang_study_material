package main

import (
	"errors"
	"fmt"
	"sync"
)

// Node represents an individual server instance with its own local memory.
type Node struct {
	ID   string
	Data map[string]string
	mu   sync.RWMutex
}

// Cluster manages our nodes and simulates network connectivity/partitions.
type Cluster struct {
	Nodes      map[string]*Node
	Mode       string                  // "CP" (Consistency-focused) or "AP" (Availability-focused)
	Partitions map[string]map[string]bool // map[nodeA][nodeB] = true means network is broken between them
	mu         sync.RWMutex
}

// NewCluster initialises a cluster with the given node IDs and mode.
func NewCluster(mode string, nodeIDs ...string) *Cluster {
	nodes := make(map[string]*Node)
	for _, id := range nodeIDs {
		nodes[id] = &Node{
			ID:   id,
			Data: make(map[string]string),
		}
	}
	return &Cluster{
		Nodes:      nodes,
		Mode:       mode,
		Partitions: make(map[string]map[string]bool),
	}
}

// SetPartition cuts off or restores communication between two nodes.
func (c *Cluster) SetPartition(nodeA, nodeB string, partitioned bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.Partitions[nodeA] == nil {
		c.Partitions[nodeA] = make(map[string]bool)
	}
	if c.Partitions[nodeB] == nil {
		c.Partitions[nodeB] = make(map[string]bool)
	}

	c.Partitions[nodeA][nodeB] = partitioned
	c.Partitions[nodeB][nodeA] = partitioned
}

// CanCommunicate checks if nodeA can successfully reach nodeB.
func (c *Cluster) CanCommunicate(nodeA, nodeB string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.Partitions[nodeA] == nil {
		return true
	}
	return !c.Partitions[nodeA][nodeB]
}

// GetReachableNodes returns a list of all nodes a specific node can talk to.
func (c *Cluster) GetReachableNodes(fromNode string) []string {
	var reachable []string
	for toNode := range c.Nodes {
		if c.CanCommunicate(fromNode, toNode) {
			reachable = append(reachable, toNode)
		}
	}
	return reachable
}

// Write attempts to save data via a "Coordinator" node.
func (c *Cluster) Write(coordinatorID string, key, value string) error {
	reachable := c.GetReachableNodes(coordinatorID)
	quorumSize := (len(c.Nodes) / 2) + 1

	if c.Mode == "CP" {
		// CP Mode: We must reach a majority quorum to guarantee Consistency.
		if len(reachable) < quorumSize {
			return fmt.Errorf("CP Write Failed on %s: Node cannot reach majority quorum (only %d/%d reachable)", 
				coordinatorID, len(reachable), len(c.Nodes))
		}

		// Write to all reachable nodes (which is at least a majority)
		for _, nodeID := range reachable {
			node := c.Nodes[nodeID]
			node.mu.Lock()
			node.Data[key] = value
			node.mu.Unlock()
		}
		return nil
	}

	// AP Mode: Stay Available at all costs. Accept the write locally and propagate best-effort.
	for _, nodeID := range reachable {
		node := c.Nodes[nodeID]
		node.mu.Lock()
		node.Data[key] = value
		node.mu.Unlock()
	}
	if len(reachable) < len(c.Nodes) {
		fmt.Printf("[AP Log] Write on %s succeeded but only partially replicated to %v due to partition\n", coordinatorID, reachable)
	}
	return nil
}

// Read attempts to retrieve data from a specific node.
func (c *Cluster) Read(nodeID string, key string) (string, error) {
	node, exists := c.Nodes[nodeID]
	if !exists {
		return "", errors.New("node not found")
	}

	reachable := c.GetReachableNodes(nodeID)
	quorumSize := (len(c.Nodes) / 2) + 1

	if c.Mode == "CP" {
		// CP Mode: Prevent reading stale data from a minority partition.
		if len(reachable) < quorumSize {
			return "", fmt.Errorf("CP Read Failed on %s: Isolated node cannot guarantee strong consistency", nodeID)
		}
	}

	node.mu.RLock()
	val, ok := node.Data[key]
	node.mu.RUnlock()

	if !ok {
		return "", errors.New("key not found")
	}
	return val, nil
}

func main() {
	fmt.Println("=== 1. RUNNING CP MODE (Consistency + Partition Tolerance) ===")
	cpCluster := NewCluster("CP", "Node-1", "Node-2", "Node-3")

	// Network is healthy. Write propagates everywhere.
	_ = cpCluster.Write("Node-1", "title", "Go Design Patterns")
	
	// Simulate Network Partition: Node-3 is completely isolated from Node-1 and Node-2.
	// Node-1 & Node-2 can still talk to each other (majority of 2/3).
	cpCluster.SetPartition("Node-1", "Node-3", true)
	cpCluster.SetPartition("Node-2", "Node-3", true)

	// Test a write on the majority partition (Node-1)
	err := cpCluster.Write("Node-1", "title", "Go SOLID & CAP")
	if err == nil {
		fmt.Println("Write to Node-1 (majority): Succeeded!")
	}

	// Test a write on the isolated minority partition (Node-3)
	err = cpCluster.Write("Node-3", "title", "Rogue Update")
	if err != nil {
		fmt.Printf("Write to Node-3 (minority): Failed as expected -> %v\n", err)
	}

	// Read from Node-3 (minority)
	_, err = cpCluster.Read("Node-3", "title")
	if err != nil {
		fmt.Printf("Read from Node-3 (minority): Failed as expected -> %v\n", err)
	}

	fmt.Println("\n=== 2. RUNNING AP MODE (Availability + Partition Tolerance) ===")
	apCluster := NewCluster("AP", "Node-1", "Node-2", "Node-3")

	// Write initial value
	_ = apCluster.Write("Node-1", "title", "Go Design Patterns")

	// Simulate same Network Partition: Node-3 isolated from Node-1 & Node-2.
	apCluster.SetPartition("Node-1", "Node-3", true)
	apCluster.SetPartition("Node-2", "Node-3", true)

	// Test write on the majority partition (Node-1)
	_ = apCluster.Write("Node-1", "title", "Go SOLID & CAP")

	// Test write on the isolated node (Node-3)
	// It accepts the write locally because it must remain available!
	_ = apCluster.Write("Node-3", "title", "Rogue Update")

	// Read from both partitions to see the data split (Inconsistency)
	val1, _ := apCluster.Read("Node-1", "title")
	val3, _ := apCluster.Read("Node-3", "title")
	fmt.Printf("Read from Node-1: %s\n", val1)
	fmt.Printf("Read from Node-3: %s (STALE/CONFLICTING DATA!)\n", val3)
}