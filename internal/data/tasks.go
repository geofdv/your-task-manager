package data

import "time"

type Task struct {
	TaskID int64 `json:"taskid"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	Status string `json:"status"`
	Progress Progress `json:"progress"`
	CreatedAt time.Time `json:"-"`
	Expires time.Time `json:"-"`
	Version int32 `json:"version"`
}
