package ApiEntity

import "../Constants"

// 動画送信Entity.
type SendVideo struct {
	// コンテンツタイプ(3固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// オリジナル動画URL.
	OriginalContentUrl string `json:"originalContentUrl"`
	// プレビュー動画URL.
	PreviewImageUrl string `json:"previewImageUrl"`
}

// 新規動画送信Entity.
func NewSendVideo() *SendVideo {
	entity := SendVideo{
		ContentType: Constants.ContentTypeVideo,
		ToType:      1,
	}
	return &entity
}
