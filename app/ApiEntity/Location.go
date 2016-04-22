package ApiEntity

import (
	"bytes"
	"encoding/json"
	"log"
)

// 位置情報Entity.
type Location struct {
	// タイトル.
	Title string `json:"title"`
	// 住所.
	Address string `json:"address"`
	// 緯度.
	Latitude float64 `json:"latitude"`
	// 経度.
	Longitude float64 `json:"longitude"`
}

// 新規位置情報.
func NewLocation(
	jsonStr string,
) *Location {
	var entity Location
	dec := json.NewDecoder(bytes.NewBuffer([]byte(jsonStr)))
	err := dec.Decode(&entity)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &entity
}
