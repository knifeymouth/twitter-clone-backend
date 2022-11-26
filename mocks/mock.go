package mocks

import (
	"time"
	"twitter-clone/structs"
)

var MockTweets = []structs.Tweet{
	{
		Id:      1,
		UserId:  1,
		Content: "First tweet!",
		Img: structs.Image{
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
		Img: structs.Image{
			Src:    "2",
			Width:  100,
			Height: 100,
		},
		CreatedAt: time.Now(),
	},
}
