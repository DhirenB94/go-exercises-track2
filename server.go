package main

import (
	"encoding/json"
	"net/http"
)

type PostServer struct {
	http.Handler
}

func NewPostServer() *PostServer {
	p := new(PostServer)
	router := http.NewServeMux()
	router.Handle("/posts", http.HandlerFunc(p.postsHandler))
	p.Handler = router
	return p
}

func (p *PostServer) postsHandler(w http.ResponseWriter, r *http.Request) {
	postsJson, err := json.Marshal(PostLists)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(postsJson)
}
