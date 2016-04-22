package Api

import (
	"../Constants"
)

// ユーザプロフィール取得API.
type UserProfile struct {
	*HTTPClient
}

// 新規取得.
func NewUserProfile(mid string) *UserProfile {
	// 値を初期化
	userProfile := UserProfile{}
	userProfile.HTTPClient = NewClient()
	userProfile.RequestUrl = "https://trialbot-api.line.me/v1/profiles?mids=" + mid
	userProfile.RequestHeader = map[string]string{
		"X-Line-ChannelID":             Constants.ChannelId,
		"X-Line-ChannelSecret":         Constants.ChannelSecret,
		"X-Line-Trusted-User-With-ACL": Constants.MID,
	}
	userProfile.RequestMethod = "GET"
	return &userProfile
}
