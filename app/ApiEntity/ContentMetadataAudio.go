package ApiEntity

import (
	"encoding/json"
	"bytes"
	"log"
)

// 音声メタデータ.
type ContentMetadataAudio struct {
	// ？.
	AT_RECV_MODE string `json:"AT_RECV_MODE"`
	// 音声再生長.
	AUDLEN string `json:"AUDLEN"`
	// ？.
	OBS_POP string `json:"OBS_POP"`
	// ？.
	SKIP_BADGE_COUNT string `json:"SKIP_BADGE_COUNT"`
}

// 新規音声メタデータ.
func NewContentMetadataAudio(
	jsonStr string,
) *ContentMetadataAudio {
	var metadata ContentMetadataAudio
	dec := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	err := dec.Decode(&metadata)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &metadata
}
