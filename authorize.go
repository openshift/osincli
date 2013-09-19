package osincli

import (
	"errors"
	"net/http"
	"net/url"
)

type AuthorizeRequestType string

const (
	CODE AuthorizeRequestType = "code"
	//TOKEN                      = "token"	// token not supported in server applications (uses url fragment)
)

// Authorize request information
type AuthorizeRequest struct {
	client *Client
	Type   AuthorizeRequestType
}

// Authorization data
type AuthorizeData struct {
	Code  string
	State string
}

func (c *Client) NewAuthorizeRequest(t AuthorizeRequestType) *AuthorizeRequest {
	return &AuthorizeRequest{
		client: c,
		Type:   t,
	}
}

func (c *AuthorizeRequest) GetAuthorizeUrl() (*url.URL, error) {
	return c.GetAuthorizeUrlWithParams("")
}

func (c *AuthorizeRequest) GetAuthorizeUrlWithParams(state string) (*url.URL, error) {
	u, err := url.Parse(c.client.Config.AuthorizeUrl)
	if err != nil {
		return nil, err
	}

	uq := u.Query()
	uq.Add("response_type", string(c.Type))
	uq.Add("client_id", c.client.Config.ClientId)
	uq.Add("redirect_url", c.client.Config.RedirectUrl)
	if c.client.Config.Scope != "" {
		uq.Add("scope", c.client.Config.Scope)
	}
	if state != "" {
		uq.Add("state", state)
	}
	u.RawQuery = uq.Encode()

	return u, nil
}

func (c *AuthorizeRequest) HandleRequest(r *http.Request) (*AuthorizeData, error) {
	r.ParseForm()

	var ad *AuthorizeData

	if c.Type == CODE {
		if r.Form.Get("error") != "" {
			return nil, NewError(r.Form.Get("error"), r.Form.Get("error_description"), r.Form.Get("error_uri"), r.Form.Get("state"))
		} else if r.Form.Get("code") == "" {
			return nil, errors.New("Requested parameter not sent")
		}
		ad = &AuthorizeData{
			Code:  r.Form.Get("code"),
			State: r.Form.Get("state"),
		}
	} else {
		return nil, errors.New("Unsupported response type")
	}

	return ad, nil
}
