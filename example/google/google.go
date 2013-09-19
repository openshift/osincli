package main

import (
	osincli "../.."
	"fmt"
)

func main() {
	config := &osincli.ClientConfig{
		ClientId:                 "123",
		ClientSecret:             "1234",
		AuthorizeUrl:             "https://accounts.google.com/o/oauth2/auth",
		TokenUrl:                 "https://accounts.google.com/o/oauth2/token",
		RedirectUrl:              "oob",
		ErrorsInStatusCode:       true,
		SendClientSecretInParams: true,
	}
	client := osincli.NewClient(config)

	areq := client.NewAuthorizeRequest(osincli.CODE)

	u, err := areq.GetAuthorizeUrl()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorize URL: %s\n", u.String())

	areqdata := &osincli.AuthorizeData{
		Code: "abcdefg",
	}

	treq := client.NewAccessRequest(osincli.AUTHORIZATION_CODE, areqdata)

	u2, err := treq.GetTokenUrl()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Access URL: %s\n", u2.String())
}
