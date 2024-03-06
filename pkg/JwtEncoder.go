package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

type JwtEncoder struct {
	SecretKey string
}

func (instance *JwtEncoder) NewJwt(header string, payload string) string {
	var encodedHeader = base64.StdEncoding.EncodeToString([]byte(header))
	var encodedPayload = base64.StdEncoding.EncodeToString([]byte(payload))
	encodedSignature := instance.sign(encodedHeader, encodedPayload)

	return encodedHeader + "." + encodedPayload + "." + encodedSignature
}

func (instance *JwtEncoder) sign(encodedHeader string, encodedPayload string) string {
	hmac := hmac.New(sha256.New, []byte(instance.SecretKey))
	var fullEncode string = encodedHeader + "." + encodedPayload
	hmac.Write([]byte(fullEncode))
	signature := hmac.Sum(nil)
	encodedSignature := hex.EncodeToString(signature)

	return encodedSignature
}
