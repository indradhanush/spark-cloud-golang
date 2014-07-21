package main

import (
	"fmt"
	"github.com/indradhanush/spark-cloud-golang/gospark"
)

func main() {

	aTokenService := gospark.NewAccessTokenService(gospark.UserName,
		gospark.Password)
	aTokenService.ListAllAccessTokens()
	fmt.Println(aTokenService.ATokenList)

	for _, token := range aTokenService.ATokenList.Tokens {
		fmt.Println(token.Token, token.ExpiresAt, token.Client)
	}

	oauthToken, _ := aTokenService.GetAccessToken()
	fmt.Println(oauthToken.AccessToken)
	fmt.Println(oauthToken.TokenType)
	fmt.Println(oauthToken.ExpiresIn)

}
