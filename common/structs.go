package common

type Pagination struct {
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}

type ShowData struct {
	Data interface{} `json:"data"`
}
