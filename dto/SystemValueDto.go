package dto

type SystemValueDto struct {
	ID        uint   `json:"id"`
	Module    string `json:"module"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	IsEncrypt bool   `json:"is_encrypt"`
}
