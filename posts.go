package main

import "time"

var PostLists []Posts

type Posts struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	PublishedAt time.Time `json:"publishedAt"`
	Author      Author    `json:"author"`
}

type Author struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
