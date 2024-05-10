package telegraph

import (
	telegraph "github.com/OlegChuev/go-telegraph/api/types"
	"github.com/OlegChuev/go-telegraph/pkg"
)

// This object represents a get Telegraph account info request params.
type RevokeAccessTokenParams struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token" validate:"nonnil"`
}

// Use this method to revoke access_token and generate a new one, for example,
// if the user would like to reset all connected sessions,
// or you have reasons to believe the token was compromised.
// On success, returns an Account object with new access_token and auth_url fields.
func (c *Client) RevokeAccessToken(params *RevokeAccessTokenParams) (account *telegraph.Account, err error) {
	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "revokeAccessToken",
		Params:   params,
		Model:    telegraph.Account{},
	}

	// Make request
	result, err := pkg.PerformRequest(c.URL, &apiReq)
	if err != nil {
		return
	}

	// Cast interface{} -> Model type
	account = result.(*telegraph.Account)

	// Set AccessToken to the Client
	c.SetAccessToken(account.AccessToken)

	return
}
