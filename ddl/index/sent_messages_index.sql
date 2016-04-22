-- 送信済みメッセージテーブル ユニークインデックス
CREATE UNIQUE INDEX sent_messages_index ON sent_messages (
    to_mid,
    message_id
);
