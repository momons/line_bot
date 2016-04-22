package DatabaseEntity

import "time"

// メッセージテーブルEntity.
type Messages struct {
	// ID.
	Id int64 `db:"id"`
	// メッセージID.
	MessageId string `db:"message_id"`
	// コンテンツタイプ.
	ContentType int `db:"content_type"`
	// 送信者MID.
	FromMid string `db:"from_mid"`
	// 作成日.
	CreateTime int64 `db:"create_time"`
	// 受信者ID群.
	ToMids string `db:"to_mids"`
	// メッセージID.
	ToType int `db:"to_type"`
	// メッセージ詳細情報.
	ContentMetadata string `db:"content_metadata"`
	// 送信テキスト.
	Text string `db:"text"`
	// ロケーションデータ.
	Location string `db:"location"`
	// 返信済み.
	Replied int `db:"replied"`
	// 更新日時.
	UpdateAt time.Time `db:"update_at"`
	// 作成日時.
	CreateAt time.Time `db:"create_at"`
}
