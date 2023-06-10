package models

import "errors"

var (
	// ErrNoRecord is used when a record is not found in the database
	ErrNoRecord = errors.New("models: no matching record found")
	// ErrInvalidCredentials is used when a user tries to login with an incorrect email or password
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// ErrDuplicateEmail is used when a user tries to signup with an email address that's already in use
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
