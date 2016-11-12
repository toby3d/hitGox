package hitGox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/valyala/fasthttp"
)

// AccessToken is the equivalent of a regular authToken. You are now able to use this on the hitbox API just like any other authToken.
type AccessToken struct {
	AccessToken string `json:"access_token,omitempty"`
}

// Exchange get an authentication authToken rather than account information.
//
// The hash value is a Base64 encode of the app.Token and app.Secret. As an example, you can open up the Chrome/Firefox Developer Tool, go to the console and type btoa("app_authToken"+"app_secret"); and the result would be your hash.
func (app Application) Exchange(requestToken string) (AccessToken, error) {
	hash := base64.StdEncoding.EncodeToString([]byte(app.Token + app.Secret))

	args := fasthttp.AcquireArgs()
	args.Add("request_token", requestToken)
	args.Add("app_token", app.Token)
	args.Add("hash", hash)
	statusCode, body, err := fasthttp.Post(nil, API+"/oauth/exchange", args)
	if statusCode != 200 || err != nil {
		return AccessToken{}, err
	}
	var obj AccessToken
	if err = json.NewDecoder(bytes.NewReader(body)).Decode(&obj); err != nil {
		return AccessToken{}, err
	}
	return obj, nil
}
