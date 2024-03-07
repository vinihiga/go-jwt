package models

type ClaimsModel struct {
	Iss string // Issuer
	Sub string // Subject
	Aud string // Audience
	Exp int64  // Expiration Time in miliseconds
	Nbf string // Not before
	Iat int64  // Issued At in miliseconds
	Jti string // JWT ID
}
