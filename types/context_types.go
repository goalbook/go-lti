package types

import (
	"fmt"
	"net/url"
	"strings"
)

type contextType string

const (
	ctxTypeBase = "urn:lti:context-type:ims/lis/"

	CtxTypeCourseTemplate = contextType(ctxTypeBase + "CourseTemplate")
	CtxTypeCourseOffering = contextType(ctxTypeBase + "CourseOffering")
	CtxTypeCourseSection  = contextType(ctxTypeBase + "CourseSection")
	CtxTypeGroup          = contextType(ctxTypeBase + "Group")
)

var contextTypes = []contextType{
	CtxTypeCourseOffering,
	CtxTypeCourseSection,
	CtxTypeCourseTemplate,
	CtxTypeGroup,
}

type LTIContextType []contextType

func NewLTIContextType(s ...contextType) (*LTIContextType, error) {
	ctArr := []contextType{}

	for _, t := range s {
		valid := false

		for _, validT := range contextTypes {
			if t == validT {
				valid = true
				break
			}
		}

		if !valid {
			return nil, errInvalidContextType
		}

		ctArr = append(ctArr, t)
	}

	res := (LTIContextType(ctArr))
	return &res, nil
}

func (l *LTIContextType) AddContextTypes(ctxs ...contextType) error {
	var temp []contextType
	for _, ctx := range ctxs {
		temp = append(temp, ctx)
	}

	*l = temp
	return nil
}

func (l *LTIContextType) RemoveContextTypes(ctxs ...contextType) error {
	var remove = map[contextType]bool{}
	for _, ctx := range ctxs {
		remove[ctx] = true
	}

	var temp []contextType
	for _, ctx := range *l {
		if cut, _ := remove[ctx]; !cut {
			temp = append(temp, ctx)
		}
	}

	*l = temp
	return nil
}

func (l *LTIContextType) HasContextType(ctx contextType) bool {
	for _, t := range *l {
		if ctx == t {
			return true
		}
	}
	return false
}

/*
	JSON Marshaler/Unmarshaler interface
*/
func (l LTIContextType) MarshalJSON() ([]byte, error) {
	var ctxs []string
	var ctxsStr string

	for _, ctx := range l {
		ctxs = append(ctxs, string(ctx))
	}

	ctxsStr = strings.Join(ctxs, ",")
	ctxsStr = fmt.Sprintf("\"%s\"", ctxsStr)

	return []byte(ctxsStr), nil
}

func (l *LTIContextType) UnmarshalJSON(in []byte) error {
	var newL *LTIContextType
	var err error
	var cleanedIn = string(in)
	cleanedIn = strings.TrimPrefix(cleanedIn, "\"")
	cleanedIn = strings.TrimSuffix(cleanedIn, "\"")

	ctxArr := []contextType{}
	for _, ctxStr := range strings.Split(cleanedIn, ",") {
		ctxArr = append(ctxArr, contextType(ctxStr))
	}

	newL, err = NewLTIContextType(ctxArr...)
	if err != nil {
		return err
	}

	*l = *newL
	return nil
}

/*
	URL Form Encoding interface
*/
func (l *LTIContextType) EncodeValues(key string, val *url.Values) error {
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
