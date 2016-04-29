package types

import (
	"testing"

	"bytes"
	"encoding/json"
	"fmt"
)

var badCtxType contextType = "Bad context type"

func TestContextTypes(t *testing.T) {
	ctxTypeGroupUrn := contextType("urn:lti:context-type:ims/lis/Group")
	if CtxTypeGroup != ctxTypeGroupUrn {
		t.Errorf("%s context type not equal to urn %s", CtxTypeGroup, ctxTypeGroupUrn)
	}

	/*
		Test LTIContextType funcs
	*/
	var c *LTIContextType
	var err error

	c, err = NewLTIContextType(badCtxType)
	if err != errInvalidContextType {
		t.Errorf("Incorrect error: \"%s\" should return a \"%v\" error.", badCtxType, errInvalidContextType)
	}

	c, err = NewLTIContextType(CtxTypeGroup)
	if !c.HasContextType(CtxTypeGroup) {
		t.Errorf("Bad Comparison: %s should be equivalent to %s", c, CtxTypeGroup)
	}

	if c.HasContextType(CtxTypeCourseTemplate) {
		t.Errorf("Bad Comparison: %s should not be equivalent to %s", c, CtxTypeCourseTemplate)
	}

	/*
		Marshal  valid payload and check for roles
	*/
	var testLoad, bHeaders, expectedHeaders []byte

	partialHeaders := struct {
		ContextType LTIContextType `json:"context_type"`
	}{}

	testLoad = []byte(fmt.Sprintf(`{"context_type":"%s"}`, CtxTypeGroup))
	err = json.Unmarshal(testLoad, &partialHeaders)
	if err != nil {
		t.Errorf("Error Marshalling: %v", err)
	}
	if !partialHeaders.ContextType.HasContextType(CtxTypeGroup) {
		t.Errorf("Bad Unmarshal: %v should be equivalent to %s", partialHeaders.ContextType, CtxTypeGroup)
	}

	bHeaders, err = json.Marshal(partialHeaders)
	if err != nil {
		t.Errorf("Error Marshalling: %v", err)
	}
	expectedHeaders = []byte(
		fmt.Sprintf(`{"context_type":"%s"}`,
			CtxTypeGroup,
		))
	if !bytes.Equal(bHeaders, expectedHeaders) {
		t.Errorf("Bad Marshal: expected and actual marshaled byte array do not match.")
		t.Logf("Expected:\n%s", expectedHeaders)
		t.Logf("Actual:\n%s", bHeaders)
	}

	/*
		Marhsal invalid payload
	*/
	testLoad = []byte(fmt.Sprintf(`{"context_type":"%s"}`, badCtxType))
	err = json.Unmarshal(testLoad, &partialHeaders)
	if err != errInvalidContextType {
		t.Errorf("Incorrect error: \"%s\" should return a \"%v\" error.", badCtxType, errInvalidContextType)
	}
}
