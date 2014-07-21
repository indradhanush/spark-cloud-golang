package gospark

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

// OAuthRequest is a struct for the body of a request to get an access
// token.
type OAuthRequest struct {
	GrantType string `json:"grant_type, omitempty"`
	UserName  string `json:"username, omitempty"`
	Password  string `json:"password, omitempty"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Get returns an AccessToken in the form of a OAuthResponse object.
func (oreq *OAuthRequest) Get() (*OAuthResponse, error) {

	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl: BaseUrl,
		Endpoint: "/oauth/token"})

	form := url.Values{}
	form.Set("grant_type", oreq.GrantType)
	form.Set("username", oreq.UserName)
	form.Set("password", oreq.Password)

	req, err := http.NewRequest("POST", urlStr, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(BasicAuthId, BasicAuthPassword)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	oauthResp := &OAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(oauthResp)
	if err != nil {
		return nil, err
	}

	return oauthResp, nil
}
