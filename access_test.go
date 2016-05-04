package osincli

import "testing"

func TestGetTokenUrl(t *testing.T) {
	clientConfig := ClientConfig{
		ClientId:     "myclient",
		ClientSecret: "mysecret",
		TokenUrl:     "https://example.com/token",
		AuthorizeUrl: "https://example.com/authorize",
		RedirectUrl:  "/",
	}

	testcases := map[string]struct {
		Type   AccessRequestType
		Data   AuthorizeData
		Params map[string]string

		URL string
	}{
		"client credentials": {
			Type: CLIENT_CREDENTIALS,
			Data: AuthorizeData{State: "mystate", Code: "mycode", Username: "myusername", Password: "mypassword"},
			URL:  "https://example.com/token?grant_type=client_credentials&redirect_uri=%2F",
		},
		"code grant": {
			Type: AUTHORIZATION_CODE,
			Data: AuthorizeData{State: "mystate", Code: "mycode", Username: "myusername", Password: "mypassword"},
			URL:  "https://example.com/token?code=mycode&grant_type=authorization_code&redirect_uri=%2F",
		},
		"refresh grant": {
			Type: REFRESH_TOKEN,
			Data: AuthorizeData{State: "mystate", Code: "mycode", Username: "myusername", Password: "mypassword"},
			URL:  "https://example.com/token?grant_type=refresh_token&redirect_uri=%2F&refresh_token=mycode",
		},
		"password grant": {
			Type: PASSWORD,
			Data: AuthorizeData{State: "mystate", Code: "mycode", Username: "myusername", Password: "mypassword"},
			URL:  "https://example.com/token?grant_type=password&password=mypassword&redirect_uri=%2F&username=myusername",
		},
		"password grant with custom params": {
			Type:   PASSWORD,
			Data:   AuthorizeData{},
			Params: map[string]string{"username": "customuser", "password": "custompw"},
			URL:    "https://example.com/token?grant_type=password&password=custompw&redirect_uri=%2F&username=customuser",
		},
	}

	client, err := NewClient(&clientConfig)
	if err != nil {
		t.Fatal(err)
	}

	for k, tc := range testcases {
		req := client.NewAccessRequest(tc.Type, &tc.Data)
		req.CustomParameters = tc.Params
		url := req.GetTokenUrl().String()
		if url != tc.URL {
			t.Errorf("%s: Expected %s, got %s", k, tc.URL, url)
		}
	}
}
