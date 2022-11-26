package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"twitter-clone/mocks"
	"twitter-clone/structs"
)

func GetAllTweets(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(mocks.MockTweets)
}

func GetTweet(w http.ResponseWriter, requestId int64) {
	tweetIndex, err := findTweet(requestId, mocks.MockTweets)

	if err != nil {
		json.NewEncoder(w).Encode(structs.CommonResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	tweet := mocks.MockTweets[tweetIndex]
	json.NewEncoder(w).Encode(structs.CommonResponse{
		Code: 200,
		Data: tweet,
	})
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

func UpdateTweet(w http.ResponseWriter, r *http.Request) {
	// get and validate id
	idString := r.URL.Query().Get("id")
	if idString == "" {
		json.NewEncoder(w).Encode(structs.CommonResponse{
			Code:    404,
			Message: "id is required",
		})
		return
	}

	// convert id into integer and validate
	requestId, err := strconv.Atoi(idString)
	if err != nil {
		json.NewEncoder(w).Encode(structs.CommonResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	// find and validate tweet from tweet data
	tweetIndex, err := findTweet(int64(requestId), mocks.MockTweets)
	if err != nil {
		json.NewEncoder(w).Encode(structs.CommonResponse{
			Code:    404,
			Message: err.Error(),
		})
		return
	}

	// get tweet from found index and update
	tweet := mocks.MockTweets[tweetIndex]
	var incomingTweet structs.Tweet
	json.NewDecoder(r.Body).Decode(&incomingTweet)
	newTweet := structs.Tweet{
		Id:        tweet.Id,
		UserId:    tweet.UserId,
		Content:   incomingTweet.Content,
		Img:       incomingTweet.Img,
		CreatedAt: tweet.CreatedAt,
		UpdatedAt: time.Now(),
	}
	mocks.MockTweets[tweetIndex] = newTweet

	// return json
	json.NewEncoder(w).Encode(structs.CommonResponse{
		Code: 200,
		Data: newTweet,
	})
}

func findTweet(id int64, tweets []structs.Tweet) (int, error) {
	for index, tweet := range tweets {
		if tweet.Id == id {
			return index, nil
		}
	}

	return 0, errors.New("tweet not found")
}
