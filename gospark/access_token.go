package gospark

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// AccessTokenService is an entrypoint for any AccessToken related operation.
type AccessTokenService struct {
	Oreq       *OAuthRequest
	AToken     *AccessToken
	ATokenList *AccessTokenList
}

func NewAccessTokenService() *AccessTokenService {
	aTokenService := &AccessTokenService{}
	aTokenService.AToken = NewAccessToken()
	aTokenService.ATokenList = NewAccessTokenList()

	return aTokenService
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

	req.SetBasicAuth(UserName, Password)

	if err != nil {
		return err
	}

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
	fmt.Println(s.ATokenList.Tokens)

	return nil
}
