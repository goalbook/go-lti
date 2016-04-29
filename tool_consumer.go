package lti

import (
	"net/http"
)

type LTIToolConsumer struct {
	LTIHeaders LTIStdHeaders
}

func NewLTIToolConsumer() *LTIToolConsumer {
	return &LTIToolConsumer{}
}

func (c *LTIToolConsumer) SendLTIReqeust(l *LTIStdHeaders) (*http.Response, error) {
	return nil, nil
}
