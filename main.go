package main

import (
	"fmt"
	"net/http"
)

// app.get('/', function(req, res) {
// 		res.send("Hello World")
// })

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// app.get('/json', function(req, res) {
// 		res.json({ message :"ok"})
// })

func returnJson(w http.ResponseWriter, r *http.Request) {

	jsonData := []byte(`{"message":"ok"}`)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func handleRequest() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/json", returnJson)
	http.HandleFunc("/users", GetUsers)
	http.ListenAndServe(":8080", nil)
}

func main() {
	fmt.Println("Server is starting")
	fmt.Println("Service is running at port 8080")
	handleRequest()
}
