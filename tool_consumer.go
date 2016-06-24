package lti

import (
	"net/http"
)

// Tool consumer is not yet implemnted.
type LTIToolConsumer struct {
	LTIHeaders LTIStdHeaders
}

// **Not implemented**
func NewLTIToolConsumer() *LTIToolConsumer {
	return &LTIToolConsumer{}
}

// **Not implemented**
func (c *LTIToolConsumer) SendLTIReqeust(l *LTIStdHeaders) (*http.Response, error) {
	return nil, nil
}
