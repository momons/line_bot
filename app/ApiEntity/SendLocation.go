package ApiEntity

import "../Constants"

// 位置情報送信Entity.
type SendLocation struct {
	// コンテンツタイプ(7固定).
	ContentType int `json:"contentType"`
	// 送信者タイプ (1:ユーザ).
	ToType int `json:"toType"`
	// 位置説明.
	Text string `json:"text"`
	// 位置情報詳細.
	Location SendLocationLocation `json:"location"`
}

// 位置情報送信詳細Entity.
type SendLocationLocation struct {
	// Textと同じ値を設定する.
	Title string `json:"title"`
	// 緯度.
	Latitude float64 `json:"latitude"`
	// 経度.
	Longitude float64 `json:"longitude"`
}

// 新規位置情報送信Entity.
func NewSendLocation() *SendLocation {
	entity := SendLocation{
		ContentType: Constants.ContentTypeLocation,
		ToType:      1,
	}
	return &entity
}
