package ApiEntity

import "../Constants"

// 画像送信Entity.
type SendImage struct {
	// コンテンツタイプ(2固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// オリジナル画像URL 1024×1024.
	OriginalContentUrl string `json:"originalContentUrl"`
	// プレビュー画像URL 240×240.
	PreviewImageUrl string `json:"previewImageUrl"`
}

// 新規画像送信Entity.
func NewSendImage() *SendImage {
	entity := SendImage{
		ContentType: Constants.ContentTypeImage,
		ToType:      1,
	}
	return &entity
}
