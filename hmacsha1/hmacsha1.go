package hmacsha1

/*
	OAuth Message Signature Base String
	Adopted and modified from https://github.com/mrjones/oauth/blob/master/oauth.go
*/

import (
	"encoding/base64"
	"fmt"
	"sort"
	"strings"

	"net/url"

	"crypto/hmac"
	"crypto/sha1"
)

func CheckMAC(message, consumerSecret, tokenSecret, oauthSig string) bool {
	key := percentEncode(consumerSecret) + "&" + percentEncode(tokenSecret)

	hashfun := hmac.New(sha1.New, []byte(key))
	hashfun.Write([]byte(message))
	rawSignature := hashfun.Sum(nil)
	base64signature := base64.StdEncoding.EncodeToString(rawSignature)

	return hmac.Equal([]byte(base64signature), []byte(oauthSig))
}

func RequestSignatureBaseString(method string, absUrl string, params url.Values) string {
	params.Del("oauth_signature")

	var paramsArr sort.StringSlice
	for key, _ := range params {
		paramsArr = append(paramsArr, fmt.Sprintf("%s=%s", percentEncode(key), percentEncode(params.Get(key))))
	}

	sort.Sort(paramsArr)
	paramString := strings.Join(paramsArr, "&")
	return method + "&" + percentEncode(absUrl) + "&" + percentEncode(paramString)
}

/*
	Helper Funcs
*/
func percentEncode(s string) string {
	t := make([]byte, 0, 3*len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isEscapable(c) {
			t = append(t, '%')
			t = append(t, "0123456789ABCDEF"[c>>4])
			t = append(t, "0123456789ABCDEF"[c&15])
		} else {
			t = append(t, s[i])
		}
	}
	return string(t)
}

func isEscapable(b byte) bool {
	return !('A' <= b && b <= 'Z' || 'a' <= b && b <= 'z' || '0' <= b && b <= '9' || b == '-' || b == '.' || b == '_' || b == '~')

}
