package pkg

import (
	"crypto/sha256"
	"encoding/hex"
)

type JwtEncoder struct {
	SecretKey string
}

func (instance *JwtEncoder) NewJwt(config string, json string) string {
	var header = sha256.Sum256([]byte(config))
	var encodedHeader = hex.EncodeToString(header[:])

	var claim = sha256.Sum256([]byte(json))
	var encodedClaim = hex.EncodeToString(claim[:])

	var signature = sha256.Sum256([]byte(instance.SecretKey))
	var encodedSignature = hex.EncodeToString(signature[:])

	return encodedHeader + "." + encodedClaim + "." + encodedSignature
}
