package data

import "time"

type Task struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"-"`
	Expires time.Time `json:"-"`
}