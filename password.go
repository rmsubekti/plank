package plank

import (
	"errors"
	"regexp"
)

// Password is a type for password.
// # example
// 		plank.Password("password").Validate() // error
// 		plank.Password("5uperP@ssw0rd").Ok() // true
type Password string

var upper = regexp.MustCompile(`[[:upper:]]+`)
var lower = regexp.MustCompile(`[[:lower:]]+`)
var punct = regexp.MustCompile(`[[:punct:]]+`)
var digit = regexp.MustCompile(`[[:digit:]]+`)

// Validate checks if a given password is valid.
//
// It takes no parameters.
// It returns an error if the password does not meet the required criteria.
func (p Password) Validate() (err error) {
	if !upper.MatchString(string(p)) {
		err = errors.New("password should contain one or more uppercase character")
	}
	if !lower.MatchString(string(p)) {
		err = errors.New("password should contain one or more lowercase character")
	}
	if !punct.MatchString(string(p)) {
		err = errors.New("password should contain one or more special character")
	}
	if !digit.MatchString(string(p)) {
		err = errors.New("password should contain one or more digit character")
	}
	return
}

// Ok checks if the password meets the requirements.
//
// It checks if the password contains at least one uppercase letter,
// one lowercase letter, one punctuation character, and one digit.
// Returns true if the password meets the requirements, false otherwise.
func (p Password) Ok() bool {
	if !upper.MatchString(string(p)) {
		return false
	}
	if !lower.MatchString(string(p)) {
		return false
	}
	if !punct.MatchString(string(p)) {
		return false
	}
	return digit.MatchString(string(p))
}
