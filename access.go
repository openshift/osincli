package osincli

import (
	"net/url"
)

type AccessRequestType string

const (
	AUTHORIZATION_CODE AccessRequestType = "authorization_code"
	REFRESH_TOKEN                        = "refresh_token"
	PASSWORD                             = "password"
	CLIENT_CREDENTIALS                   = "client_credentials"
	IMPLICIT                             = "__implicit"
)

// Access request information
type AccessRequest struct {
	client        *Client
	Type          AccessRequestType
	AuthorizeData *AuthorizeData
}

// Access data
type AccessData struct {
}

func (c *Client) NewAccessRequest(t AccessRequestType, ad *AuthorizeData) *AccessRequest {
	return &AccessRequest{
		client:        c,
		Type:          t,
		AuthorizeData: ad,
	}
}

func (c *AccessRequest) GetTokenUrl() (*url.URL, error) {
	u, err := url.Parse(c.client.Config.TokenUrl)
	if err != nil {
		return nil, err
	}

	uq := u.Query()
	uq.Add("grant_type", string(c.Type))
	uq.Add("code", c.AuthorizeData.Code)
	uq.Add("redirect_url", c.client.Config.RedirectUrl)
	if c.client.Config.SendClientSecretInParams {
		uq.Add("client_id", c.client.Config.ClientId)
		uq.Add("client_secret", c.client.Config.ClientSecret)
	}
	u.RawQuery = uq.Encode()

	return u, nil
}
