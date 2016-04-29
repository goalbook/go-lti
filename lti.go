package lti

import (
	"time"
)

var OAuthTimestampValidationRange time.Duration = time.Minute * 10

func acceptTimestamp(tstamp int) bool {
	t := time.Unix(int64(tstamp), 0)

	if t.After(time.Now().Add(OAuthTimestampValidationRange)) || t.Before(time.Now().Add(-OAuthTimestampValidationRange)) {
		return false
	}
	return true
}
