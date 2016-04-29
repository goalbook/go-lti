package lti

import (
	"testing"

	"github.com/goalbook/goalbook-auth/auth/lti/types"
	"net/url"

	"bytes"
	"fmt"
	"io/ioutil"
)

var testPrimaryEmail = "mohammads@goalbookapp.com"
var testFullName = "Mo Samman"

var badCtxType = "badCtxType"
var badLTIMessage = "badLTIMessage"
var badLTIVersion = "badLTIVersion"

type LTINonStdHeaders struct {
	LTIStdHeaders
	NonStandardHeader string
}

func TestParseJsonBody(t *testing.T) {
	var err error
	var body string
	var l *LTIStdHeaders

	template := `{
		"context_type": "%s",
		"lti_message_type": "%s",
		"lti_version": "%s",
		"roles":"%s,%s",
		"lis_person_contact_email_primary":"%s",
		"lis_person_name_full": "%s"
	}`

	/*
		Bad context type test
	*/
	body = fmt.Sprintf(template,
		//types.CtxTypeGroup,
		badCtxType,
		types.LTIBasicMessage,
		types.LTIVersion1_0,
		types.SysAdmin,
		types.InstAdmin,
		testPrimaryEmail,
		testFullName,
	)
	l, err = parseJsonBody(bytes.NewReader([]byte(body)))
	if err == nil {
		t.Errorf("Expected error: Bad context type should trigger error")
	}

	/*
		Bad lti message type test
	*/
	body = fmt.Sprintf(template,
		types.CtxTypeGroup,
		badLTIMessage,
		types.LTIVersion1_0,
		types.SysAdmin,
		types.InstAdmin,
		testPrimaryEmail,
		testFullName,
	)
	l, err = parseJsonBody(bytes.NewReader([]byte(body)))
	if err == nil {
		t.Errorf("Expected error: Bad message type should trigger error")
	}

	/*
		Bad lti version type test
	*/
	body = fmt.Sprintf(template,
		types.CtxTypeGroup,
		types.LTIBasicMessage,
		badLTIVersion,
		types.SysAdmin,
		types.InstAdmin,
		testPrimaryEmail,
		testFullName,
	)
	l, err = parseJsonBody(bytes.NewReader([]byte(body)))
	if err == nil {
		t.Errorf("Expected error: Bad version type should trigger error")
	}

	/*
		Parse valid body test
	*/
	body = fmt.Sprintf(template,
		types.CtxTypeGroup,
		types.LTIBasicMessage,
		types.LTIVersion1_0,
		types.SysAdmin,
		types.InstAdmin,
		testPrimaryEmail,
		testFullName,
	)
	l, err = parseJsonBody(bytes.NewReader([]byte(body)))
	if err != nil {
		t.Errorf("Unexpected error: body parse should have succeeded but got error '%s'", err)
	}
	if !l.ContextType.HasContextType(types.CtxTypeGroup) {
		t.Errorf("Bad marshal: Actual context type %s should be %s", l.ContextType, types.CtxTypeGroup)
	}
	if l.LTIMessageType != types.LTIMessage(types.LTIBasicMessage) {
		t.Errorf("Bad marshal: Actual message type %s should be %s", l.LTIMessageType, types.LTIBasicMessage)
	}
	if l.LTIVersion != types.LTIVersion(types.LTIVersion1_0) {
		t.Errorf("Bad marshal: Actual message type %s should be %s", l.LTIVersion, types.LTIVersion1_0)
	}
	if l.LISPersonPrimaryEmail != testPrimaryEmail {
		t.Errorf("Bad marshal: Actual primary email %s should be %s", l.LISPersonPrimaryEmail, testPrimaryEmail)
	}
	if l.LISPersonFullName != testFullName {
		t.Errorf("Bad marshal: Actual person full name %s should be %s", l.LISPersonFullName, testFullName)
	}
	if !l.Roles.HasSystemRole(types.SysAdmin) {
		t.Errorf("Bad marshal:  Of the present roles, %s, missing %s system role", l.Roles.GetSystemRoles(), types.SysAdmin)
	}
	if !l.Roles.HasInstitutionRole(types.InstAdmin) {
		t.Errorf("Bad marshal: Of the present roles, %s, missing %s institution role", l.Roles.GetInstitutionRoles(), types.InstAdmin)
	}
	if len(l.Roles.GetContextRoles()) != 0 {
		t.Errorf("Bad marshal: There should be no context roles present, but", l.Roles.GetContextRoles())
	}
	if len(l.Roles.GetUndefinedRoles()) != 0 {
		t.Errorf("Bad marshal: There should be no undefined roles present, but", l.Roles.GetUndefinedRoles())
	}
}

