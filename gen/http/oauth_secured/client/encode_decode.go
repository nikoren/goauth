// Code generated by goa v3.0.10, DO NOT EDIT.
//
// oauth_secured HTTP client encoders and decoders
//
// Command:
// $ goa gen goauth/design

package client

import (
	"bytes"
	"context"
	oauthsecured "goauth/gen/oauth_secured"
	"io/ioutil"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildMethodRequest instantiates a HTTP request object with method and path
// set to call the "oauth_secured" service "method" endpoint
func (c *Client) BuildMethodRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MethodOauthSecuredPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("oauth_secured", "method", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeMethodRequest returns an encoder for requests sent to the
// oauth_secured method server.
func EncodeMethodRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*oauthsecured.MethodPayload)
		if !ok {
			return goahttp.ErrInvalidType("oauth_secured", "method", "*oauthsecured.MethodPayload", v)
		}
		{
			head := p.OauthToken
			req.Header.Set("Authorization", head)
		}
		body := NewMethodRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("oauth_secured", "method", err)
		}
		return nil
	}
}

// DecodeMethodResponse returns a decoder for responses returned by the
// oauth_secured method endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeMethodResponse may return the following errors:
//	- "invalid-scopes" (type oauthsecured.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type oauthsecured.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeMethodResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_secured", "method", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body MethodInvalidScopesResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_secured", "method", err)
			}
			return nil, NewMethodInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body MethodUnauthorizedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("oauth_secured", "method", err)
			}
			return nil, NewMethodUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("oauth_secured", "method", resp.StatusCode, string(body))
		}
	}
}