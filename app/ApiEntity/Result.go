package ApiEntity

// 結果リクエストEntity.
type Result struct {
	// .
	Id string `json:"id"`
	// .
	From string `json:"from"`
	// .
	FromChannel int64 `json:"fromChannel"`
	// .
	To []string `json:"to"`
	// .
	ToChannel int64 `json:"toChannel"`
	// .
	EventType string `json:"eventType"`
	// .
	Content Content `json:"content"`
}
