package main

import (
	"fmt"
	"github.com/indradhanush/spark-cloud-golang/gospark"
	"reflect"
)

func main() {
	core := gospark.NewSparkCore()
	oauthReq := &gospark.OAuthRequest{}
	oauthReq.GrantType = "password"
	oauthReq.Username = gospark.UserName
	oauthReq.Password = gospark.Password
	oauthToken, err := core.Oauth.Get(oauthReq)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Response received.")
		ref := reflect.TypeOf(oauthToken)
		fmt.Println(ref)
		fmt.Println(oauthToken)
		fmt.Println(oauthToken.AccessToken)
		fmt.Println(oauthToken.TokenType)
		fmt.Println(oauthToken.ExpiresIn)
	}
}
