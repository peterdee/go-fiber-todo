package main

type Todo struct {
	Completed bool   `json:"completed"`
	Id        string `json:"id"`
	Text      string `json:"text"`
}
