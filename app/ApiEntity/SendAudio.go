package ApiEntity

import "../Constants"

// 音声送信Entity.
type SendAudio struct {
	// コンテンツタイプ(4固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// オリジナル音声URL.
	OriginalContentUrl string `json:"originalContentUrl"`
	// コンテンツ詳細データ.
	ContentMetadata SendAudioMetaData `json:"contentMetadata"`
}

// 音声送信詳細Entity.
type SendAudioMetaData struct {
	// 音声時間(ms).
	AUDLEN string `json:"AUDLEN"`
}

// 新規音声送信Entity.
func NewSendAudio() *SendAudio {
	entity := SendAudio{
		ContentType: Constants.ContentTypeAudio,
		ToType:      1,
	}
	return &entity
}
