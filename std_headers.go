package lti

import (
	"github.com/goalbook/goalbook-auth/auth/lti/types"

	"bytes"
	"io"
	"io/ioutil"

	"encoding/json"
	"github.com/goalbook/goalbook-auth/auth/lti/Godeps/_workspace/src/github.com/google/go-querystring/query"
	"github.com/goalbook/goalbook-auth/auth/lti/Godeps/_workspace/src/github.com/gorilla/schema"

	"net/url"
)

const (
	ContextID    = "context_id"
	ContextLabel = "context_label"
	ContextTitle = "context_title"
	ContextType  = "context_type"

	LaunchPresHeight    = "launch_presentation_height"
	LaunchPresLocale    = "launch_presentation_locale"
	LaunchPresTarget    = "launch_presentation_document_target"
	LaunchPresReturnURL = "launch_presentation_return_url"
	LaunchPresWidth     = "launch_presentation_width"

	LISCourseOfferingSID = "lis_course_offering_sourcedid"
	LISCourseSectionSID  = "lis_course_section_sourcedid"

	LISPersonFamilyName   = "lis_person_name_family"
	LISPersonFullName     = "lis_person_name_full"
	LISPersonGivenName    = "lis_person_name_given"
	LISPersonPrimaryEmail = "lis_person_contact_email_primary"
	LISPersonSID          = "lis_person_sourcedid"

	LISResultSID = "lis_result_souredid"

	LTIMessageType = "lti_message_type"
	LTIVersion     = "lti_version"

	OAuthCallback        = "oauth_callback"
	OAuthConsumerKey     = "oauth_consumer_key"
	OAuthNonce           = "oauth_nonce"
	OAuthSignature       = "oauth_signature"
	OAuthSignatureMethod = "oauth_signature_method"
	OAuthTimestamp       = "oauth_timestamp"
	OAuthVersion         = "oauth_version"

	ResourceLinkID    = "resource_link_id"
	ResourceLinkTitle = "resource_link_title"
	ResourceLinkDesc  = "resource_link_description"

	Roles = "roles"

	ToolConsumerInstGUID         = "tool_consumer_instance_guid"
	ToolConsumerInstName         = "tool_consumer_instance_name"
	ToolConsumerInstDesc         = "tool_consumer_instance_description"
	ToolConsumerInstURL          = "tool_consumer_instance_url"
	ToolConsumerInstContactEmail = "tool_consumer_instance_contact_email"

	UserId    = "user_id"
	UserImage = "user_image"
)

