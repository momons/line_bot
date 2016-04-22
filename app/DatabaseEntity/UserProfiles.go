package DatabaseEntity

import (
	"../Constants"
	"time"
)

type UserProfiles struct {
	// ID.
	Id int64 `db:"id"`
	// ユーザコード.
	Mid string `db:"mid"`
	// ユーザ名.
	DisplayName string `db:"display_name"`
	// 画像URL.
	PictureUrl string `db:"picture_url"`
	// ステータスメッセージ.
	StatusMessage string `db:"status_message"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}

// 期限チェック.
func (entity *UserProfiles) IsLimit() bool {
	// 期限チェック
	if time.Now().Unix() > entity.UpdateAt.AddDate(0, 0, Constants.UserProfilePeriod).Unix() {
		return true
	}
	return false
}
