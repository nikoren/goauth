// Code generated by goa v3.0.10, DO NOT EDIT.
//
// oauth_secured HTTP client types
//
// Command:
// $ goa gen goauth/design

package client

import (
	oauthsecured "goauth/gen/oauth_secured"
)

// MethodRequestBody is the type of the "oauth_secured" service "method"
// endpoint HTTP request body.
type MethodRequestBody struct {
	Data string `form:"data" json:"data" xml:"data"`
}

// MethodInvalidScopesResponseBody is the type of the "oauth_secured" service
// "method" endpoint HTTP response body for the "invalid-scopes" error.
type MethodInvalidScopesResponseBody string

// MethodUnauthorizedResponseBody is the type of the "oauth_secured" service
// "method" endpoint HTTP response body for the "unauthorized" error.
type MethodUnauthorizedResponseBody string

// NewMethodRequestBody builds the HTTP request body from the payload of the
// "method" endpoint of the "oauth_secured" service.
func NewMethodRequestBody(p *oauthsecured.MethodPayload) *MethodRequestBody {
	body := &MethodRequestBody{
		Data: p.Data,
	}
	return body
}

// NewMethodInvalidScopes builds a oauth_secured service method endpoint
// invalid-scopes error.
func NewMethodInvalidScopes(body MethodInvalidScopesResponseBody) oauthsecured.InvalidScopes {
	v := oauthsecured.InvalidScopes(body)
	return v
}

// NewMethodUnauthorized builds a oauth_secured service method endpoint
// unauthorized error.
func NewMethodUnauthorized(body MethodUnauthorizedResponseBody) oauthsecured.Unauthorized {
	v := oauthsecured.Unauthorized(body)
	return v
}
