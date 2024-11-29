package service

import (
	"testing"

	"github.com/oklog/ulid/v2"
)

func Test_create_happyPath(t *testing.T) {
	auth := NewAuthService()
	id := ulid.Make().String()

	token, expires, err := auth.CreateToken(id)
	if err != nil {
		t.Fatalf("Expected err to be nil, got %s", err)
	}

	if token == "" {
		t.Fatal("Expected token to be populated")
	}

	if expires.IsZero() {
		t.Fatalf("Expected expires to be populated, got %s", expires)
	}
}

func Test_parse_returnsId(t *testing.T) {
	auth := NewAuthService()
	ulid := ulid.Make().String()

	token, _, err := auth.CreateToken(ulid)
	if err != nil {
		t.Fatalf("Expected err to be nil creating token, got %s", err)
	}

	id, err := auth.Authenticate(token)
	if err != nil {
		t.Fatalf("Expected err to be nil parsing jwt, got %s", err)
	}

	if id != ulid {
		t.Fatalf("Expected ulid to be returned from token, ulid: %s, parse returned %s", ulid, id)
	}
}
