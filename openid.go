package jwtgomiddleware

type OpenidWellKnown struct {
	Issuer  string `json:"issuer"`
	JWKSURL string `json:"jwks_uri"`
}
