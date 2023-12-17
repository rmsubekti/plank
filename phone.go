package plank

import (
	"errors"
	"regexp"
)

// Phone is a type for phone numbers.
// # example
// 		plank.Phone("0123456789").Ok() // true
// 		plank.Phone("example.com").Validate() // not valid phone number
type Phone string

var rxPhone = regexp.MustCompile(`^(?:[[:digit:]]+)$`)

// Validate checks if the phone number is valid.
//
// It takes no parameters.
// It returns an error if the phone number is not valid.
func (p Phone) Validate() (err error) {
	if !rxPhone.MatchString(string(p)) {
		return errors.New("not valid phone number")
	}
	return
}

// Ok checks if the phone number is valid.
//
// It takes no parameters.
// It returns a boolean value indicating whether the phone number is valid or not.
func (p Phone) Ok() bool {
	return rxPhone.MatchString(string(p))
}
