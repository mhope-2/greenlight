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


	// Initialize a new Validator instance.
	v := validator.New()

	// Use the Check() method to execute our validation checks. This will add the
	// provided key and error message to the errors map if the check does not evaluate to true.
	v.Check(input.Title != "", "title", "must be provided")
	v.Check(len(input.Title) <= 500, "title", "must not be more than 500 bytes long")

	v.Check(input.Year != 0, "year", "must be provided")
	v.Check(input.Year >= 1888, "year", "must be greater than 1888")
	v.Check(input.Year <= int32(time.Now().Year()), "year", "must not be in the future")

	v.Check(input.Runtime != 0, "runtime", "must be provided")
	v.Check(input.Runtime > 0, "runtime", "must be a positive integer")
	v.Check(input.Genres != nil, "genres", "must be provided")
	v.Check(len(input.Genres) >= 1, "genres", "must contain at least 1 genre")
	v.Check(len(input.Genres) <= 5, "genres", "must not contain more than 5 genres")

	// Note that we're using the Unique helper in the line below to check that all
	// values in the input.Genres slice are unique.
	v.Check(validator.Unique(input.Genres), "genres", "must not contain duplicate values")
	// Use the Valid() method to see if any of the checks failed. If they did, then use
	//the failedValidationResponse() helper to send a response to the client, passing // in the v.Errors map.
	if !v.Valid() {
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












