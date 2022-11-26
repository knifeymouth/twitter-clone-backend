package services

import (
	"encoding/json"
	"net/http"
	"time"

	"twitter-clone/mocks"
	"twitter-clone/structs"
)

func GetAllTweets(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(mocks.MockTweets)
}

func CreateTweet(w http.ResponseWriter, r *http.Request) {
	id := len(mocks.MockTweets)

	var t structs.Tweet
	json.NewDecoder(r.Body).Decode(&t)

	t.Id = int64(id)
	t.CreatedAt = time.Now()
	json.NewEncoder(w).Encode(t)

	mocks.MockTweets = append(mocks.MockTweets, t)
}
