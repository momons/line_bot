package ApiEntity

import (
	"bytes"
	"encoding/json"
	"log"
)

// コンタクトメタデータ.
type ContentMetadataContact struct {
	// ユーザID.
	Mid string `json:"mid"`
	// ユーザ名.
	DisplayName string `json:"displayName"`
}

// 新規ステッカーメタデータ.
func NewContentMetadataContact(
	jsonObject interface{},
) *ContentMetadataContact {
	// 一度文字列に変換.
	jsonStr, _ := json.Marshal(jsonObject)
	// 再度オブジェクト化.
	var metadata ContentMetadataContact
	dec := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	err := dec.Decode(&metadata)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &metadata
}
