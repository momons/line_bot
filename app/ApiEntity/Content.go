package ApiEntity

// Property list of messages.
type Content struct {
	// メッセージID.
	Id string `json:"id"`
	// コンテンツタイプ.
	ContentType int `json:"contentType"`
	// 送信者MID.
	From string `json:"from"`
	// 作成日 UNIXタイム.
	CreatedTime int64 `json:"createdTime"`
	// 受信者ID群.
	To []string `json:"to"`
	// 受信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// メッセージ詳細情報.
	ContentMetadata interface{} `json:"contentMetadata"`
	// 送信テキスト (最大10000文字).
	Text string `json:"text"`
	// ロケーションデータ (ContentTypeが7の場合のみ).
	Location *Location `json:"location"`
}
