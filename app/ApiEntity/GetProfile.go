package ApiEntity

import (
	"bytes"
	"encoding/json"
	"log"
)

// ユーザプロフィールEntity.
type GetProfile struct {
	// プロフィール.
	Contacts []GetProfileContacts `json:"contacts"`
	// カウント.
	Count int `json:"count"`
	// 合計.
	Total int `json:"total"`
	// 開始インデックス.
	Start int `json:"start"`
	// 表示パラメータ値.
	Display int `json:"display"`
}

// ユーザプロフィールコンタクトEntity.
type GetProfileContacts struct {
	// ニックネーム.
	DisplayName string `json:"displayName"`
	// ユーザID.
	Mid string `json:"mid"`
	// 画像URL.
	PictureUrl string `json:"pictureUrl"`
	// ステータスメッセージ.
	StatusMessage string `json:"statusMessage"`
}

// バイナリデータからEntity作成.
func NewGetProfile(data []byte) *GetProfile {
	dec := json.NewDecoder(bytes.NewReader(data))
	var entity GetProfile
	err := dec.Decode(&entity)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &entity
}
