package main

import (
	"fmt"
	"github.com/mhope-2/greenlight/internal/data"
	"github.com/mhope-2/greenlight/internal/validator"
	"net/http"
	"time"
)



// Add a createMovieHandler for the "POST /v1/movies" endpoint.
func (app *Application) createMovieHandler(w http.ResponseWriter, r *http.Request) {

	// Declare an anonymous struct to hold the information that we expect
	var input struct {
		Title string 				`json:"title"`
		Year int32					`json:"year"`
		Runtime int32				`json:"runtime"`
		Genres []string				`json:"genres"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		// Use the new badRequestResponse() helper.
		app.badRequestResponse(w, r, err)
		return
	}


	// Copy the values from the input struct to a new Movie struct.
	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: data.Runtime(input.Runtime),
		Genres:  input.Genres,
	}

	// Initialize a new Validator.
	v := validator.New()

	// Call the ValidateMovie() function and return a response containing the errors if
	// any of the checks fail.

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
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












