package Manager

import (
	"../ApiEntity"
	"../Constants"
	"../DatabaseEntity"
	"encoding/json"
	"log"
	"time"
)

// メッセージマネージャ.
type Messages struct {
}

// 新規メッセージマネージャ作成.
func NewMessages() *Messages {
	manager := Messages{}
	return &manager
}

// 未処理の情報を取得
func (manager *Messages) SelectUnsentList(
	mid string,
) *[]DatabaseEntity.Messages {

	var entities []DatabaseEntity.Messages

	err := DB.Where(
		"message_id NOT IN ( SELECT message_id FROM sent_messages WHERE to_mid = ? ) AND from_mid != ?",
		mid,
		mid,
		//"message_id NOT IN ( SELECT message_id FROM sent_messages WHERE to_mid = ? )",
		//mid,
	).Find(&entities).Error
	if err != nil {
		log.Println(err)
		return nil
	}

	return &entities
}

// 未返信のデータを取得、送信済みに更新
func (manager *Messages) SelectUpdateNotReply() (bool, *[]DatabaseEntity.Messages) {

	// 現在日付取得
	nowAt := time.Now()

	tx := DB.Begin()

	var entities []DatabaseEntity.Messages

	err := tx.Where(
		"replied = ?",
		Constants.StatusTypeNotReply,
	).Order("update_at DESC").Find(&entities).Error
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return false, nil
	}

	// 返信済みに更新
	for _, message := range entities {
		message.Replied = Constants.StatusTypeReplied
		message.UpdateAt = nowAt
		err := tx.Save(&message).Error
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return false, nil
		}
	}

	tx.Commit()

	return true, &entities
}

// 追加.
func (manager *Messages) Insert(
	requestEntity ApiEntity.ReceivedMessage,
) bool {

	// 現在日付取得
	nowAt := time.Now()

	tx := DB.Begin()

	for _, result := range requestEntity.Result {

		toMidsStr, _ := json.Marshal(result.Content.To)
		metaDataStr, _ := json.Marshal(result.Content.ContentMetadata)
		locationStr, _ := json.Marshal(result.Content.Location)

		insertEntity := DatabaseEntity.Messages{
			MessageId:       result.Content.Id,
			ContentType:     result.Content.ContentType,
			FromMid:         result.Content.From,
			CreateTime:      result.Content.CreatedTime,
			ToMids:          string(toMidsStr),
			ToType:          result.Content.ToType,
			ContentMetadata: string(metaDataStr),
			Text:            result.Content.Text,
			Location:        string(locationStr),
			Replied:         0,
			UpdateAt:        nowAt,
			CreateAt:        nowAt,
		}
		err := tx.Create(&insertEntity).Error
		if err != nil {
			log.Println(err)
			tx.Rollback()
			return false
		}
	}

	tx.Commit()

	return true
}
