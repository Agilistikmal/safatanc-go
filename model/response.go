package model

type Response struct {
	StatusCode    int    `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Data          any    `json:"data,omitempty"`
}

type Responses struct {
	StatusCode    int    `json:"status_code,omitempty"`
	StatusMessage string `json:"status_message,omitempty"`
	Data          any    `json:"data,omitempty"`
	Page          int    `json:"page,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	TotalPage     int    `json:"total_page,omitempty"`
}
