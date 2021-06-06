package main

import (
	"fmt"
	"github.com/mhope-2/greenlight/internal/data"
	"net/http"
	"time"
)



// Add a createMovieHandler for the "POST /v1/movies" endpoint.
func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	// Declare an anonymous struct to hold the information that we expect
	var input struct {
		Title string 		`json:"title"`
		Year int32			`json:"year"`
		Runtime int32		`json:"runtime"`
		Genres []string		`json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)

}



// createMovies handler
func (app *Application) showMovieHandler(w http.ResponseWriter, r *http.Request){

	// get id param and convert into base 10 int
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		// Use the notFoundResponse() helper.
		app.notFoundResponse(w, r)
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
		// Use the serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
	}

}












