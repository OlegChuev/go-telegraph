package telegraph

import (
	telegraph "github.com/OlegChuev/go-telegraph/api/types"
	"github.com/OlegChuev/go-telegraph/pkg"
)

// This object represents a create Telegraph account request params.
type CreateAccountParams struct {
	// Required. Account name, helps users with several accounts remember which they are currently using.
	// Displayed to the user above the "Edit/Publish" button on Telegra.ph, other users don't see this name.
	ShortName string `json:"short_name" validate:"min=1,max=32"`

	// Default author name used when creating new articles.
	AuthorName string `json:"author_name,omitempty" validate:"min=0,max=128"`

	// Default profile link, opened when users click on the author's name below the title.
	// Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url,omitempty" validate:"min=0,max=512"`
}

// Use this method to create a new Telegraph account.
// Most users only need one account, but this can be useful for channel administrators
// who would like to keep individual author names and profile links for each of their channels.
// On success, request returns an Account object with the regular
// fields and an additional access_token field.
func (c *Client) CreateAccount(params *CreateAccountParams) (account *telegraph.Account, err error) {
	// Prepare request config
	apiReq := telegraph.ApiReq{
		Endpoint: "createAccount",
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
