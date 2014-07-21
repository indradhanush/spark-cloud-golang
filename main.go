package main

import (
	"fmt"
	"github.com/indradhanush/spark-cloud-golang/gospark"
)

func main() {
	oauthReq := &gospark.OAuthRequest{"password", gospark.UserName,
		gospark.Password}

	oauthToken, _ := oauthReq.Get()
	fmt.Println(oauthToken.AccessToken)
	fmt.Println(oauthToken.TokenType)
	fmt.Println(oauthToken.ExpiresIn)

	aTokenService := gospark.NewAccessTokenService()
	aTokenService.ListAllAccessTokens()
	fmt.Println(aTokenService.ATokenList)

	for _, token := range aTokenService.ATokenList.Tokens {
		fmt.Println(token.Token, token.ExpiresAt, token.Client)
	}

}
