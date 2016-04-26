package Api

import (
	"../Util"
	"../ApiEntity"
	"../Constants"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
	"bytes"
	"github.com/satori/go.uuid"
)

// テキスト送信イベント.
func (api *ReceiveMessage) ReceiveDiffusionImage(w rest.ResponseWriter, req *rest.Request) {

	// TODO:ゆくゆくはマルチパートに。とりあえず、そのままの画像データ。
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

	// コンテンツ画像にリサイズ
	imageData := Util.GetImageUtil().ResizeContentMetadata(bodyData.Bytes())
	if imageData == nil {
			rest.Error(w, "", http.StatusBadRequest)
			return
	}
	Util.GetFileUtil().SaveFile(messageId, imageData, Constants.ContentTypeImage)
	// プレビュー画像にリサイズ
	previewData := Util.GetImageUtil().ResizePreview(bodyData.Bytes())
	if previewData == nil {
		rest.Error(w, "", http.StatusBadRequest)
		return
	}
	Util.GetFileUtil().SaveFile(messageId + "_pre", previewData, Constants.ContentTypeImage)

	go func() {
		// ユーザ一覧取得
		entities := api.userProfilesManager.SelectList()
		mids := make([]string, len(entities))
		for index, entity := range entities {
			mids[index] = entity.Mid
		}

		if len(mids) > 0 {
			sendImageEntity := ApiEntity.NewSendImage()
			sendImageEntity.OriginalContentUrl = Util.GetFileUtil().GetUrlPath(messageId, Constants.ContentTypeImage)
			sendImageEntity.PreviewImageUrl = Util.GetFileUtil().GetUrlPath(messageId + "_pre", Constants.ContentTypeImage)
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

