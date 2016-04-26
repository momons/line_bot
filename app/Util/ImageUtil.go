package Util

import (
	"bytes"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"log"
)

// 画像ユーティル.
type ImageUtil struct {
}

// 画像ユーティルインスタンス.
var imageUtilInstance *ImageUtil

// インスタンス取得.
func GetImageUtil() *ImageUtil {
	if imageUtilInstance == nil {
		// インスタンス作成.
		imageUtilInstance = &ImageUtil{}
	}
	return imageUtilInstance
}

// プレビュー用に画像リサイズ.
func (util *ImageUtil) ResizePreview(
	imageData []byte,
) []byte {
	return util.ResizeImage(imageData, 240)
}

// メタデータ用に画像リサイズ.
func (util *ImageUtil) ResizeContentMetadata(
	imageData []byte,
) []byte {
	return util.ResizeImage(imageData, 1024)
}

// 画像リサイズ.
func (util *ImageUtil) ResizeImage(
	imageData []byte,
	maxSize uint,
) []byte {

	// 画像変換.
	image, _, err := image.Decode(bytes.NewBuffer(imageData))
	if err != nil {
		log.Println(err)
		return nil
	}

	// サイズの大きいほうを採用
	var width uint = 0
	var height uint = 0
	if image.Bounds().Dx() > image.Bounds().Dy() && image.Bounds().Dx() > int(maxSize) {
		width = maxSize
	} else if image.Bounds().Dx() <= image.Bounds().Dy() && image.Bounds().Dy() > int(maxSize) {
		height = maxSize
	}
	if width == 0 && height == 0 {
		return imageData
	}

	resizedimg := resize.Resize(width, height, image, resize.Lanczos3)

	buff := new(bytes.Buffer)
	err = jpeg.Encode(buff, resizedimg, nil)
	if err != nil {
		log.Println(err)
		return nil
	}

	return buff.Bytes()
}
