package telegraph

// This object represents a list of Telegraph articles
// belonging to an account. Most recently created articles first.
type PageList struct {
	// Total number of pages belonging to the target Telegraph account.
	TotalCount int `json:"total_count"`

	// Requested pages of the target Telegraph account.
	Pages []Page `json:"pages"`
}
