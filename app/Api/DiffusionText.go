package Api

import (
	"log"
	"../ApiEntity"
	"../Constants"
	"github.com/ant0ine/go-json-rest/rest"
	"net/http"
)

// テキスト送信イベント.
func (api *ReceiveMessage) ReceiveDiffusionText(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	requestEntity := ApiEntity.DiffusionText{}
	err := req.DecodeJsonPayload(&requestEntity)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// トークンチェック
	if requestEntity.AppToken != Constants.AppToken {
		rest.Error(w, "", http.StatusBadRequest)
		return
	}

	go func() {
		// ユーザ一覧取得
		entities := api.userProfilesManager.SelectList()
		mids := make([]string, len(entities))
		for index, entity := range entities {
			mids[index] = entity.Mid
		}

		if len(mids) > 0 {
			sendTextEntity := ApiEntity.NewSendText()
			sendTextEntity.Text = requestEntity.Text
			sendMessageEntity := ApiEntity.SendMessage{
				To:        mids,
				ToChannel: Constants.SendMessageChannelId,
				EventType: Constants.SendMessageEventType,
				Content:   sendTextEntity,
			}

			// 送信
			sendMessage := NewSendMessage(sendMessageEntity)
			sendMessage.Send()
		}
	}()

	// 即OKを返却
	w.WriteJson(map[string]string{"status": "OK"})
}

