// Code generated by goa v3.0.10, DO NOT EDIT.
//
// oauth_secured service
//
// Command:
// $ goa gen goauth/design

package oauthsecured

import (
	"context"

	"goa.design/goa/v3/security"
)

// The secured service exposes endpoints that require valid authorization
// credentials.
type Service interface {
	// This action requires secured oauth scopes.
	Method(context.Context, *MethodPayload) (res string, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "oauth_secured"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"method"}

// MethodPayload is the payload type of the oauth_secured service method method.
type MethodPayload struct {
	OauthToken string
	Data       string
}

// Credentials are invalid
type Unauthorized string

// Token scopes are invalid
type InvalidScopes string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}
