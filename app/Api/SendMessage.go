package Api

import (
	"../Constants"
	"encoding/json"
)

// メッセージ送信API.
type SendMessage struct {
	*HTTPClient
}

// 新規取得.
func NewSendMessage(postData interface{}) *SendMessage {
	sendMessage := SendMessage{}
	sendMessage.HTTPClient = NewClient()
	sendMessage.RequestUrl = "https://trialbot-api.line.me/v1/events"
	sendMessage.RequestHeader = map[string]string{
		"Content-Type":                 "application/json; charset=utf-8",
		"X-Line-ChannelID":             Constants.ChannelId,
		"X-Line-ChannelSecret":         Constants.ChannelSecret,
		"X-Line-Trusted-User-With-ACL": Constants.MID,
	}
	sendMessage.RequestMethod = "POST"
	sendMessage.RequestBody, _ = json.Marshal(postData)
	return &sendMessage
}
