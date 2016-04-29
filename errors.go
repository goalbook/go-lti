package lti

import (
	"errors"
)

const (
	errLogBase = "LTI Provider Error Log: "
	logBase    = "LTI Provider Log: "
)

/*
	Unexported errors
*/
var (
	errBadContentType = errors.New("Unsupported content type")

	errInvalidRequest = errors.New("We could not validate your request to use our tool.")

	// Signature Validation errors
	errLogInvalidSignature = errors.New(errLogBase + "Invalid OAuth Signature Provided.")

	// Timestamp Valdiation errors
	errLogInvalidTimestamp = errors.New(errLogBase + "OAuth Timestamp is outside the acceptable range.")

	// Nonce Validation errors
	errLogUsedNonce = errors.New(errLogBase + "OAuth Nonce has already been used.")
)

/*
	Exported errors
*/
var (
	// Resource Link ID errors
	ErrLogUnknownResourceLinkId = errors.New(errLogBase + "Provided Resource Link ID does not have an associated resource.")
)
