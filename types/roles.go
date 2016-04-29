package types

import (
	"fmt"
	"strings"

	"net/url"
)

/*
	System role types
*/
type systemRole string

const (
	sRoleBase = "urn:lti:sysrole:ims/lis/"

	SysSysAdmin     = systemRole(sRoleBase + "SysAdmin")
	SysSupport      = systemRole(sRoleBase + "SysSupport")
	SysCreator      = systemRole(sRoleBase + "Creator")
	SysAccountAdmin = systemRole(sRoleBase + "AccountAdmin")
	SysUser         = systemRole(sRoleBase + "User")
	SysAdmin        = systemRole(sRoleBase + "Administrator")
	SysNone         = systemRole(sRoleBase + "None")
)

/*
	Institutional role types
*/
type institutionRole string

const (
	iRoleBase = "urn:lti:instrole:ims/lis/"

	InstStudent            = institutionRole(iRoleBase + "Student")
	InstFaculty            = institutionRole(iRoleBase + "Faculty")
	InstMember             = institutionRole(iRoleBase + "Member")
	InstLearner            = institutionRole(iRoleBase + "Learner")
	InstInstructor         = institutionRole(iRoleBase + "Instructor")
	InstMentor             = institutionRole(iRoleBase + "Mentor")
	InstStaff              = institutionRole(iRoleBase + "Staff")
	InstAlumni             = institutionRole(iRoleBase + "Alumni")
	InstProspectiveStudent = institutionRole(iRoleBase + "ProspectiveStudent")
	InstGuest              = institutionRole(iRoleBase + "Guest")
	InstOther              = institutionRole(iRoleBase + "Other")
	InstAdmin              = institutionRole(iRoleBase + "Administrator")
	InstObserver           = institutionRole(iRoleBase + "Observer")
	InstNone               = institutionRole(iRoleBase + "None")
)

/*
	Context role types
*/
type contextRole string

