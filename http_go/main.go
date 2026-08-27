package main

import (
	"fmt"
	"net/http"
)

type user struct{
	id int32
	name string
	date_joined string
}

var users []user
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",handleRoot)
	mux.HandleFunc("POST /user",createUser)
	fmt.Println("Server is running on port: 8000")
	http.ListenAndServe(":8000",mux)
}

func handleRoot(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, World!")
}

func createUser(w http.ResponseWriter, r *http.Request){
	var u user
	fmt.Scan(&u.id, &u.name, &u.date_joined)
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "User created successfully")
}

//Handler Function Definition : 
// A handler function is a function that takes two parameters: an http.ResponseWriter and an *http.Request. The http.ResponseWriter is used to send a response back to the client, while the *http.Request contains information about the incoming request, such as the URL, headers, and body.
// The reason http.Request is a pointer is to avoid copying the entire request struct, which can be large and inefficient. By passing a pointer, we can access the request data directly without creating a new copy of it. This improves performance and reduces memory usage, especially for large requests.
