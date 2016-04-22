-- 送信済みメッセージ
CREATE TABLE sent_messages (
    -- ID
    id bigserial PRIMARY KEY,
    -- 送信者ID
    to_mid VARCHAR(64) NOT NULL,
    -- メッセージID
    message_id VARCHAR(64) NOT NULL,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
