package ApiEntity

// テキスト拡散Entity.
type DiffusionText struct {
	// アプリトークン.
	AppToken string `json:"appToken"`
	// テキスト.
	Text string `json:"text"`
}

