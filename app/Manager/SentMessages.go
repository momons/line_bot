package Manager

import (
	"../DatabaseEntity"
	"log"
	"time"
)

// 送信済みメッセージマネージャ.
type SentMessages struct {
}

// 新規メッセージマネージャ作成.
func NewSentMessages() *SentMessages {
	manager := SentMessages{}
	return &manager
}

// 1件取得.
func (manager *SentMessages) SelectOne(
	mid string,
	messageId string,
) *DatabaseEntity.SentMessages {

	var entity DatabaseEntity.SentMessages

	err := DB.Where(
		"to_mid = ? AND message_id = ?",
		mid,
		messageId,
	).First(&entity).Error
	if err != nil {
		return nil
	}

	return &entity
}

// 送信済みチェック.
func (manager *SentMessages) isSent(
	mid string,
	messageId string,
) bool {
	return manager.SelectOne(mid, messageId) != nil
}

// 追加.
func (manager *SentMessages) Insert(
	mid string,
	messageIds []string,
) bool {

	// 現在日付取得
	nowAt := time.Now()

	tx := DB.Begin()

	for _, messageId := range messageIds {

		// 送信済みチェック
		if manager.isSent(mid, messageId) {
			continue
		}

		insertEntity := DatabaseEntity.SentMessages{
			ToMid:     mid,
			MessageId: messageId,
			UpdateAt:  nowAt,
			CreateAt:  nowAt,
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
