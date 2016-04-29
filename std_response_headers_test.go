package lti

import (
	"testing"
)

type responseTestStruct struct {
	responseStruct   *LTIStdResponse
	expectedResponse string
}

func TestSerializeStdLTIResponse(t *testing.T) {
	tests := []responseTestStruct{
		{
			responseStruct: &LTIStdResponse{
				LTIMessage: "A message from the tool provider.",
				LTILog:     "LTI log entry: A message from the tool provider.",
			},
			expectedResponse: "lti_log=LTI+log+entry%3A+A+message+from+the+tool+provider.&lti_msg=A+message+from+the+tool+provider.",
		},
		{
			responseStruct: &LTIStdResponse{
				LTIErrorMessage: "An error message from the tool provider.",
				LTIErrorLog:     "LTI error log entry: An error message from the tool provider.",
			},
			expectedResponse: "lti_errorlog=LTI+error+log+entry%3A+An+error+message+from+the+tool+provider.&lti_errormsg=An+error+message+from+the+tool+provider.",
		},
	}

	for _, test := range tests {
		seralized, err := test.responseStruct.Serialize()
		if err != nil {
			t.Error("Error Serializing: Serializing LTI Standard Response failed with error '%s'.", err)
		}
		if seralized != test.expectedResponse {
			t.Logf("Actual:\n%s\n", seralized)
			t.Logf("Expected:\n%s\n", test.expectedResponse)
			t.Error("Bad Serialization: Acutal LTI Standard Response serialization does not match expected serialization.")
		}
	}
}
