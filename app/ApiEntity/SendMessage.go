package ApiEntity

// メッセージ送信Entity.
type SendMessage struct {
	// 送信先ID.
	To []string `json:"to"`
	// 送信先チャンネル.
	ToChannel string `json:"toChannel"`
	// イベントタイプ.
	EventType string `json:"eventType"`
	// コンテンツ.
	Content interface{} `json:"content"`
}
