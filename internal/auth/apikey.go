package auth

import (
    "errors"
    "net/http"
    "strings"
)

func GetAPIKey(headers http.Header) (string, error) {
    authHeader := headers.Get("Authorization")

    if authHeader == "" {
        return "", errors.New("missing authorization header")
    }

    const prefix = "ApiKey "

    if !strings.HasPrefix(authHeader, prefix) {
        return "", errors.New("invalid authorization header")
    }

    return strings.TrimPrefix(authHeader, prefix), nil
}