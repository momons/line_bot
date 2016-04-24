package Api

import "../Constants"

// メッセージコンテンツプレビュー取得API.
type GetMessageContentPreview struct {
	*HTTPClient
}

// 新規取得.
func NewGetMessageContentPreview(messageId string) *GetMessageContent {
	// 値を初期化
	getMessageContent := GetMessageContent{}
	getMessageContent.HTTPClient = NewClient()
	getMessageContent.RequestUrl = "https://trialbot-api.line.me/v1/bot/message/" + messageId + "/content/preview"
	getMessageContent.RequestHeader = map[string]string{
		"X-Line-ChannelID":             Constants.ChannelId,
		"X-Line-ChannelSecret":         Constants.ChannelSecret,
		"X-Line-Trusted-User-With-ACL": Constants.MID,
	}
	getMessageContent.RequestMethod = "GET"
	return &getMessageContent
}
