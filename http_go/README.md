# HTTP in Golang
- The most basic way to create a http port is :
``` 
func main(){
	http.HandleFunc("/hello-world", func(w http.ResponseWriter,r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe(":8080",nil)
}
```
- Pointer at http.Request is used to directly access the value of request struct without creating copy and getting bulky data.
