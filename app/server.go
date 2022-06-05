package main

import (
	"encoding/json"
	"net/http"
)

func PostServer(w http.ResponseWriter, r *http.Request) {
	postsJson, err := json.Marshal(PostLists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(postsJson)
}
