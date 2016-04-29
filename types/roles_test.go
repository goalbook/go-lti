package types

import (
	"testing"

	"bytes"
	"encoding/json"
	"fmt"
)

func TestSysRoleTypes(t *testing.T) {
	sysAdminUrn := systemRole("urn:lti:sysrole:ims/lis/SysAdmin")
	if SysSysAdmin != sysAdminUrn {
		t.Errorf("%s system role type not equal to urn %s", SysSysAdmin, sysAdminUrn)
	}
}

func TestInstRoleTypes(t *testing.T) {
	instStudentUrn := institutionRole("urn:lti:instrole:ims/lis/Student")
	if InstStudent != instStudentUrn {
		t.Errorf("%s system role type not equal to urn %s", InstStudent, instStudentUrn)
	}
}

func TestCtxRoleTypes(t *testing.T) {
	ctxLearnerUrn := contextRole("urn:lti:role:ims/lis/Learner")
	if CtxLearner != ctxLearnerUrn {
		t.Errorf("%s system role type not equal to urn %s", CtxLearner, ctxLearnerUrn)
	}

	ctxLearner_LearnerUrnString := contextRole("urn:lti:role:ims/lis/Learner/Learner")
	if CtxLearner_Learner != ctxLearner_LearnerUrnString {
		t.Errorf("%s system role type not equal to %s", CtxLearner_Learner, ctxLearner_LearnerUrnString)
	}
}

func TestLTIRoles(t *testing.T) {
	var undefinedRole string = "undefinedRole"

	/*
		Marshal payload and check for roles
	*/
	testLoad := []byte(fmt.Sprintf(`{"roles":"%s,%s,%s,%s"}`, SysAdmin, InstAdmin, CtxAdmin, undefinedRole))
	partialHeaders := struct {
		Roles LTIRoles `json:"roles"`
	}{}
	json.Unmarshal(testLoad, &partialHeaders)

	if !partialHeaders.Roles.HasSystemRole(SysAdmin) {
		t.Errorf("Bad Unmarshal: %s is missing from system roles.", SysAdmin)
		t.Log(partialHeaders.Roles.systemRoles)
	}

	if !partialHeaders.Roles.HasInstitutionRole(InstAdmin) {
		t.Errorf("Bad Unmarshal: %s is missing from institution roles.", InstAdmin)
		t.Log(partialHeaders.Roles.institutionRoles)
	}

	if !partialHeaders.Roles.HasContextRole(CtxAdmin) {
		t.Errorf("Bad Unmarshal: %s is missing from context roles.", CtxAdmin)
		t.Log(partialHeaders.Roles.contextRoles)
	}

	/*
		Add roles and check for them
	*/
	partialHeaders.Roles.AddSystemRoles(SysSupport)
	if !partialHeaders.Roles.HasSystemRole(SysSupport) {
		t.Errorf("Bad Role Addition: %s was not added to system roles.", SysSupport)
		t.Log(partialHeaders.Roles.systemRoles)
	}
	partialHeaders.Roles.AddInstitutionRoles(InstStaff)
	if !partialHeaders.Roles.HasInstitutionRole(InstStaff) {
		t.Errorf("Bad Role Addition: %s was not added to institution roles.", InstStaff)
		t.Log(partialHeaders.Roles.systemRoles)
	}
	partialHeaders.Roles.AddContextRoles(CtxContentDev)
	if !partialHeaders.Roles.HasContextRole(CtxContentDev) {
		t.Errorf("Bad Role Addition: %s was not added to context roles.", CtxContentDev)
		t.Log(partialHeaders.Roles.systemRoles)
	}

	/*
		Remove some roles and check that they are removed
	*/
	partialHeaders.Roles.RemoveSystemRoles(SysAdmin, SysCreator)
	if partialHeaders.Roles.HasSystemRole(SysAdmin) {
		t.Errorf("Bad Role Removal: %s was not removed from system roles.", SysAdmin)
		t.Log(partialHeaders.Roles.systemRoles)
	}
	partialHeaders.Roles.RemoveInstitutionRoles(InstAdmin)
	if partialHeaders.Roles.HasInstitutionRole(InstAdmin) {
		t.Errorf("Bad Role Removal: %s was not removed from institution roles.", InstAdmin)
		t.Log(partialHeaders.Roles.institutionRoles)
	}
	partialHeaders.Roles.RemoveContextRoles(CtxAdmin, CtxContentDev)
	if partialHeaders.Roles.HasContextRole(CtxAdmin) {
		t.Errorf("Bad Role Removal: %s was not removed from context roles.", CtxAdmin)
		t.Log(partialHeaders.Roles.contextRoles)
	}
	if partialHeaders.Roles.HasContextRole(CtxContentDev) {
		t.Errorf("Bad Role Removal: %s was not removed from context roles.", CtxContentDev)
		t.Log(partialHeaders.Roles.contextRoles)
	}

	/*
		Marshal roles and validate output
	*/
	bHeaders, err := json.Marshal(partialHeaders)
	if err != nil {
		t.Errorf("Error Marshalling: %v", err)
	}
	expectedHeaders := []byte(
		fmt.Sprintf(`{"roles":"%s,%s,%s"}`,
			SysSupport,
			InstStaff,
			undefinedRole,
		))
	if !bytes.Equal(bHeaders, expectedHeaders) {
		t.Errorf("Bad Marshal: expected and actual marshaled byte array do not match.")
		t.Logf("Expected:\n%s", expectedHeaders)
		t.Logf("Actual:\n%s", bHeaders)
	}
}
