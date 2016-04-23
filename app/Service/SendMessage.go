package Service

import (
	"../Api"
	"../ApiEntity"
	"../Constants"
	"../DatabaseEntity"
	"../Manager"
	"../Util"
	"math/rand"
	"time"
)

// メッセージ送信サービス
type SendMessage struct {
	// メッセージマネージャ.
	messagesManager *Manager.Messages
	// 送信済みマネージャ.
	sentMessages *Manager.SentMessages
}

// 新規メッセージ送信サービス.
func NewSendMessage() *SendMessage {
	service := SendMessage{
		messagesManager: Manager.NewMessages(),
		sentMessages:    Manager.NewSentMessages(),
	}
	return &service
}

// 実行.
func (service *SendMessage) Run() {

	for {
		// 未送信メッセージチェック
		isSuccess, notReplyMessages := service.messagesManager.SelectUpdateNotReply()
		if isSuccess {
			// ユーザIDでさまり
			mids := service.summaryMids(notReplyMessages)
			for mid, _ := range mids {

				requestEntity := ApiEntity.SendMessage{
					To:        []string{mid},
					ToChannel: Constants.SendMessageChannelId,
					EventType: Constants.SendMessageEventType,
				}

				// メッセージ取得
				message := service.getSendMessage(mid)
				if message == nil {
					// 送信するメッセージがない。
					requestEntity.Content = service.createSendText("...返事がない...しかばねのようだ...")
				} else {
					// 処理追加
					service.sentMessages.Insert(mid, []string{message.MessageId})

					// テキスト
					switch message.ContentType {
					case Constants.ContentTypeText:
						requestEntity.Content = service.createSendText(message.Text)
					case Constants.ContentTypeImage:
						requestEntity.Content = service.createSendImage(message)
					case Constants.ContentTypeVideo:
						requestEntity.Content = service.createSendVideo(message)
					case Constants.ContentTypeAudio:
						requestEntity.Content = service.createSendAudio(message)
					case Constants.ContentTypeLocation:
						requestEntity.Content = service.createSendLocation(message)
					case Constants.ContentTypeSticker:
						requestEntity.Content = service.createSendSticker(message)
					default:
						requestEntity.Content = service.createSendText("...返事がない...しかばねのようだ...")
					}
				}

				// 送信済みとして処理
				go func() {
					sendMessage := Api.NewSendMessage(requestEntity)
					sendMessage.Send()
				}()
			}
		}
		// ウエイト
		time.Sleep(Constants.SendMessagePollingWaitTime * time.Second)
	}
}

// ユーザIDでサマリ.
func (service *SendMessage) summaryMids(
	messages *[]DatabaseEntity.Messages,
) map[string]string {

	mids := map[string]string{}

	for _, message := range *messages {
		mids[message.FromMid] = message.FromMid
	}

	return mids
}

// 送信メッセージを1つ取得.
func (service *SendMessage) getSendMessage(
	mid string,
) *DatabaseEntity.Messages {

	// 対象のメッセージ
	messages := service.messagesManager.SelectUnsentList(mid)
	if messages == nil || len(*messages) <= 0 {
		return nil
	}

	// ランダムで抽出
	index := rand.Int() % len(*messages)

	return &(*messages)[index]
}

// 送信テキスト作成
func (service *SendMessage) createSendText(
	text string,
) *ApiEntity.SendText {
	sendEntity := ApiEntity.NewSendText()
	sendEntity.Text = text
	return sendEntity
}

// 送信画像作成
func (service *SendMessage) createSendImage(
	message *DatabaseEntity.Messages,
) *ApiEntity.SendImage {
	sendEntity := ApiEntity.NewSendImage()
	sendEntity.OriginalContentUrl = Util.GetFileUtil().GetUrlPath(message.MessageId, Constants.ContentTypeImage)
	sendEntity.PreviewImageUrl = Util.GetFileUtil().GetUrlPath(message.MessageId+"_pre", Constants.ContentTypeImage)
	return sendEntity
}

// 送信動画作成
func (service *SendMessage) createSendVideo(
	message *DatabaseEntity.Messages,
) *ApiEntity.SendVideo {
	sendEntity := ApiEntity.NewSendVideo()
	sendEntity.OriginalContentUrl = Util.GetFileUtil().GetUrlPath(message.MessageId, Constants.ContentTypeVideo)
	sendEntity.PreviewImageUrl = Constants.UrlMetadata + "video.jpg"
	return sendEntity
}

// 送信音声作成
func (service *SendMessage) createSendAudio(
	message *DatabaseEntity.Messages,
) *ApiEntity.SendAudio {
	sendEntity := ApiEntity.NewSendAudio()
	sendEntity.OriginalContentUrl = Util.GetFileUtil().GetUrlPath(message.MessageId, Constants.ContentTypeAudio)
	// 音声情報取得
	metadata := ApiEntity.NewContentMetadataAudio(message.ContentMetadata)
	if metadata != nil {
		sendEntity.ContentMetadata.AUDLEN = metadata.AUDLEN
	}
	return sendEntity
}

// 送信位置情報作成
func (service *SendMessage) createSendLocation(
	message *DatabaseEntity.Messages,
) *ApiEntity.SendLocation {
	sendEntity := ApiEntity.NewSendLocation()
	// 位置情報取得
	location := ApiEntity.NewLocation(message.Location)
	if location != nil {
		sendEntity.Text = location.Address
		sendEntity.Location.Title = location.Address
		sendEntity.Location.Latitude = location.Latitude
		sendEntity.Location.Longitude = location.Longitude
	}
	return sendEntity
}

// 送信ステッカー情報作成
func (service *SendMessage) createSendSticker(
	message *DatabaseEntity.Messages,
) *ApiEntity.SendSticker {
	sendEntity := ApiEntity.NewSendSticker()
	// スタンプ情報取得
	metadata := ApiEntity.NewContentMetadataSticker(message.ContentMetadata)
	if metadata != nil {
		sendEntity.ContentMetadata.STKID = metadata.STKID
		sendEntity.ContentMetadata.STKPKGID = metadata.STKPKGID
		sendEntity.ContentMetadata.STKVER = metadata.STKVER
	}
	return sendEntity
}
