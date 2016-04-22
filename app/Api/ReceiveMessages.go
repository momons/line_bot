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
	// API
	api *rest.Api
	// ポート
	port int
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
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	receiveMessage.api.SetApp(router)

	return &receiveMessage
}

// メッセージ受信スタート.
func (receiveMessage *ReceiveMessage) Start() {
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(receiveMessage.port), receiveMessage.api.MakeHandler()))
}

// メッセージ受信イベント.
func (receiveMessage *ReceiveMessage) Receive(w rest.ResponseWriter, req *rest.Request) {

	// リクエスト取得
	requestEntity := ApiEntity.ReceivedMessage{}
	err := req.DecodeJsonPayload(&requestEntity)
	if err != nil {
		log.Println(err)
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// ヘッダ情報チェック

	// メタデータ取得
	go receiveMessage.fetchContentMetadata(requestEntity)

	// ユーザ情報取得
	go receiveMessage.fetchUserProfile(requestEntity)

	// 情報をテーブルに保存するのみ
	messagesManagar := Manager.NewMessages()
	messagesManagar.Insert(requestEntity)

	// 即OKを返却
	w.WriteJson(map[string]string{"status": "OK"})
}

// コンテンツメタデータを取得.
func (receiveMessage *ReceiveMessage) fetchContentMetadata(
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
		}
	}
}

// ユーザ情報を取得.
func (receiveMessage *ReceiveMessage) fetchUserProfile(
	requestEntity ApiEntity.ReceivedMessage,
) {
	userProfilesManager := Manager.NewUserProfiles()
	for _, result := range requestEntity.Result {
		hasUser := userProfilesManager.HasUser(result.From)
		if !hasUser {
			log.Println("Fetch user profile.", result.From)
			// ユーザ情報取得
			userProfile := NewUserProfile(result.From)
			userProfile.Send()
			if userProfile.ResponseHttpStatus == http.StatusOK {
				log.Println("Fetch user profile success.", result.From)
				entity := ApiEntity.NewGetProfile(userProfile.ResponseBody)
				if entity != nil {
					// 保存
					userProfilesManager.UpdateInsert(*entity)
				}
			} else {
				log.Println("Fetch user profile failure.", result.From)
			}
		}
	}
}
