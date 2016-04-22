package ApiEntity

// メッセージ受信リクエストEntity.
type ReceivedMessage struct {
	// 結果.
	Result []Result `json:"result"`
}
