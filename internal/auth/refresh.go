package auth

import (
    "crypto/rand"
    "encoding/hex"
)

func MakeRefreshToken() string {
    bytes := make([]byte, 32)
    rand.Read(bytes)
    return hex.EncodeToString(bytes)
}