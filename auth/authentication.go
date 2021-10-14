package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"log"
	"strings"
	"time"
)

func New(key []byte) *Auth {
	var k []byte
	if key == nil {
		k = genKey()
	} else {
		copy(k, key)
	}

	return &Auth{k}
}

// VerifyToken ...
func (a *Auth) VerifyToken(token string, period int64) (bool, string, error) {
	tokenS := strings.Split(token, ".")
	if len(tokenS) != 2 {
		return false, "", fmt.Errorf("invalid token")
	}
	payloadB, err := base64.RawStdEncoding.DecodeString(tokenS[0])
	if err != nil {
		return false, "", err
	}

	var newPlS Jwt
	if err = json.Unmarshal(payloadB, &newPlS); err != nil {
		return false, "", err
	}

	fmt.Printf("%+v\n", newPlS)
	if newPlS.Iat+period < time.Now().UTC().Unix() {
		return false, "", fmt.Errorf("token expirou")
	}

	return tokenS[1] == a.Sign(payloadB), newPlS.Username, nil
}

// Sign ...
func (a *Auth) Sign(json []byte) string {
	var hasher hash.Hash = hmac.New(sha256.New, a.key)
	hasher.Write(json)

	return base64.RawStdEncoding.EncodeToString(hasher.Sum(nil))
}

// CreateToken ...
func (a *Auth) CreateToken(plS Jwt) string {
	json, err := json.Marshal(plS)
	if err != nil {
		log.Fatal(err)
	}
	payload := base64.RawStdEncoding.EncodeToString(json)
	sign := a.Sign(json)
	return payload + "." + sign
}
