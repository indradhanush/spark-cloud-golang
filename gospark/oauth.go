package gospark

// OAuthService is an entrypoint for any OAuth2 related operation.
type OAuthService struct {
	sparkCore *SparkCore `endpoint:/oauth/token`
}

// OAuthRequest is a struct for the body of a request to get an access
// token.
type OAuthRequest struct {
	GrantType string `json:"grant_type, omitempty"`
	Username  string `json:"username, omitempty"`
	Password  string `json:"password, omitempty"`
}

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// Get returns an AccessToken in the form of a OAuthResponse object.
func (s *OAuthService) Get(oreq *OAuthRequest) (*OAuthResponse, error) {
	url := BaseUrl + GetEndpoint(*oreq)
	req, err := s.sparkCore.NewRequest("POST", url, oreq)
	req.SetBasicAuth(BasicAuthId, BasicAuthPassword)
	if err != nil {
		return nil, err
	}

	resp, err := s.sparkCore.Do(req, &OAuthResponse{})
	if err != nil {
		return nil, err
	}

	if r, ok := resp.(OAuthResponse); ok {
		return &r, err
	} else {
		return nil, err
	}

}
