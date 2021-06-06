package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"strconv"

)

// Define an envelope type.
type envelope map[string]interface{}


// Retrieve the "id" URL parameter from the current request context, then convert it to
//an integer and return it. If the operation isn't successful, return 0 and an error.
//
func (app *Application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil
}


// writeJSON method
func (app *Application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {

	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Append a newline to make it easier to view in terminal applications.
	//js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}
	// Add the "Content-Type: application/json" header, then write the status code and // JSON response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}


// readJSON() method to handle json decode errors
func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	// Decode the request body into the target destination.
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		// If there is an error during decoding, start the triage...
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError

		switch {
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)

		default:
			return err
		}
	}
	return nil
}