type LTIStdHeaders struct {
	ContextID    string                `json:"context_id" url:"context_id,omitempty"`
	ContextLabel string                `json:"context_label" url:"context_label,omitempty"`
	ContextTitle string                `json:"context_title" url:"context_title,omitempty"`
	ContextType  *types.LTIContextType `json:"context_type" url:"context_type,omitempty"`

	LaunchPresHeight    string `json:"launch_presentation_height" url:"launch_presentation_height,omitempty"`
	LaunchPresLocale    string `json:"launch_presentation_locale" url:"launch_presentation_locale,omitempty"`
	LaunchPresTarget    string `json:"launch_presentation_document_target" url:"launch_presentation_document_target,omitempty"`
	LaunchPresReturnURL string `json:"launch_presentation_return_url" url:"launch_presentation_return_url,omitempty"`
	LaunchPresWidth     string `json:"launch_presentation_width" url:"launch_presentation_width,omitempty"`

	LISCourseOfferingSID string `json:"lis_course_offering_sourcedid" url:"lis_course_offering_sourcedid,omitempty"`
	LISCourseSectionSID  string `json:"lis_course_section_sourcedid" url:"lis_course_section_sourcedid,omitempty"`

	LISPersonFamilyName   string `json:"lis_person_name_family" url:"lis_person_name_family,omitempty"`
	LISPersonFullName     string `json:"lis_person_name_full" url:"lis_person_name_full,omitempty"`
	LISPersonGivenName    string `json:"lis_person_name_given" url:",omitempty"`
	LISPersonPrimaryEmail string `json:"lis_person_contact_email_primary" url:"lis_person_contact_email_primary,omitempty"`
	LISPersonSID          string `json:"lis_person_sourcedid" url:"lis_person_sourcedid,omitempty"`

	LISResultSID string `json:"lis_result_souredid" url:"lis_result_souredid,omitempty"`

	LTIMessageType types.LTIMessage `json:"lti_message_type" url:"lti_message_type,omitempty"`
	LTIVersion     types.LTIVersion `json:"lti_version" url:"lti_version,omitempty"`

	OAuthCallback        string `json:"oauth_callback" url:"oauth_callback,omitempty"`
	OAuthConsumerKey     string `json:"oauth_consumer_key" url:"oauth_consumer_key,omitempty"`
	OAuthNonce           string `json:"oauth_nonce" url:"oauth_nonce,omitempty"`
	OAuthSignature       string `json:"oauth_signature" url:"oauth_signature,omitempty"`
	OAuthSignatureMethod string `json:"oauth_signature_method" url:"oauth_signature_method,omitempty"`
	OAuthTimestamp       int    `json:"oauth_timestamp" url:"oauth_timestamp,omitempty"`
	OAuthVersion         string `json:"oauth_version" url:"oauth_version,omitempty"`

	ResourceLinkID    string `json:"resource_link_id" url:"resource_link_id,omitempty"`
	ResourceLinkTitle string `json:"resource_link_title" url:"resource_link_title,omitempty"`
	ResourceLinkDesc  string `json:"resource_link_description" url:"resource_link_description,omitempty"`

	Roles *types.LTIRoles `json:"roles" url:"roles,omitempty"`

	ToolConsumerInstGUID         string `json:"tool_consumer_instance_guid" url:"tool_consumer_instance_guid,omitempty"`
	ToolConsumerInstName         string `json:"tool_consumer_instance_name" url:"tool_consumer_instance_name,omitempty"`
	ToolConsumerInstDesc         string `json:"tool_consumer_instance_description" url:"tool_consumer_instance_description,omitempty"`
	ToolConsumerInstURL          string `json:"tool_consumer_instance_url" url:"tool_consumer_instance_url,omitempty"`
	ToolConsumerInstContactEmail string `json:"tool_consumer_instance_contact_email" url:"tool_consumer_instance_contact_email,omitempty"`

	UserId    string `json:"user_id" schema:"user_id" url:"user_id,omitempty"`
	UserImage string `json:"user_image" schema:"user_image" url:"user_image,omitempty"`
}

// Expert from http://www.gorillatoolkit.org/pkg/Schema
// Note: it is a good idea to set a Decoder instance as a package global, because it caches meta-data about structs, and a instance can be shared safely
var formDecoder *schema.Decoder

func init() {
	formDecoder = schema.NewDecoder()
	formDecoder.SetAliasTag("url")
	formDecoder.IgnoreUnknownKeys(true)
}

func parseJsonBody(body io.Reader) (*LTIStdHeaders, error) {
	l := &LTIStdHeaders{}
	err := json.NewDecoder(body).Decode(l)
	return l, err
}

func parseUrlEncodedForm(m map[string][]string) (*LTIStdHeaders, error) {
	l := &LTIStdHeaders{}
	err := formDecoder.Decode(l, m)
	return l, err
}

func ParseXMLEncodedFile(filename string) (*LTIStdHeaders, error) { return nil, nil }

func serializeJsonBody(l LTIStdHeaders) (io.ReadCloser, error) {
	b, err := json.Marshal(l)
	if err != nil {
		return nil, err
	}
	return ioutil.NopCloser(bytes.NewReader(b)), nil
}

func serializeUrlEncodedForm(l LTIStdHeaders) (url.Values, error) {
	return query.Values(l)
}
