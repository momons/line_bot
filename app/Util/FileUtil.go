package Util

import (
	"../Constants"
	"io/ioutil"
	"log"
	"os"
)

// ファイルユーティル.
type FileUtil struct {
}

// ファイルユーティルインスタンス.
var fileUtilInstance *FileUtil

// インスタンス取得.
func GetFileUtil() *FileUtil {
	if fileUtilInstance == nil {
		// インスタンス作成.
		fileUtilInstance = &FileUtil{}
	}
	return fileUtilInstance
}

// データ読み込み.
func (util *FileUtil) LoadFile(messageId string, contentType int) *[]byte {
	// ファイル読み込み
	data, err := ioutil.ReadFile(util.getFilePath(messageId, contentType))
	if err != nil {
		log.Println(err)
		return nil
	}
	return &data
}

// データ保存.
func (util *FileUtil) SaveFile(messageId string, data []byte, contentType int) bool {
	// ファイル書き込み
	err := ioutil.WriteFile(util.getFilePath(messageId, contentType), data, os.ModePerm)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// ファイルパスを取得.
func (util *FileUtil) getFilePath(
	messageId string,
	contentType int,
) string {
	fileName := GetHashUtil().Hash(Constants.HashHeader, messageId, Constants.HashFooter)
	return util.getDir(contentType) + fileName + util.getExtension(contentType)
}

// URLパスを取得.
func (util *FileUtil) GetUrlPath(
	messageId string,
	contentType int,
) string {
	fileName := GetHashUtil().Hash(Constants.HashHeader, messageId, Constants.HashFooter)
	return util.getPath(contentType) + fileName + util.getExtension(contentType)
}

// コンテンツタイプよりディレクトリ取得.
func (util *FileUtil) getDir(
	contentType int,
) string {
	switch contentType {
	case Constants.ContentTypeImage:
		return Constants.DirImage
	case Constants.ContentTypeVideo:
		return Constants.DirImage
	case Constants.ContentTypeAudio:
		return Constants.DirAudio
	}
	return ""
}

// コンテンツタイプよりURL取得.
func (util *FileUtil) getPath(
	contentType int,
) string {
	switch contentType {
	case Constants.ContentTypeImage:
		return Constants.UrlImage
	case Constants.ContentTypeVideo:
		return Constants.UrlImage
	case Constants.ContentTypeAudio:
		return Constants.UrlAudio
	}
	return ""
}

func (util *FileUtil) getExtension(
	contentType int,
) string {
	switch contentType {
	case Constants.ContentTypeImage:
		return ".jpg"
	case Constants.ContentTypeVideo:
		return ".mp4"
	case Constants.ContentTypeAudio:
		return ".m4a"
	}
	return ""
}

// ディレクトリを作成.
func (util *FileUtil) CreateDir(
	dirPath string,
) bool {
	// ファイル情報取得
	fileInfo, err := os.Stat(dirPath)
	if err != nil || !fileInfo.IsDir() {
		// ディレクトリでなかったら作成
		err := os.Mkdir(dirPath, 0777)
		if err != nil {
			log.Println(err)
			return false
		}
	}
	return true
}
