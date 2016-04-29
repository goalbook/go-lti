package lti

import (
	"github.com/goalbook/goalbook-auth/auth/lti/hmacsha1"

	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type LTIToolProvider struct {
	LTIHeaders  *LTIStdHeaders
	LTIResponse *LTIStdResponse

	ltiConsumerRequest *http.Request

	requestProxyPath   string
	requestProxyScheme string
}

func NewLTIToolProvider(r *http.Request) (*LTIToolProvider, error) {
	var err error

	provider := LTIToolProvider{
		ltiConsumerRequest: r,
		LTIResponse:        &LTIStdResponse{},
	}

	err = r.ParseForm()
	if err != nil {
		return nil, err
	}

	contType := r.Header.Get("Content-Type")
	if contType == "application/x-www-form-urlencoded" {
		provider.LTIHeaders, err = parseUrlEncodedForm(r.Form)
	} else if contType == "application/json" {
		provider.LTIHeaders, err = parseJsonBody(r.Body)
	} else {
		return nil, errors.New(fmt.Sprintf("%s: %s", errBadContentType, contType))
	}

	if err != nil {
		return nil, err
	}

	return &provider, nil
}

func (tp *LTIToolProvider) ValidateRequest(consumerSecret string, checkTimestamp, checkNonce bool) (bool, error) {
	var err error

	defer func() {
		if err != nil {
			tp.LTIResponse.LTIErrorMessage = errInvalidRequest.Error()
			tp.LTIResponse.LTIErrorLog = err.Error()
		}
	}()

	// First check OAuth Signature
	req := tp.ltiConsumerRequest

	// Create fully qualified URL
	var requestUrl *url.URL
	if tp.requestProxyPath != "" {
		req.URL.Path = fmt.Sprintf("%s%s", tp.requestProxyPath, req.URL.Path)
	}
	if !req.URL.IsAbs() {
		requestUrl = req.URL
		if tp.requestProxyScheme != "" {
			requestUrl.Scheme = tp.requestProxyScheme
		} else if req.TLS == nil {
			requestUrl.Scheme = "http"
		} else {
			requestUrl.Scheme = "https"
		}
		requestUrl.Host = req.Host
	} else {
		requestUrl = req.URL
	}

	reqStr := hmacsha1.RequestSignatureBaseString(req.Method, requestUrl.String(), req.Form)

	if !hmacsha1.CheckMAC(reqStr, consumerSecret, "", tp.LTIHeaders.OAuthSignature) {
		err = errLogInvalidSignature
		return false, err
	}

	// Second verify that timestamp is withing acceptable range
	if checkTimestamp {
		tstamp := tp.LTIHeaders.OAuthTimestamp

		if !acceptTimestamp(tstamp) {
			err = errLogInvalidTimestamp
			return false, err
		}
	}
	// Third, make sure unique nonce
	if checkNonce {
		// TODO: Nonce verification
	}

	return true, nil
}

func (tp *LTIToolProvider) CreateReturnURL() (*url.URL, error) {
	if tp.LTIHeaders == nil || tp.LTIHeaders.LaunchPresReturnURL == "" {
		return nil, nil
	}

	urlParams, err := tp.LTIResponse.Serialize()
	if err != nil {
		return nil, err
	}

	returnUrl := tp.LTIHeaders.LaunchPresReturnURL + "?" + urlParams
	return url.Parse(returnUrl)
}

// TODO: remove this, pull proxy path from headers
func (tp *LTIToolProvider) SetProxyPathPrefix(proxyPath string) {
	proxyPath = strings.TrimPrefix(proxyPath, "/")
	proxyPath = strings.TrimSuffix(proxyPath, "/")
	tp.requestProxyPath = proxyPath
}

func (tp *LTIToolProvider) SetProxyScheme(scheme string) {
	tp.requestProxyScheme = scheme
}
