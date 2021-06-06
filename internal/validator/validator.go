package validator

import "regexp"

// Declare a regular expression for sanity checking the format of email addresses
var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-z])")
	)

// Define a new Validator type which contains a map of validation errors.
type Validator struct {
	Errors map[string]string
}