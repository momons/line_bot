package ApiEntity

import "../Constants"

// ステッカー送信Entity.
type SendSticker struct {
	// コンテンツタイプ(8固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// コンテンツ詳細.
	ContentMetadata SendStickerMetaData `json:"contentMetadata"`
}

// ステッカー送信詳細Entity.
type SendStickerMetaData struct {
	// ステッカーID.
	STKID string `json:"STKID"`
	// ステッカーのパッケージID.
	STKPKGID string `json:"STKPKGID"`
	// ステッカーのバージョン.
	STKVER string `json:"STKVER"`
}

// 新規ステッカー送信Entity.
func NewSendSticker() *SendSticker {
	entity := SendSticker{
		ContentType: Constants.ContentTypeSticker,
		ToType:      1,
	}
	return &entity
}
