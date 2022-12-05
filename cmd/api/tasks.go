package main

import (
	"fmt"
	"net/http"

	"github.com/OneGeag/task-manager/internal/data"
)

func (app *application) createTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new task")
}

func (app *application) showTaskHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	task := data.Task{
		ID: id,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"task": task}, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server have got some issue", http.StatusInternalServerError)
	}
}