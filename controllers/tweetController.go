package controllers

import (
	"net/http"
	"strconv"
	"twitter-clone/services"
)

func Tweet(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		idString := r.URL.Query().Get("id")

		if idString == "" {
			services.GetAllTweets(w)
			return
		}

		id, err := strconv.Atoi(idString)

		if err != nil {
			services.GetAllTweets(w)
		}

		services.GetTweet(w, int64(id))

	case "POST":
		services.CreateTweet(w, r)

	case "PUT":
		services.UpdateTweet(w, r)
	}

}
