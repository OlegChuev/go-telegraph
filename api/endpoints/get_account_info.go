package telegraph

import (
	telegraph "github.com/OlegChuev/go-telegraph/api/types"
	"github.com/OlegChuev/go-telegraph/pkg"
)

// This object represents a get Telegraph account info request params.
type GetAccountInfoParams struct {
	// Required. Access token of the Telegraph account.
	AccessToken string `json:"access_token" validate:"nonnil"`

	// List of account fields to return.
	// Available fields: short_name, author_name, author_url, auth_url, page_count.
	// default = [“short_name”,“author_name”,“author_url”]
	Fields []string `json:"fields,omitempty"`
}

// Use this method to get information about a Telegraph account.
// Returns an Account object on success.
func (c *Client) GetAccountInfo(params *GetAccountInfoParams) (account *telegraph.Account, err error) {
	// If token is not passed, try to fetch latest from client
	if len(params.AccessToken) == 0 {
		params.AccessToken = c.GetAccessToken()
	}

	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "getAccountInfo",
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

	return
}
