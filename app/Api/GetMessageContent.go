package Api

import "../Constants"

// メッセージコンテンツ取得API.
type GetMessageContent struct {
	*HTTPClient
}

// 新規取得.
func NewGetMessageContent(messageId string) *GetMessageContent {
	// 値を初期化
	getMessageContent := GetMessageContent{}
	getMessageContent.HTTPClient = NewClient()
	getMessageContent.RequestUrl = "https://trialbot-api.line.me/v1/bot/message/" + messageId + "/content"
	getMessageContent.RequestHeader = map[string]string{
		"X-Line-ChannelID":             Constants.ChannelId,
		"X-Line-ChannelSecret":         Constants.ChannelSecret,
		"X-Line-Trusted-User-With-ACL": Constants.MID,
	}
	getMessageContent.RequestMethod = "GET"
	return &getMessageContent
}
