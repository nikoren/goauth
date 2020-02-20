package oauth

import (
	"context"
	jwt "github.com/dgrijalva/jwt-go"
	oauthsecured "goauth/gen/oauth_secured"
	"log"

	"goa.design/goa/v3/security"
)
var (
	// ErrInvalidToken is the error returned when the JWT token is invalid.
	ErrInvalidToken error = oauthsecured.Unauthorized("invalid token")

	// ErrInvalidTokenScopes is the error returned when the scopes provided in
	// the JWT token claims are invalid.
	ErrInvalidTokenScopes error = oauthsecured.InvalidScopes("invalid scopes in token")

	// Key is the key used in JWT authentication
	Key = []byte("-----BEGIN CERTIFICATE-----\nMIIDNjCCAh6gAwIBAgIIPrvfAxjPdh8wDQYJKoZIhvcNAQEFBQAwPjE8MDoGA1UE\nAxMzc2NoZWR1bGVyLmxlYXJuY2hlY2stMjU4MzAxLmlhbS5nc2VydmljZWFjY291\nbnQuY29tMB4XDTIwMDIyMDE1MzYxMVoXDTIyMDMwODAwNTczM1owPjE8MDoGA1UE\nAxMzc2NoZWR1bGVyLmxlYXJuY2hlY2stMjU4MzAxLmlhbS5nc2VydmljZWFjY291\nbnQuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAxAX0K1AWO4Ix\n9p8cZ+rtpDSLN5AoFAUD3hXVyjMlq2+SUMVP7MEQBRRsfFmPsHLAMgvFz3iclfzg\nx3Be504rje/7xXuFyx4geoNTpvL0yTj2JqbFLs4BPigq1ihUu027qvOOKU+7b/ns\nxd1K0nnlvnjoQu1Dk8In/MWhP5U3hCNwUFyue25lY76orjpXy9nzbAJ+4r82i0Zz\nkGXYPrM4pHxbZi+FRedwyxQfAfiBzeEt4VqlDRBoqs3lD/vCjeH0gxhewc8Kfo95\n77Dp/fXg00TPd1I/nB/eSpZ4vysxOBP2WD/w8LvBtwWDZKTmP/mYmvszo8fOdFud\nVtGglOXbuwIDAQABozgwNjAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB/wQEAwIHgDAW\nBgNVHSUBAf8EDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOCAQEAPLpdNp2c\n1LihqrZ6F8EqH4hhDAf3IKQHN4pn4InP6nCTtLLi/D7OrM4zJbI4feLLK2nKC62a\ndFYi0NLPkciQ7rMEN3YCO8sfQkM7tXjHHv5cmyHhAPxSi1FUN+llHGAV7vqdD80R\nIK+XL3byUjgKL4TiOg4eZVjcz+T/g8ChfteuaSPFzPbuR47N1O28gmr9PnXvbM5S\nuyqUYTw/1ofBqH5aMDz9GDdk3gxYcmyDT5sXwq/LL+8QjqeDBgdH15UTtUfTSwc0\nNZ1OZGnfN74FxozyCCkyhjqahFHeEE/bug4LsREAqnSmsMAfLg5HhUn+h4CkIzzV\njAqkgA48qq2D+Q==\n-----END CERTIFICATE-----\n")
)

// oauth_secured service example implementation.
// The example methods log the requests and return zero values.
type oauthSecuredsrvc struct {
	logger *log.Logger
}

// NewOauthSecured returns the oauth_secured service implementation.
func NewOauthSecured(logger *log.Logger) oauthsecured.Service {
	return &oauthSecuredsrvc{logger}
}

// OAuth2Auth implements the authorization logic for service "oauth_secured"
// for the "oauth2" security scheme.
func (s *oauthSecuredsrvc) OAuth2Auth(ctx context.Context, token string, scheme *security.OAuth2Scheme) (context.Context, error) {
	claims := make(jwt.MapClaims)
	// authorize request

	// 1. parse JWT token, token key is hardcoded to "secret" in this example
	_, err := jwt.ParseWithClaims(token, claims,
		func(_ *jwt.Token) (interface{}, error) { return Key, nil })
	if err != nil {
		return ctx, ErrInvalidToken
	}

	// 2. validate provided "scopes" claim
	if claims["scopes"] == nil {
		return ctx, ErrInvalidTokenScopes
	}
	scopes, ok := claims["scopes"].([]interface{})
	if !ok {
		return ctx, ErrInvalidTokenScopes
	}
	scopesInToken := make([]string, len(scopes))
	for _, scp := range scopes {
		scopesInToken = append(scopesInToken, scp.(string))
	}
	if err := scheme.Validate(scopesInToken); err != nil {
		return ctx, oauthsecured.InvalidScopes(err.Error())
	}
	return ctx, nil
}

// This action requires secured oauth scopes.
func (s *oauthSecuredsrvc) Method(ctx context.Context, p *oauthsecured.MethodPayload) (res string, err error) {
	s.logger.Print("oauthSecured.method:" + p.Data + " " + p.OauthToken)
	return
}
