-- メッセージ
CREATE TABLE messages (
    -- ID
    id bigserial PRIMARY KEY,
    -- メッセージID
    message_id VARCHAR(64) NOT NULL,
    -- コンテンツタイプ
    content_type INTEGER,
    -- 送信者MID
    from_mid VARCHAR(64),
    -- 作成日
    create_time BIGINT,
    -- 受信者ID群.
    to_mids TEXT,
    -- 受信者タイプ
    to_type INTEGER,
    -- メッセージ詳細情報.
    content_metadata TEXT,
    -- 送信テキスト
    text TEXT,
    -- ロケーションデータ
    location TEXT,
    -- 送信
    replied INTEGER,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
