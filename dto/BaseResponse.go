package dto

type BaseResponse struct {
	Status    string      `json:"status"`
	RequestID string      `json:"request_id"`
	TitleID   string      `json:"title_id"`
	TitleEN   string      `json:"title_en"`
	DescID    string      `json:"desc_id"`
	DescEN    string      `json:"desc_en"`
	Source    string      `json:"source"`
	Data      interface{} `json:"data"`
}
