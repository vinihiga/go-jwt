package pkg

import (
	"errors"
	"strings"
)

type JwtDecoder struct {
	SecretKey string
}

func (instance *JwtDecoder) Validate(token string) (bool, error) {
	var parts []string = strings.Split(token, ".")

	if len(parts) != 3 {
		return false, errors.New("Invalid token format")
	}

	var encodedHeader = parts[0]
	var encodedPayload = parts[1]

	var encoder = JwtEncoder{
		SecretKey: instance.SecretKey,
	}

	var encodedSignature = encoder.Sign(encodedHeader, encodedPayload)
	var duplicatedToken = encodedHeader + "." + encodedPayload + "." + encodedSignature

	return token == duplicatedToken, nil
}
