package pkg

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/vinihiga/go-jwt/pkg/models"
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

	isValid, validationErr := instance.isInsideExpirationInterval(encodedPayload)

	if validationErr != nil {
		return false, errors.New("Couldn't find expiration claim")
	} else if !isValid {
		return false, nil
	}

	var encoder = JwtEncoder{
		SecretKey: instance.SecretKey,
	}

	var encodedSignature = encoder.Sign(encodedHeader, encodedPayload)
	var duplicatedToken = encodedHeader + "." + encodedPayload + "." + encodedSignature

	return token == duplicatedToken, nil
}

func (instance *JwtDecoder) isInsideExpirationInterval(encodedPayload string) (bool, error) {
	decodedText, decodeErr := base64.StdEncoding.DecodeString(encodedPayload)

	if decodeErr != nil {
		return false, errors.New("Couldn't find expiration claim")
	}

	claims := models.ClaimsModel{}
	parseErr := json.Unmarshal(decodedText, &claims)

	if parseErr != nil {
		return false, errors.New("Couldn't parse decoded 64 text into desired claim model")
	} else if time.Now().UnixMilli() >= claims.Exp {
		return false, nil
	}

	return true, nil
}