func TestParseUrlEncodedForm(t *testing.T) {
	var l *LTIStdHeaders
	var err error

	vals := url.Values{}

	vals.Add(ContextType, badCtxType)
	vals.Add(LISPersonPrimaryEmail, testPrimaryEmail)
	vals.Add(LISPersonFullName, testFullName)
	vals.Add(LTIMessageType, string(types.LTIBasicMessage))
	vals.Add(LTIVersion, string(types.LTIVersion1_0))
	vals.Add(Roles, fmt.Sprintf("%s,%s", types.SysAdmin, types.InstAdmin))

	/*
		Bad context type test
	*/
	l, err = parseUrlEncodedForm(vals)
	if err == nil {
		t.Logf("(IGNORING) Expected error: Bad context type should trigger error")
	}

	/*
		Bad lti message type test
	*/
	vals.Del(ContextType)
	vals.Add(ContextType, string(types.CtxTypeGroup))
	vals.Del(LTIMessageType)
	vals.Add(LTIMessageType, badLTIMessage)

	l, err = parseUrlEncodedForm(vals)
	if err == nil {
		t.Logf("(IGNORING) Expected error: Bad LTI message type should trigger error")
	}

	/*
		Bad lti version type test
	*/
	vals.Del(LTIMessageType)
	vals.Add(LTIMessageType, string(types.LTIBasicMessage))
	vals.Del(LTIVersion)
	vals.Add(LTIVersion, badLTIVersion)

	l, err = parseUrlEncodedForm(vals)
	if err == nil {
		t.Logf("(IGNORING) Expected error: Bad LTI version type should trigger error")
	}

	/*
		Parse valid body test
	*/
	vals.Del(LTIVersion)
	vals.Add(LTIVersion, string(types.LTIVersion1_0))

	l, err = parseUrlEncodedForm(vals)
	if err != nil {
		t.Errorf("Unexpected error: body parse should have succeeded but got error '%s'", err)
	}
	if !l.ContextType.HasContextType(types.CtxTypeGroup) {
		t.Errorf("Bad marshal: Actual context type %s should be %s", l.ContextType, types.CtxTypeGroup)
	}
	if l.LTIMessageType != types.LTIMessage(types.LTIBasicMessage) {
		t.Errorf("Bad marshal: Actual message type %s should be %s", l.LTIMessageType, types.LTIBasicMessage)
	}
	if l.LTIVersion != types.LTIVersion(types.LTIVersion1_0) {
		t.Errorf("Bad marshal: Actual message type %s should be %s", l.LTIVersion, types.LTIVersion1_0)
	}
	if l.LISPersonPrimaryEmail != testPrimaryEmail {
		t.Errorf("Bad marshal: Actual primary email %s should be %s", l.LISPersonPrimaryEmail, testPrimaryEmail)
	}
	if l.LISPersonFullName != testFullName {
		t.Errorf("Bad marshal: Actual person full name %s should be %s", l.LISPersonFullName, testFullName)
	}
	if !l.Roles.HasSystemRole(types.SysAdmin) {
		t.Errorf("Bad marshal:  Of the present roles, %s, missing %s system role", l.Roles.GetSystemRoles(), types.SysAdmin)
	}
	if !l.Roles.HasInstitutionRole(types.InstAdmin) {
		t.Errorf("Bad marshal: Of the present roles, %s, missing %s institution role", l.Roles.GetInstitutionRoles(), types.InstAdmin)
	}
	if len(l.Roles.GetContextRoles()) != 0 {
		t.Errorf("Bad marshal: There should be no context roles present, but", l.Roles.GetContextRoles())
	}
	if len(l.Roles.GetUndefinedRoles()) != 0 {
		t.Errorf("Bad marshal: There should be no undefined roles present, but", l.Roles.GetUndefinedRoles())
	}

	/*
		Ignore custom and ext headers for now
	*/
	vals.Add("custom_lineitems_url", "line_items")
	l, err = parseUrlEncodedForm(vals)
	if err != nil {
		t.Errorf("Unexpected error: body parse should have succeeded but got error '%s'", err)
	}
}

