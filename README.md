spark-cloud-golang
==================

Golang bindings for the Spark Cloud API.

API Docs: [http://docs.spark.io/api/](http://docs.spark.io/api/)

This is a work in progress.

## Implemented
* Generate a new access token.
* List all your tokens
* Deleting an access token

## TODO
* Pretty much everything.

## Example

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