const (
	cRoleBase = "urn:lti:role:ims/lis/"

	CtxLearner                  = contextRole(cRoleBase + "Learner")
	CtxLearner_Learner          = contextRole(cRoleBase + "Learner/Learner")
	CtxLearner_NonCreditLearner = contextRole(cRoleBase + "Learner/NonCreditLeaner")
	CtxLearner_GuestLeaner      = contextRole(cRoleBase + "Learner/GuestLearner")
	CtxLearner_ExternalLearner  = contextRole(cRoleBase + "Learner/ExternalLearner")
	CtxLearner_Instructor       = contextRole(cRoleBase + "Learner/Instructor")

	CtxInstructor                    = contextRole(cRoleBase + "Instructor")
	CtxInstructor_PrimaryInstructor  = contextRole(cRoleBase + "Instructor/PrimaryInstructor")
	CtxInstructor_Lecturer           = contextRole(cRoleBase + "Instructor/Lecturer")
	CtxInstructor_GuestInstructor    = contextRole(cRoleBase + "Instructor/GuestInstructor")
	CtxInstructor_ExternalInstructor = contextRole(cRoleBase + "Instructor/ExternalInstructor")

	CtxContentDev                       = contextRole(cRoleBase + "ContentDeveloper")
	CtxContentDev_ContentDev            = contextRole(cRoleBase + "ContentDeveloper/ContentDeveloper")
	CtxContentDev_Librarian             = contextRole(cRoleBase + "ContentDeveloper/Librarian")
	CtxContentDev_ContentExpert         = contextRole(cRoleBase + "ContentDeveloper/ContentExpert")
	CtxContentDev_ExternalContentExpert = contextRole(cRoleBase + "ContentDeveloper/ExternalContentExpert")

	CtxMember        = contextRole(cRoleBase + "Member")
	CtxMember_Member = contextRole(cRoleBase + "Member/Member")

	CtxManager                   = contextRole(cRoleBase + "Manager")
	CtxManager_AreaManager       = contextRole(cRoleBase + "Manager/AreaManager")
	CtxManager_CourseCoordiantor = contextRole(cRoleBase + "Manager/CourseCoordianto")
	CtxManager_Observer          = contextRole(cRoleBase + "Manager/Observer")
	CtxManager_ExtObserver       = contextRole(cRoleBase + "Manager/ExternalObserver")

	CtxMentor                        = contextRole(cRoleBase + "Mentor")
	CtxMentor_Mentor                 = contextRole(cRoleBase + "Mentor/Mentor")
	CtxMentor_Reviewer               = contextRole(cRoleBase + "Mentor/Reviewer")
	CtxMentor_Advisor                = contextRole(cRoleBase + "Mentor/Advisor")
	CtxMentor_Auditor                = contextRole(cRoleBase + "Mentor/Auditor")
	CtxMentor_Tutor                  = contextRole(cRoleBase + "Mentor/Tutor")
	CtxMentor_LearningFacilitator    = contextRole(cRoleBase + "Mentor/LearningFacilitator")
	CtxMentor_ExtAdvisor             = contextRole(cRoleBase + "Mentor/ExternalAdvisor")
	CtxMentor_ExtReviewer            = contextRole(cRoleBase + "Mentor/ExternalReviewer")
	CtxMentor_ExtAuditor             = contextRole(cRoleBase + "Mentor/ExternalAuditor")
	CtxMentor_ExtTutor               = contextRole(cRoleBase + "Mentor/ExternalTutor")
	CtxMentor_ExtLearningFacilitator = contextRole(cRoleBase + "Mentor/ExternalLearningFacilitator")

	CtxAdmin             = contextRole(cRoleBase + "Administrator")
	CtxAdmin_Admin       = contextRole(cRoleBase + "Administrator/Administrator")
	CtxAdmin_Support     = contextRole(cRoleBase + "Administrator/Support")
	CtxAdmin_Dev         = contextRole(cRoleBase + "Administrator/Developer")
	CtxAdmin_SysAdmin    = contextRole(cRoleBase + "Administrator/SystemAdministrator")
	CtxAdmin_ExtSysAdmin = contextRole(cRoleBase + "Administrator/ExternalSystemAdministrator")
	CtxAdmin_ExtDev      = contextRole(cRoleBase + "Administrator/ExternalDeveloper")
	CtxAdmin_ExtSupport  = contextRole(cRoleBase + "Administrator/ExternalSupport")

	CtxTA                    = contextRole(cRoleBase + "TeachingAssistant")
	CtxTA_TA                 = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistant")
	CtxTA_Section            = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantSection")
	CtxTA_SectionAssociation = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantSectionAssociation")
	CtxTA_Offering           = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantOffering")
	CtxTA_Template           = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantTemplate")
	CtxTA_Group              = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantGroup")
	CtxTA_Grader             = contextRole(cRoleBase + "TeachingAssistant/TeachingAssistantGrader")
)

type LTIRoles struct {
	systemRoles      []systemRole
	institutionRoles []institutionRole
	contextRoles     []contextRole

	undefinedRoles []string
}

func NewLTIRoles(r string) *LTIRoles {
	l := LTIRoles{}

	rolesArr := strings.Split(r, ",")
	for _, role := range rolesArr {
		if role == "" {
			continue
		}

		if strings.Contains(role, sRoleBase) {
			l.systemRoles = append(l.systemRoles, systemRole(role))
		} else if strings.Contains(role, iRoleBase) {
			l.institutionRoles = append(l.institutionRoles, institutionRole(role))
		} else if strings.Contains(role, cRoleBase) {
			l.contextRoles = append(l.contextRoles, contextRole(role))
		} else {
			l.undefinedRoles = append(l.undefinedRoles, role)
		}
	}

	return &l
}

func (l *LTIRoles) GetSystemRoles() []systemRole {
	return l.systemRoles
}

func (l *LTIRoles) GetInstitutionRoles() []institutionRole {
	return l.institutionRoles
}

func (l *LTIRoles) GetContextRoles() []contextRole {
	return l.contextRoles
}

