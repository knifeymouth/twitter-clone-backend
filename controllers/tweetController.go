package controllers

import (
	"net/http"
	"twitter-clone/services"
)

func Tweet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		services.GetAllTweets(w)

	case "POST":
		services.CreateTweet(w, r)
	}

}
