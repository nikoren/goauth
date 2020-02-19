package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("multi_auth", func() {
	Title("Security Example API")
	Description("This API demonstrates the use of the goa security DSL")
	Docs(func() { // Documentation links
		Description("Security example README")
		URL("https://github.com/goadesign/goa/tree/master/example/security/README.md")
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