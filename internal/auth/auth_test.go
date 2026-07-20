package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJWT(t *testing.T) {

	userID := uuid.New()
	secret := "my-secret"

	token, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	gotID, err := ValidateJWT(token, secret)
	if err != nil {
		t.Fatal(err)
	}

	if gotID != userID {
		t.Fatalf("expected %v got %v", userID, gotID)
	}
}

func TestExpiredJWT(t *testing.T) {

	userID := uuid.New()
	secret := "secret"

	token, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, secret)

	if err == nil {
		t.Fatal("expected expired token error")
	}
}

func TestWrongSecret(t *testing.T) {

	userID := uuid.New()

	token, err := MakeJWT(userID, "secret1", time.Hour)
	if err != nil {
		t.Fatal(err)
	}

	_, err = ValidateJWT(token, "secret2")

	if err == nil {
		t.Fatal("expected invalid signature")
	}
}

func TestGetBearerToken(t *testing.T) {

	h := http.Header{}
	h.Set("Authorization", "Bearer abc123")

	token, err := GetBearerToken(h)

	if err != nil {
		t.Fatal(err)
	}

	if token != "abc123" {
		t.Fatal("wrong token")
	}
}