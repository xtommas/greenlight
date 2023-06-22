package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xtommas/greenlight/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		Id:        id,
		CreatedAt: time.Now(),
		Title:     "Blade Runner 2049",
		Runtime:   163,
		Genres:    []string{"action", "drama", "mystery", "sci-fi", "thriller"},
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
