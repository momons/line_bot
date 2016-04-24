package ApiEntity

import (
	"bytes"
	"encoding/json"
	"log"
)

// ステッカーメタデータ.
type ContentMetadataSticker struct {
	// ？.
	AT_RECV_MODE string `json:"AT_RECV_MODE"`
	// ？.
	SKIP_BADGE_COUNT string `json:"SKIP_BADGE_COUNT"`
	// ステッカーID.
	STKID string `json:"STKID"`
	// ？.
	STKOPT string `json:"STKOPT"`
	// パッケージID.
	STKPKGID string `json:"STKPKGID"`
	// ？.
	STKTXT string `json:"STKTXT"`
	// ステッカーバージョン.
	STKVER string `json:"STKVER"`
}

// 新規ステッカーメタデータ.
func NewContentMetadataSticker(
	jsonStr string,
) *ContentMetadataSticker {
	var metadata ContentMetadataSticker
	dec := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	err := dec.Decode(&metadata)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &metadata
}
