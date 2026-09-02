package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey abcde")

	got, err := GetAPIKey(header)
	want := "abcde"
	if err != nil {
		t.Fatal(err)
	}
	if !(got == want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

func TestGetApiKeyErr(t *testing.T) {
	header := http.Header{}
	header.Set("noauth", "ApiKey abcde")

	_, err := GetAPIKey(header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatal(err)
	}
}
