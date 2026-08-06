package main

import "net/http"


func main(){
	http.HandleFunc("/hello-world", func(w http.ResponseWriter,r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080",nil)
}

//Handler Function Definition : 
// A handler function is a function that takes two parameters: an http.ResponseWriter and an *http.Request. The http.ResponseWriter is used to send a response back to the client, while the *http.Request contains information about the incoming request, such as the URL, headers, and body.
// The reason http.Request is a pointer is to avoid copying the entire request struct, which can be large and inefficient. By passing a pointer, we can access the request data directly without creating a new copy of it. This improves performance and reduces memory usage, especially for large requests.
