package models

type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}