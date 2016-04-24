package Api

import (
	"../ApiEntity"
	"../Constants"
	"../Manager"
	"../Util"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"strconv"
)

// メッセージ受信API.
type ReceiveMessage struct {
	// API.
	api *rest.Api
	// ポート.
	port int
	// ユーザプロフィールマネージャ.
	userProfilesManager *Manager.UserProfiles
	// メッセージマネージャ.
	messageManager *Manager.Messages
}

// 新規取得.
func NewReceiveMessage(port int) *ReceiveMessage {
	receiveMessage := ReceiveMessage{}
	receiveMessage.api = rest.NewApi()
	receiveMessage.port = port
	// 開発スタックモード
	receiveMessage.api.Use(
		//rest.DefaultDevStack...
		[]rest.Middleware{
			&rest.AccessLogApacheMiddleware{},
			&rest.TimerMiddleware{},
			&rest.RecorderMiddleware{},
			&rest.PoweredByMiddleware{},
			&rest.RecoverMiddleware{
				EnableResponseStackTrace: true,
			},
			&rest.JsonIndentMiddleware{},
			//&rest.ContentTypeCheckerMiddleware{},
		}...,
	)
	// ルーティング設定
	router, err := rest.MakeRouter(
		&rest.Route{"POST", "/", receiveMessage.Receive},
		&rest.Route{"POST", "/_send_text", receiveMessage.ReceiveDiffusionText},
		&rest.Route{"POST", "/_send_sticker", receiveMessage.ReceiveDiffusionSticker},
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	receiveMessage.api.SetApp(router)

	receiveMessage.userProfilesManager = Manager.NewUserProfiles()
	receiveMessage.messageManager = Manager.NewMessages()

	return &receiveMessage
}

// メッセージ受信スタート.
func (api *ReceiveMessage) Start() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(api.port), api.api.MakeHandler()))
}

// メッセージ受信イベント.
func (api *ReceiveMessage) Receive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	requestEntity := ApiEntity.ReceivedMessage{}
	err := req.DecodeJsonPayload(&requestEntity)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ヘッダ情報チェック
	if !api.isValidRequest(requestEntity) {
		rest.Error(w, "", http.StatusBadRequest)
		return
	}

	// メタデータ取得
	go api.fetchContentMetadata(requestEntity)

	// ユーザ情報取得
	go api.fetchUserProfile(requestEntity)

	// 情報をテーブルに保存するのみ
	api.messageManager.Insert(requestEntity)

	// 即OKを返却
	w.WriteJson(map[string]string{"status": "OK"})
}

// ヘッダ情報チェック
func (api *ReceiveMessage) isValidRequest(
	requestEntity ApiEntity.ReceivedMessage,
) bool {

	// to_midを判定、１つでも違う宛先が入っていたらダメ。
	for _, result := range requestEntity.Result {
		if len(result.Content.To) <= 0 || result.Content.To[0] != Constants.MID {
			return false
		}
	}

	return true
}

// コンテンツメタデータを取得.
func (api *ReceiveMessage) fetchContentMetadata(
	requestEntity ApiEntity.ReceivedMessage,
) {
	for _, result := range requestEntity.Result {
		if result.Content.ContentType == Constants.ContentTypeImage ||
			result.Content.ContentType == Constants.ContentTypeVideo ||
			result.Content.ContentType == Constants.ContentTypeAudio {
			log.Println("Fetch content metadata.", result.Content.Id)
			getMessageContent := NewGetMessageContent(result.Content.Id)
			getMessageContent.Send()
			if getMessageContent.ResponseHttpStatus == http.StatusOK {
				// メタデータを保存.
				Util.GetFileUtil().SaveFile(result.Content.Id, getMessageContent.ResponseBody, result.Content.ContentType)
				if result.Content.ContentType == Constants.ContentTypeImage {
					// プレビュー用データを取得.
					imageData := Util.GetImageUtil().ResizePreview(getMessageContent.ResponseBody)
					if imageData != nil {
						Util.GetFileUtil().SaveFile(result.Content.Id+"_pre", imageData, result.Content.ContentType)
					}
				}
				log.Println("Fetch content metadata success.", result.Content.Id)
			} else {
				log.Println("Fetch content metadata failure.", result.Content.Id)
			}
		} else if result.Content.ContentType == Constants.ContentTypeContact {
			metadata := ApiEntity.NewContentMetadataContact(result.Content.ContentMetadata)
			if metadata != nil {
				// 保存
				api.userProfilesManager.UpdateInsertForMetadata(*metadata)
			}
		}
	}
}

// ユーザ情報を取得.
func (api *ReceiveMessage) fetchUserProfile(
	requestEntity ApiEntity.ReceivedMessage,
) {
	for _, result := range requestEntity.Result {
		if result.Content.ContentType == Constants.ContentTypeContact {
			continue
		}
		hasUser := api.userProfilesManager.HasUser(result.Content.From)
		if !hasUser {
			log.Println("Fetch user profile.", result.Content.From)
			// ユーザ情報取得
			userProfile := NewUserProfile(result.Content.From)
			userProfile.Send()
			if userProfile.ResponseHttpStatus == http.StatusOK {
				log.Println("Fetch user profile success.", result.Content.From)
				entity := ApiEntity.NewGetProfile(userProfile.ResponseBody)
				if entity != nil {
					// 保存
					api.userProfilesManager.UpdateInsertForGetProfile(*entity)
				}
			} else {
				log.Println("Fetch user profile failure.", result.Content.From)
			}
		}
	}
}

