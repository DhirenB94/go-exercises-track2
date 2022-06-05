package main

import (
	"encoding/json"
	"net/http"
)

type PostServer struct {
	router *http.ServeMux
}

func NewPostServer() *PostServer {
	p := &PostServer{router: http.NewServeMux()}
	p.router.Handle("/posts", http.HandlerFunc(p.postsHandler))

	return p
}

func (p *PostServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
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
