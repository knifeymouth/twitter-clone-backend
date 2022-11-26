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

func GetTweet(w http.ResponseWriter, requestId int64) {
	var tweet structs.Tweet

	for _, value := range mocks.MockTweets {
		if value.Id == requestId {
			tweet = value
			break
		}
	}

	if tweet == (structs.Tweet{}) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(tweet)
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
