package lti

import (
	"testing"

	"bytes"
	"net/http"
	"net/url"
)

// Sample Request taken from LTI test console
// http://ltiapps.net/test/tc.php
var (
	sampleMethod = "POST"
	sampleUrl    = "http://ltiapps.net/test/tp.php"
	sampleSecret = "secret"
	sampleForm   = url.Values{
		"launch_presentation_return_url": {"http://ltiapps.net/test/tc-return.php"},
		"lti_message_type":               {"basic-lti-launch-request"},
		"lti_version":                    {"LTI-1p0"},
		"oauth_callback":                 {"about:blank"},
		"oauth_consumer_key":             {"jisc.ac.uk"},
		"oauth_nonce":                    {"91587875faf645f3dc3e6917eca96cef"},
		"oauth_signature":                {"p//XBFQxwl6eiGaFQ0b4wqgsDmU="},
		"oauth_signature_method":         {"HMAC-SHA1"},
		"oauth_timestamp":                {"1443215685"},
		"oauth_version":                  {"1.0"},
		"resource_link_id":               {"429785226"},
	}
)

func TestToolProviderValidation(t *testing.T) {
	var err error
	var valid bool
	var retUrl *url.URL

	/*
		Test creating LTI Provider and validating LTI request
	*/
	r, err := http.NewRequest(sampleMethod, sampleUrl, bytes.NewBufferString(sampleForm.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		t.Errorf("Error: Request creation failed with '%s' error.", err)
	}

	provider, err := NewLTIToolProvider(r)
	if err != nil {
		t.Errorf("Error: LTI Provider creation failed with '%s' error.", err)
	}

	/*
		Validation should pass when not checking timestamp
	*/
	valid, err = provider.ValidateRequest("secret", false, false)
	if err != nil {
		t.Errorf("Error: Validation failed with error '%s'", err)
	} else if !valid {
		t.Errorf("Bad Validation Result: Signatures do not match")
	}

	retUrl, err = provider.CreateReturnURL()
	if err != nil {
		t.Errorf("Error: LTI Provider failed to create return URL with '%s' error.", err)
	}

	/*
		Validation should NOT pass when checking timestamp
	*/
	valid, err = provider.ValidateRequest("secret", true, false)
	if err == nil {
		t.Errorf("Expected Error: Expected validation error but none returned")
	} else if err != errLogInvalidTimestamp {
		t.Errorf("Unexpected Error: Validation failed with error '%s' but was expecteding error '%s'", err, errLogInvalidTimestamp)
	}

	retUrl, err = provider.CreateReturnURL()
	expectedRetUrl := "http://ltiapps.net/test/tc-return.php?lti_errorlog=LTI+Provider+Error+Log%3A+OAuth+Timestamp+is+outside+the+acceptable+range.&lti_errormsg=We+could+not+validate+your+request+to+use+our+tool."
	if err != nil {
		t.Errorf("Error: LTI Provider failed to create return URL with '%s' error.", err)
	} else if retUrl.String() != expectedRetUrl {
		t.Errorf("Error: LTI Provider failed to create the corret return URL.")
		t.Logf("Actual:\n%s\n", retUrl.String())
		t.Logf("Expected:\n%s\n", retUrl.String())
	}
}
