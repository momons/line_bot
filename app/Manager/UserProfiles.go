package Manager

import (
	"../ApiEntity"
	"../DatabaseEntity"
	"log"
	"time"
)

// ユーザプロフィールマネージャ.
type UserProfiles struct {
}

// 新規ユーザプロフィールマネージャ作成.
func NewUserProfiles() *UserProfiles {
	manager := UserProfiles{}
	return &manager
}

// 1件取得
func (manager *UserProfiles) SelectOne(
	mid string,
) *DatabaseEntity.UserProfiles {

	var entity DatabaseEntity.UserProfiles

	err := DB.Where(
		"mid = ?",
		mid,
	).First(&entity).Error
	if err != nil {
		return nil
	}

	return &entity
}

// 既存チェック.
func (manager *UserProfiles) HasUser(
	mid string,
) bool {
	entity := manager.SelectOne(mid)
	if entity == nil {
		return false
	}
	return !entity.IsLimit()
}

// 追加.
func (manager *UserProfiles) UpdateInsert(
	entity ApiEntity.GetProfile,
) bool {

	// 現在日付取得
	nowAt := time.Now()

	tx := DB.Begin()

	for _, contact := range entity.Contacts {

		// 既存チェック
		getEntity := manager.SelectOne(contact.Mid)
		if getEntity == nil {
			// 作成
			insertEntity := DatabaseEntity.UserProfiles{
				Mid:           contact.Mid,
				DisplayName:   contact.DisplayName,
				PictureUrl:    contact.PictureUrl,
				StatusMessage: contact.StatusMessage,
				UpdateAt:      nowAt,
				CreateAt:      nowAt,
			}
			err := tx.Create(&insertEntity).Error
			if err != nil {
				log.Println(err)
				tx.Rollback()
				return false
			}
		} else {
			getEntity.DisplayName = contact.DisplayName
			getEntity.PictureUrl = contact.PictureUrl
			getEntity.StatusMessage = contact.StatusMessage
			getEntity.UpdateAt = nowAt
			err := tx.Save(getEntity).Error
			if err != nil {
				log.Println(err)
				tx.Rollback()
				return false
			}
		}
	}

	tx.Commit()

	return true
}
