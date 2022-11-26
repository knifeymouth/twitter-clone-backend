package main

import (
	"log"
	"net/http"
	"twitter-clone/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Tweet)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
