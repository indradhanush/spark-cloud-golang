package gospark

import (
	"testing"
)

var aTokenService = NewAccessTokenService(UserName, Password)

func TestGetAccessToken(t *testing.T) {

	_, err := aTokenService.GetAccessToken()
	if err != nil {
		t.Error(err)
	}
}

func TestListAllAccessTokens(t *testing.T) {

	err := aTokenService.ListAllAccessTokens()
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteAccessToken(t *testing.T) {

	// Not touching 0th and 1st tokens in list of tokens as they
	// are for "user" and "spark-ide". Also, there will always be
	// a third token as TestListAllAccessTokens is called first.
	// TODO: Might change if using Mock.
	delTokenResp, err := aTokenService.DeleteAccessToken(
		aTokenService.ATokenList.Tokens[2])
	if err != nil {
		t.Error(err)
	}
	if delTokenResp.Status == false {
		t.Error("Failed to delete token.")
	}
}
