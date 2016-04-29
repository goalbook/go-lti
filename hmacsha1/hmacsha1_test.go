package hmacsha1

import (
	"net/url"
	"testing"
)

// Values from https://dev.twitter.com/oauth/overview/creating-signatures
var (
	testConsumerSecret = "kAcSOqF21Fu85e7zjz7ZN2U4ZRhfV3WpwPAoE3Z7kBw"
	testTokenSecret    = "LswwdoUaIvS8ltyTt5jkRh4J50vUPVVHtR2YPi5kE"
	testBaseString     = "POST&https%3A%2F%2Fapi.twitter.com%2F1%2Fstatuses%2Fupdate.json&include_entities%3Dtrue%26oauth_consumer_key%3Dxvz1evFS4wEEPTGEFPHBog%26oauth_nonce%3DkYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg%26oauth_signature_method%3DHMAC-SHA1%26oauth_timestamp%3D1318622958%26oauth_token%3D370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb%26oauth_version%3D1.0%26status%3DHello%2520Ladies%2520%252B%2520Gentlemen%252C%2520a%2520signed%2520OAuth%2520request%2521"
	testRes            = "tnnArxj06cWHq44gCs1OSKk/jLY="
)

func TestCheckMAC(t *testing.T) {
	if !CheckMAC(testBaseString, testConsumerSecret, testTokenSecret, testRes) {
		t.Error("Bad MAC Check: Expected and generate MAC do not match.")
	}
}

// Values from https://dev.twitter.com/oauth/overview/creating-signatures
var (
	sampleMethod = "POST"
	sampleUrl    = "https://api.twitter.com/1/statuses/update.json"
	sampleForm   = url.Values{
		"status":                 {"Hello Ladies + Gentlemen, a signed OAuth request!"},
		"oauth_consumer_key":     {"xvz1evFS4wEEPTGEFPHBog"},
		"oauth_nonce":            {"kYjzVBB8Y0ZFabxSWbWovY3uYSQ2pTgmZeNu2VS4cg"},
		"oauth_signature_method": {"HMAC-SHA1"},
		"oauth_timestamp":        {"1318622958"},
		"oauth_token":            {"370773112-GmHxMAgYyLbNEtIKZeRNFsMKPR9EyMZeS9weJAEb"},
		"oauth_version":          {"1.0"},
		"include_entities":       {"true"},
	}
)

func TestRequestSignatureBaseString(t *testing.T) {
	sigBaseStr := RequestSignatureBaseString(sampleMethod, sampleUrl, sampleForm)
	if sigBaseStr != testBaseString {
		t.Logf("Actual:\n%s\n", sigBaseStr)
		t.Logf("Expected:\n%s\n", testBaseString)
		t.Error("Bad Request String: Percent encoded request string does not equal expected string.")
	}
}
