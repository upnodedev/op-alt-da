package common

import "errors"

// ErrNotFound is returned when a key is not found in the store.
var ErrNotFound = errors.New("not found")

// ErrInvalidInput is returned when the input is not valid for posting to the DA storage.
var ErrInvalidInput = errors.New("invalid input")

// ErrInvalidCommitment is returned when the commitment cannot be parsed into a known commitment type.
var ErrInvalidCommitment = errors.New("invalid commitment")

// ErrCommitmentMismatch is returned when the commitment does not match the given input.
var ErrCommitmentMismatch = errors.New("commitment mismatch")
