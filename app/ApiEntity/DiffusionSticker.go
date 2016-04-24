package ApiEntity

// ステッカー拡散Entity.
type DiffusionSticker struct {
	// アプリトークン.
	AppToken string `json:"appToken"`
	// ステッカーID.
	STKID string `json:"STKID"`
	// ステッカーのパッケージID.
	STKPKGID string `json:"STKPKGID"`
	// ステッカーのバージョン.
	STKVER string `json:"STKVER"`
}
