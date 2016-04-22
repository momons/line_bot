-- ユーザプロフィール
CREATE TABLE user_profiles (
    -- ID
    id bigserial PRIMARY KEY,
    -- ユーザコード
    mid VARCHAR(64) NOT NULL,
    -- ユーザ名
    display_name TEXT,
    -- 画像URL
    picture_url TEXT,
    -- ステータスメッセージ
    status_message TEXT,
    -- 更新日時
    update_at TIMESTAMP,
    -- 作成日時
    create_at TIMESTAMP
);
