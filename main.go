package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello"{
	http.Error(w, "404 not found", http.StatusNotFound)
	return
}
if r.Method != "GET"{
	http.Error(w,"Method is not supported", http.StatusNotFound)
	return
}
fmt.Fprintf(w,"Hello you in the hello Server")
}

func formHandler(w http.ResponseWriter, r *http.Request){
if err := r.ParseForm();err !=nil{
	fmt.Fprintf(w,"ParseForm() err: %v", err)
	return
}
fmt.Fprintf(w,"POST request Successfull \n")
name := r.FormValue("name")
email := r.FormValue("email")
message := r.FormValue("message")
fmt.Fprintf(w, "Name %s\n", name)
fmt.Fprintf(w, "Email %s\n", email)
fmt.Fprintf(w, "Message %s\n", message)
}


func main() {
	fileServer := http.FileServer(http.Dir("./static")) 
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/forms", formHandler)

	fmt.Printf("Server starting at port 8080 \n")
	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}