package types

import (
	"strings"
)

/*
	LTI Message Types
*/
type ltiMessage string

const (
	LTIBasicMessage = ltiMessage("basic-lti-launch-request")
)

var supportedMessages = []ltiMessage{
	LTIBasicMessage,
}

type LTIMessage ltiMessage

func NewLTIMessage(s ltiMessage) (LTIMessage, error) {
	for _, m := range supportedMessages {
		if s == m {
			return LTIMessage(s), nil
		}
	}

	return LTIMessage(""), errUnsupportedLTIMessage(s)
}

// LTI Message Unmarshaler interface
func (m *LTIMessage) UnmarshalJSON(in []byte) error {
	var err error
	var cleanedIn = string(in)
	cleanedIn = strings.TrimPrefix(cleanedIn, "\"")
	cleanedIn = strings.TrimSuffix(cleanedIn, "\"")

	*m, err = NewLTIMessage(ltiMessage(cleanedIn))
	return err
}

/*
	LTI Version Types
*/

type ltiVersion string

const (
	LTIVersion1_0 = ltiVersion("LTI-1p0")
)

var supportedVersions = []ltiVersion{
	LTIVersion1_0,
}

type LTIVersion ltiVersion

func NewLTIVersion(s ltiVersion) (LTIVersion, error) {
	for _, v := range supportedVersions {
		if s == v {
			return LTIVersion(s), nil
		}
	}

	return LTIVersion(""), errUnsupportedLTIVersion(s)
}

// LTI Version JSON Unmarshaler interface
func (v *LTIVersion) UnmarshalJSON(in []byte) error {
	var err error
	var cleanedIn = string(in)
	cleanedIn = strings.TrimPrefix(cleanedIn, "\"")
	cleanedIn = strings.TrimSuffix(cleanedIn, "\"")

	*v, err = NewLTIVersion(ltiVersion(cleanedIn))
	return err
}
