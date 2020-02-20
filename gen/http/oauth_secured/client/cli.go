// Code generated by goa v3.0.10, DO NOT EDIT.
//
// oauth_secured HTTP client CLI support package
//
// Command:
// $ goa gen goauth/design

package client

import (
	"encoding/json"
	"fmt"
	oauthsecured "goauth/gen/oauth_secured"
)

// BuildMethodPayload builds the payload for the oauth_secured method endpoint
// from CLI flags.
func BuildMethodPayload(oauthSecuredMethodBody string, oauthSecuredMethodOauthToken string) (*oauthsecured.MethodPayload, error) {
	var err error
	var body MethodRequestBody
	{
		err = json.Unmarshal([]byte(oauthSecuredMethodBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"data\": \"Qui sunt dolor autem.\"\n   }'")
		}
	}
	var oauthToken string
	{
		oauthToken = oauthSecuredMethodOauthToken
	}
	v := &oauthsecured.MethodPayload{
		Data: body.Data,
	}
	v.OauthToken = oauthToken
	return v, nil
}
