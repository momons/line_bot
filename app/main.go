package main

import (
	"./Api"
	"./Manager"
	"./Service"
	"flag"
	"math/rand"
	"os"
	"time"
)

// 終了コード.
var exitCode = 0

// ポート.
var apiPort int

func main() {

	// セットアップ
	isSuccess := setup()
	if !isSuccess {
		exitCode = 1
	}

	os.Exit(exitCode)
}

// セットアップ.
func setup() bool {

	// コマンドラインセットアップ
	setupCommandLine()

	// 乱数初期化
	rand.Seed(time.Now().UnixNano())

	// DBセットアップ
	isSuccess := Manager.Setup()
	if !isSuccess {
		return false
	}

	// 送信ポーリング 別スレッド
	sendMessage := Service.NewSendMessage()
	go sendMessage.Run()

	// API
	receiveMessage := Api.NewReceiveMessage(apiPort)
	if receiveMessage == nil {
		return false
	}
	receiveMessage.Start()

	return true
}

// コマンドラインセットアップ
func setupCommandLine() {

	// ポート
	flag.IntVar(&apiPort, "port", 9002, "ポートを指定します。")

	flag.Parse()
}
