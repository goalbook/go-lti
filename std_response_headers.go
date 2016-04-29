package lti

import (
	"github.com/goalbook/goalbook-auth/auth/lti/Godeps/_workspace/src/github.com/google/go-querystring/query"
)

type LTIStdResponse struct {
	LTIMessage      string `url:"lti_msg,omitempty"`
	LTIErrorMessage string `url:"lti_errormsg,omitempty"`
	LTILog          string `url:"lti_log,omitempty"`
	LTIErrorLog     string `url:"lti_errorlog,omitempty"`

	StatusCode int `url:"-"`
}

func NewLTIErrorResponse(errMsg, logMsg string, statusCode int) *LTIStdResponse {
	return &LTIStdResponse{
		LTIErrorMessage: errMsg,
		LTIErrorLog:     logMsg,

		StatusCode: statusCode,
	}
}

func NewLTIStdResposne(respMsg, logMsg string, statusCode int) *LTIStdResponse {
	return &LTIStdResponse{
		LTIMessage: respMsg,
		LTILog:     logMsg,

		StatusCode: statusCode,
	}
}

func (l *LTIStdResponse) Serialize() (string, error) {
	vals, err := query.Values(l)
	if err != nil {
		return "", err
	}

	return vals.Encode(), nil
}
