package DatabaseEntity

import "time"

type SentMessages struct {
	// ID.
	Id int64 `db:"id"`
	// 送信者ID.
	ToMid string `db:"to_mid"`
	// メッセージID.
	MessageId string `db:"message_id"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}
