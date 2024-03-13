package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/vinihiga/go-jwt/pkg/models"
)

type JwtEncoder struct {
	SecretKey string
	Algorithm string
}

// Generates a new JWT based on the provided payload.
//
// It takes a payload of type models.ClaimsModel and returns a string.
func (instance *JwtEncoder) NewJwt(payload models.ClaimsModel) string {
	var header string = fmt.Sprintf(`{ "typ": "JWT", "alg": "%s" }`, instance.Algorithm)
	var encodedHeader = base64.StdEncoding.EncodeToString([]byte(header))

	var json, _ = json.Marshal(payload)
	var encodedPayload = base64.StdEncoding.EncodeToString(json)

	encodedSignature := instance.Sign(encodedHeader, encodedPayload)

	return encodedHeader + "." + encodedPayload + "." + encodedSignature
}

// Implements signing for supported algorithms.
//
// For now, experimental version, it only supports HMAC-SHA256 (HS256).
func (instance *JwtEncoder) Sign(encodedHeader string, encodedPayload string) string {
	var result string

	// TODO: Implement support to another algorithms instead only HS256
	if instance.Algorithm == "HS256" {
		hmac := hmac.New(sha256.New, []byte(instance.SecretKey))
		var fullEncode string = encodedHeader + "." + encodedPayload
		hmac.Write([]byte(fullEncode))
		signature := hmac.Sum(nil)
		result = hex.EncodeToString(signature)
	}

	return result
}
