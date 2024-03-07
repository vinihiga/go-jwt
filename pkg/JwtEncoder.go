package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"

	"github.com/vinihiga/go-jwt/pkg/models"
)

type JwtEncoder struct {
	SecretKey string
}

func (instance *JwtEncoder) NewJwt(header string, payload models.ClaimsModel) string {
	var encodedHeader = base64.StdEncoding.EncodeToString([]byte(header))

	var json, _ = json.Marshal(payload)
	var encodedPayload = base64.StdEncoding.EncodeToString(json)

	encodedSignature := instance.Sign(encodedHeader, encodedPayload)

	return encodedHeader + "." + encodedPayload + "." + encodedSignature
}

func (instance *JwtEncoder) Sign(encodedHeader string, encodedPayload string) string {
	hmac := hmac.New(sha256.New, []byte(instance.SecretKey))
	var fullEncode string = encodedHeader + "." + encodedPayload
	hmac.Write([]byte(fullEncode))
	signature := hmac.Sum(nil)
	encodedSignature := hex.EncodeToString(signature)

	return encodedSignature
}
