package main

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *Application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {

	// Create a map which holds the information that we want to send in the response.
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}

	// Pass the map to the json.Marshal() function.
	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil{
		// Use the serverErrorResponse() helper.
		app.serverErrorResponse(w, r, err)
		return
	}

}
