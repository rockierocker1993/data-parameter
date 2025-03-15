package dto

type LookupValueDto struct {
	ID     uint   `json:"id"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	TextId string `json:"text_id"`
	TextEn string `json:"text_en"`
}
