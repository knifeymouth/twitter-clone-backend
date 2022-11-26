package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type tweet struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userid"`
	Content   string    `json:"content"`
	Img       image     `json:"img"`
	CreatedAt time.Time `json:"createdAt"`
}

type image struct {
	Src    string `json:"src"`
	Width  int16  `json:"width"`
	Height int16  `json:"height"`
}

var mockTweets = []tweet{
	{
		Id:      1,
		UserId:  1,
		Content: "First tweet!",
		Img: image{
			Src:    "1",
			Width:  100,
			Height: 100,
		},
		CreatedAt: time.Now(),
	},
	{
		Id:        2,
		UserId:    2,
		Content:   "Second tweet!",
		CreatedAt: time.Now(),
	},
	{
		Id:      3,
		UserId:  1,
		Content: "Third tweet!",
		Img: image{
			Src:    "2",
			Width:  100,
			Height: 100,
		},
		CreatedAt: time.Now(),
	},
}

func main() {
	http.HandleFunc("/", getAllTweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getAllTweets(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(mockTweets)
}
