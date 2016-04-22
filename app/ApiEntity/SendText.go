package ApiEntity

import "../Constants"

// テキスト送信Entity.
type SendText struct {
	// コンテンツタイプ (1固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// テキスト.
	Text string `json:"text"`
}

// テキスト送信Entity.
func NewSendText() *SendText {
	entity := SendText{
		ContentType: Constants.ContentTypeText,
		ToType:      1,
	}
	return &entity
}
