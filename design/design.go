package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("oauth", func() {
	Title("Security Example API")
	Description("This API demonstrates the use of the goa security DSL")
	Docs(func() { // Documentation links
		Description("Security example README")
		URL("https://github.com/goadesign/goa/tree/master/example/security/README.md")
	})

	Server("goauth", func() {
		Host("localhost", func() {
			URI("http://0.0.0.0:8080")
		})
	})
})

// OAuth2Auth defines a security scheme that uses OAuth2 tokens.
var OAuth2Auth = OAuth2Security("oauth2", func() {

	AuthorizationCodeFlow(
		"http://goa.design/authorization",
		"http://goa.design/token",
		"http://goa.design/refresh")
	Description(`Secures endpoint by requiring a valid OAuth2 
		token retrieved via the signin endpoint. 
		Supports scopes "api:read" and "api:write".`)
	Scope("scheduler", "Scheduler access")
})

var _ = Service("oauth_secured", func() {
	Description("The secured service exposes endpoints that " +
		"require valid authorization credentials.")
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Response("unauthorized", StatusUnauthorized)
	})

	Method("method", func() {
		Description("This action requires secured oauth scopes.")
		Security(OAuth2Auth, func() {
			Scope("scheduler") // and "api:write" scopes in OAuth2 claims.
		})
		Payload(func() {
			AccessToken("oauth_token", String)
			Attribute("data", String)
			Required("oauth_token", "data")
		})
		Result(String)
		Error("invalid-scopes", String, "Token scopes are invalid")
		HTTP(func() {
			POST("/secure")
			Response(StatusOK)
			Response("invalid-scopes", StatusForbidden)
		})
	})
})

// Creds defines the credentials to use for authenticating to service methods.
var Creds = Type("Creds", func() {
	Attribute("oauth_token", String, "OAuth2 token", func() {
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
	Required("oauth_token")
})