func (l *LTIRoles) GetUndefinedRoles() []string {
	return l.undefinedRoles
}

func (l *LTIRoles) AddSystemRoles(s ...systemRole) {
	l.systemRoles = append(l.systemRoles, s...)
}

func (l *LTIRoles) AddInstitutionRoles(i ...institutionRole) {
	l.institutionRoles = append(l.institutionRoles, i...)
}

func (l *LTIRoles) AddContextRoles(c ...contextRole) {
	l.contextRoles = append(l.contextRoles, c...)
}

func (l *LTIRoles) RemoveSystemRoles(s ...systemRole) {
	var remove = map[systemRole]bool{}
	for _, role := range s {
		remove[role] = true
	}

	var temp []systemRole
	for _, role := range l.systemRoles {
		if cut, _ := remove[role]; !cut {
			temp = append(temp, role)
		}
	}
	l.systemRoles = temp
}

func (l *LTIRoles) RemoveInstitutionRoles(i ...institutionRole) {
	var remove = map[institutionRole]bool{}
	for _, role := range i {
		remove[role] = true
	}

	var temp []institutionRole
	for _, role := range l.institutionRoles {
		if cut, _ := remove[role]; !cut {
			temp = append(temp, role)
		}
	}
	l.institutionRoles = temp
}

func (l *LTIRoles) RemoveContextRoles(c ...contextRole) {
	var remove = map[contextRole]bool{}
	for _, role := range c {
		remove[role] = true
	}

	var temp []contextRole
	for _, role := range l.contextRoles {
		if cut, _ := remove[role]; !cut {
			temp = append(temp, role)
		}
	}
	l.contextRoles = temp
}

func (l *LTIRoles) HasSystemRole(s systemRole) bool {
	for _, role := range l.systemRoles {
		if role == s {
			return true
		}
	}
	return false
}

func (l *LTIRoles) HasInstitutionRole(i institutionRole) bool {
	for _, role := range l.institutionRoles {
		if role == i {
			return true
		}
	}
	return false
}

func (l *LTIRoles) HasContextRole(c contextRole) bool {
	for _, role := range l.contextRoles {
		if role == c {
			return true
		}
	}
	return false
}

/*
	JSON Marshaler/Unmarshaler interface
*/
func (l LTIRoles) MarshalJSON() ([]byte, error) {
	var roles []string
	var rolesStr string

	for _, role := range l.systemRoles {
		roles = append(roles, string(role))
	}
	for _, role := range l.institutionRoles {
		roles = append(roles, string(role))
	}
	for _, role := range l.contextRoles {
		roles = append(roles, string(role))
	}
	roles = append(roles, l.undefinedRoles...)

	rolesStr = strings.Join(roles, ",")
	rolesStr = fmt.Sprintf("\"%s\"", rolesStr)

	return []byte(rolesStr), nil
}

func (l *LTIRoles) UnmarshalJSON(in []byte) error {
	var cleanedIn string = string(in)
	cleanedIn = strings.TrimPrefix(cleanedIn, "\"")
	cleanedIn = strings.TrimSuffix(cleanedIn, "\"")

	(*l) = *(NewLTIRoles(cleanedIn))

	return nil
}

/*
	URL Form Encoding interface
*/
func (l *LTIRoles) EncodeValues(key string, val *url.Values) error {
	b, err := l.MarshalJSON()
	if err != nil {
		return err
	}

	cleanedString := string(b)
	cleanedString = strings.TrimPrefix(cleanedString, "\"")
	cleanedString = strings.TrimSuffix(cleanedString, "\"")

	val.Add(key, cleanedString)

	return nil
}

/*
	URL Form Unmarshaler interface
*/
func (l *LTIRoles) UnmarshalText(in []byte) error {
	var cleanedIn string = string(in)
	cleanedIn = strings.TrimPrefix(cleanedIn, "\"")
	cleanedIn = strings.TrimSuffix(cleanedIn, "\"")

	(*l) = *(NewLTIRoles(cleanedIn))

	return nil
}
