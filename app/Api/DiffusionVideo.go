package Api

import (
	"../Util"
	"../ApiEntity"
	"../Constants"
	"github.com/ant0ine/go-json-rest/rest"
	"bytes"
	"github.com/satori/go.uuid"
)

// テキスト送信イベント.
func (api *ReceiveMessage) ReceiveDiffusionVideo(w rest.ResponseWriter, req *rest.Request) {

	// TODO:ゆくゆくはマルチパートに。とりあえず、そのままの動画データ。
	// トークンチェック
	//if requestEntity.AppToken != Constants.AppToken {
	//	rest.Error(w, "", http.StatusBadRequest)
	//	return
	//}

	// リクエスト取得 (データ=画像データ)
	bodyData := new(bytes.Buffer)
	bodyData.ReadFrom(req.Body)

	// 擬似メッセージIDを生成
	messageId := uuid.NewV4().String()

	// コンテンツ動画保存
	Util.GetFileUtil().SaveFile(messageId, bodyData.Bytes(), Constants.ContentTypeVideo)

	go func() {
		// ユーザ一覧取得
		entities := api.userProfilesManager.SelectList()
		mids := make([]string, len(entities))
		for index, entity := range entities {
			mids[index] = entity.Mid
		}

		if len(mids) > 0 {
			sendImageEntity := ApiEntity.NewSendVideo()
			sendImageEntity.OriginalContentUrl = Util.GetFileUtil().GetUrlPath(messageId, Constants.ContentTypeVideo)
			sendImageEntity.PreviewImageUrl = Constants.UrlMetadata + "video.jpg"
			sendMessageEntity := ApiEntity.SendMessage{
				To:        mids,
				ToChannel: Constants.SendMessageChannelId,
				EventType: Constants.SendMessageEventType,
				Content:   sendImageEntity,
			}

			// 送信
			sendMessage := NewSendMessage(sendMessageEntity)
			sendMessage.Send()
		}
	}()

	// 即OKを返却
	w.WriteJson(map[string]string{"status": "OK"})
}

