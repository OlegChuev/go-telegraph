package telegraph

// This object represents a Telegraph account.
type Account struct {
	// Account name, helps users with several accounts remember
	// which they are currently using. Displayed to the user
	// above the "Edit/Publish" button on Telegra.ph,
	// other users don't see this name.
	ShortName string `json:"short_name"`

	// Default author name used when creating new articles.
	AuthorName string `json:"author_name"`

	// Profile link, opened when users click on the author's name below the title.
	// Can be any link, not necessarily to a Telegram profile or channel.
	AuthorUrl string `json:"author_url"`

	// Optional. Only returned by the createAccount and revokeAccessToken method.
	// Access token of the Telegraph account.
	AccessToken string `json:"access_token"`

	// Optional. URL to authorize a browser on telegra.ph and connect it to a Telegraph account.
	// This URL is valid for only one use and for 5 minutes only.
	AuthUrl string `json:"auth_url"`

	// Optional. Number of pages belonging to the Telegraph account.
	PageCount int `json:"page_count"`
}
