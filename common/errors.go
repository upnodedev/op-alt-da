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

// ErrInsufficientBalance is returned when the wallet does not have enough balance to post the transaction.
var ErrInsufficientBalance = errors.New("insufficient balance")

// ErrWalletNotFound is returned when the wallet is not found.
var ErrWalletNotFound = errors.New("wallet not found")

// ErrDataNotFound is returned when the data mapping on-chain is not found.
var ErrDataNotFound = errors.New("data not found")
