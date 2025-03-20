package dto

type ResponseMessageDto struct {
	ID        uint   `json:"id"`
	Code      string `json:"code"`
	TitleId   string `json:"title_id"`
	TitleEn   string `json:"title_en"`
	MessageId string `json:"message_id"`
	MessageEn string `json:"message_en"`
	Source    string `json:"source"`
}
