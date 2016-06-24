package lti

import (
	"github.com/goalbook/goalbook-auth/auth/lti/Godeps/_workspace/src/github.com/google/go-querystring/query"
)

// LTIStdResponse provides ability to serialize LTI response messages per LTI 1.0 spec.
type LTIStdResponse struct {
	LTIMessage      string `url:"lti_msg,omitempty"`
	LTIErrorMessage string `url:"lti_errormsg,omitempty"`
	LTILog          string `url:"lti_log,omitempty"`
	LTIErrorLog     string `url:"lti_errorlog,omitempty"`

	StatusCode int `url:"-"`
}

// Create a new LTI response with provided user error message, error log message and error status code.
func NewLTIErrorResponse(errMsg, logMsg string, statusCode int) *LTIStdResponse {
	return &LTIStdResponse{
		LTIErrorMessage: errMsg,
		LTIErrorLog:     logMsg,

		StatusCode: statusCode,
	}
}

// Create a new LTI response with provided user message, log message and status code.
func NewLTIStdResponse(respMsg, logMsg string, statusCode int) *LTIStdResponse {
	return &LTIStdResponse{
		LTIMessage: respMsg,
		LTILog:     logMsg,

		StatusCode: statusCode,
	}
}

// Serialize LTI response fields into encoded URL query string.
func (l *LTIStdResponse) Serialize() (string, error) {
	vals, err := query.Values(l)
	if err != nil {
		return "", err
	}

	return vals.Encode(), nil
}