func TestSerializeJsonBody(t *testing.T) {
	roles := types.NewLTIRoles("")
	roles.AddSystemRoles(types.SysAdmin)
	roles.AddInstitutionRoles(types.InstAdmin)

	ctxType, _ := types.NewLTIContextType(types.CtxTypeGroup)
	ltiMsg, _ := types.NewLTIMessage(types.LTIBasicMessage)
	ltiVer, _ := types.NewLTIVersion(types.LTIVersion1_0)

	l := LTIStdHeaders{
		ContextType:           ctxType,
		LISPersonPrimaryEmail: testPrimaryEmail,
		LISPersonFullName:     testFullName,
		LTIMessageType:        ltiMsg,
		LTIVersion:            ltiVer,
		Roles:                 roles,
	}

	reader, err := serializeJsonBody(l)
	if err != nil {
		t.Errorf("Unexpected error: Serializing LTIStdHeaders failed with unexpected error '%s'", err)
	}

	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Errorf("Unexpected error: Reading serialized LTIStdHeader failed with unexpected error '%s'", err)
	}

	if !bytes.Contains(b, []byte(`"context_type":"urn:lti:context-type:ims/lis/Group"`)) {
		t.Errorf("Bad serialization: Context type serialized incorrectly.")
	}
	if !bytes.Contains(b, []byte(`"lis_person_name_full":"Mo Samman"`)) {
		t.Errorf("Bad serialization: Person name full serialized incorrectly.")
	}
	if !bytes.Contains(b, []byte(`"lis_person_contact_email_primary":"mohammads@goalbookapp.com"`)) {
		t.Errorf("Bad serialization: Person contact email serialized incorrectly.")
	}
	if !bytes.Contains(b, []byte(`"lti_message_type":"basic-lti-launch-request"`)) {
		t.Errorf("Bad serialization: LTI message type serialized incorrectly.")
	}
	if !bytes.Contains(b, []byte(`"lti_version":"LTI-1p0"`)) {
		t.Errorf("Bad serialization: LTI version serialized incorrectly.")
	}
	if !bytes.Contains(b, []byte(`"roles":"urn:lti:sysrole:ims/lis/Administrator,urn:lti:instrole:ims/lis/Administrator"`)) {
		t.Errorf("Bad serialization: Roles serialized incorrectly.")
	}
}

func TestSerializeUrlEncodedForm(t *testing.T) {
	roles := types.NewLTIRoles("")
	roles.AddSystemRoles(types.SysAdmin)
	roles.AddInstitutionRoles(types.InstAdmin)

	ctxType, _ := types.NewLTIContextType(types.CtxTypeGroup)
	ltiMsg, _ := types.NewLTIMessage(types.LTIBasicMessage)
	ltiVer, _ := types.NewLTIVersion(types.LTIVersion1_0)

	l := LTIStdHeaders{
		ContextType:           ctxType,
		LISPersonPrimaryEmail: testPrimaryEmail,
		LISPersonFullName:     testFullName,
		LTIMessageType:        ltiMsg,
		LTIVersion:            ltiVer,
		Roles:                 roles,
	}

	vals, err := serializeUrlEncodedForm(l)

	if err != nil {
		t.Errorf("Unexpected error: Serializing LTIStdHeaders failed with unexpected error '%s'", err)
	}
	if vals.Get(ContextType) != string(types.CtxTypeGroup) {
		t.Errorf("Bad serialization: Context type serialized as '%s' instead of '%s'.", vals.Get(ContextType), types.CtxTypeGroup)
	}
	if vals.Get(LISPersonFullName) != testFullName {
		t.Errorf("Bad serialization: Person name full serialized as '%s' instead of '%s'.", vals.Get(LISPersonFullName), testFullName)
	}
	if vals.Get(LISPersonPrimaryEmail) != testPrimaryEmail {
		t.Errorf("Bad serialization: Person contact email serialized as '%s' instead of '%s'.", vals.Get(LISPersonPrimaryEmail), testPrimaryEmail)
	}
	if vals.Get(LTIMessageType) != string(types.LTIBasicMessage) {
		t.Errorf("Bad serialization: LTI message type serialized as '%s' instead of '%s'.", vals.Get(LTIMessageType), types.LTIBasicMessage)
	}
	if vals.Get(LTIVersion) != string(types.LTIVersion1_0) {
		t.Errorf("Bad serialization: LTI version serialized as '%s' instead of '%s'.", vals.Get(LTIVersion), types.LTIVersion1_0)
	}
	if vals.Get(Roles) != "urn:lti:sysrole:ims/lis/Administrator,urn:lti:instrole:ims/lis/Administrator" {
		t.Errorf("Bad serialization: Roles serialized as '%s' instead of '%s'.", vals.Get(Roles), fmt.Sprintf("%s,%s", types.SysAdmin, types.InstAdmin))
	}
}
