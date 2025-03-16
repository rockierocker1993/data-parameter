package dto

type BaseResponse struct {
	StatusCode string      `json:"status_code"`
	RequestID  string      `json:"request_id"`
	TitleID    string      `json:"title_id"`
	TitleEN    string      `json:"title_en"`
	DescID     string      `json:"desc_id"`
	DescEN     string      `json:"desc_en"`
	Source     string      `json:"source"`
	Data       interface{} `json:"data"`
}
