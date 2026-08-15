package api

import (
	"net/http"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
}

type Server struct {
	*mux.Router
	shoppingItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router : mux.NewRouter(),
		shoppingItems: []Item{},
	}
	return s
}
func (s *Server) CreateShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		i.ID = uuid.New()
		s.shoppingItems = append(s.shoppingItems, i)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
	}
}
func (s *Server) GetShoppingItems() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItems); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
func (s *Server) GetShoppingItemByID() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}
		for _, item := range s.shoppingItems {
			if item.ID == id {
				w.Header().Set("Content-Type", "application/json")
				if err := json.NewEncoder(w).Encode(item); err != nil{
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}