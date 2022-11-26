package controllers

import (
	"encoding/json"
	"net/http"

	"twitter-clone/mocks"
)

func Tweet(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mocks.MockTweets)
}
