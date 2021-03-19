package model

type Package struct {
	ID          int    `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
}

type MockyResponse struct {
	Message string `json:"message"`
}
