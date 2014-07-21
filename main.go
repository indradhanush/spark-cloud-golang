package main

import (
	"fmt"
	"github.com/indradhanush/spark-cloud-golang/gospark"
)

func main() {

	aTokenService := gospark.NewAccessTokenService(gospark.UserName,
		gospark.Password)

	oauthToken, _ := aTokenService.GetAccessToken()
	fmt.Println(oauthToken.AccessToken)
	fmt.Println(oauthToken.TokenType)
	fmt.Println(oauthToken.ExpiresIn)

	aTokenService.ListAllAccessTokens()

	for _, token := range aTokenService.ATokenList.Tokens {
		fmt.Println(token.Token, token.ExpiresAt, token.Client)
	}

	delTokenResp, _ := aTokenService.DeleteAccessToken(
		aTokenService.ATokenList.Tokens[2])
	fmt.Println(delTokenResp.Status)

}
