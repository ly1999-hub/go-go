package model

type ResponseCreate struct {
	ID string `json:"id"`
}

type ResponseUpdate struct {
	ID string `json:"id"`
}

// ResponseDelete ...
type ResponseDelete struct {
	ID string `json:"id"`
}

type ResponseList struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Limit int64       `json:"limit"`
}

type All struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}
