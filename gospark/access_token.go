package gospark

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// AccessTokenService is an entrypoint for any AccessToken related operation.
type AccessTokenService struct {
	OaRequest  *OAuthRequest
	AToken     *AccessToken
	ATokenList *AccessTokenList
}

func NewAccessTokenService(username, password string) *AccessTokenService {
	aTokenService := &AccessTokenService{}
	aTokenService.OaRequest = NewOAuthRequest(username, password)
	aTokenService.AToken = NewAccessToken()
	aTokenService.ATokenList = NewAccessTokenList()

	return aTokenService
}

// OAuthRequest is a struct for the body of a request to get an access
// token.
type OAuthRequest struct {
	GrantType string `json:"grant_type, omitempty"`
	UserName  string `json:"username, omitoempty"`
	Password  string `json:"password, omitempty"`
}

func NewOAuthRequest(username, password string) *OAuthRequest {
	oaReq := &OAuthRequest{}
	oaReq.GrantType = "password"
	oaReq.UserName = username
	oaReq.Password = password

	return oaReq
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Get returns an AccessToken in the form of a OAuthResponse object.
func (s *AccessTokenService) GetAccessToken() (*OAuthResponse, error) {

	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl: BaseUrl,
		Endpoint: "/oauth/token"})

	form := url.Values{}
	form.Set("grant_type", s.OaRequest.GrantType)
	form.Set("username", s.OaRequest.UserName)
	form.Set("password", s.OaRequest.Password)

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

	defer resp.Body.Close()

	oauthResp := &OAuthResponse{}
	err = json.NewDecoder(resp.Body).Decode(oauthResp)
	if err != nil {
		return nil, err
	}

	return oauthResp, nil
}

type AccessToken struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
	Client    string `json:"client"`
}

func NewAccessToken() *AccessToken {
	aToken := &AccessToken{}
	return aToken
}

type AccessTokenList struct {
	Tokens []*AccessToken
}

func NewAccessTokenList() *AccessTokenList {
	aTokenList := &AccessTokenList{}
	aTokenList.Tokens = []*AccessToken{}
	return aTokenList
}

func (s *AccessTokenService) ListAllAccessTokens() error {

	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl: BaseUrl,
		APIVersion: APIVersion, Endpoint: "/access_tokens"})

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(s.OaRequest.UserName, s.OaRequest.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	s.ATokenList = NewAccessTokenList()
	err = json.NewDecoder(resp.Body).Decode(&s.ATokenList.Tokens)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

type DeleteAccessTokenResponse struct {
	Status bool `json:"ok"`
}

func (s *AccessTokenService) DeleteAccessToken(a *AccessToken) (
	*DeleteAccessTokenResponse, error) {

	Endpoint := "/access_tokens/" + a.Token
	urlStr := GetCompleteEndpointUrl(&APIUrl{BaseUrl, APIVersion,
		Endpoint})

	req, err := http.NewRequest("DELETE", urlStr, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(s.OaRequest.UserName, s.OaRequest.Password)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	delTokenResp := &DeleteAccessTokenResponse{}
	err = json.NewDecoder(resp.Body).Decode(delTokenResp)
	if err != nil {
		return nil, err
	}

	return delTokenResp, nil
}
