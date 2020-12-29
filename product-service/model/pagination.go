package model

type Pagination struct {
	Sort   string `json:"sort"`
	Order  string `json:"order"`
	Offset string `json:"offset"`
	Limit  string `json:"limit"`
	Search string `json:"search"`
}
