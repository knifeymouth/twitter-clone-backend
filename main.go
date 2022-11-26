package main

import (
	"encoding/json"
	"log"
	"net/http"
	"twitter-clone/mocks"
)

func main() {
	http.HandleFunc("/", getAllTweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAllTweets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mocks.MockTweets)
}
