package main

import (
	"fmt"
	"github.com/mhope-2/greenlight/internal/data"
	"net/http"
	"time"
)

// Add a createMovieHandler for the "POST /v1/movies" endpoint.
func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

// createMovies handler
func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request){

	// get id param and convert into base 10 int
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID: id,
		CreatedAt: time.Now(),
		Title: "Cassablanca",
		Runtime: 102,
		Genres: []string{"drama", "romance", "war"},
		Version: 1,
	}

	// Encode the struct to JSON and send it as the HTTP response.
	err = app.writeJSON(w, http.StatusOK,  envelope{"movie": movie}, nil)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}












