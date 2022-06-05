package main

import "time"

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
