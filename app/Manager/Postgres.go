package Manager

import (
	"../Constants"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

// データベース.
var DB *gorm.DB

// セットアップ.
func Setup() bool {

	// 接続文字列
	connectStr := "sslmode=disable host=localhost" + " dbname=" + Constants.DatabaseName + " user=" + Constants.UserId + " password=" + Constants.Password
	db, err := gorm.Open("postgres", connectStr)
	if err != nil {
		log.Println(err)
		return false
	}
	db.DB()

	// 退避
	DB = db

	return true
}
