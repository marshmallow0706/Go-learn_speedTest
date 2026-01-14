package main

import (
	"fmt"

	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	// 1. スピードテスト用のクライアントを作成
	client := speedtest.New()

	// 2. 一番近いサーバーを取得
	serverList, _ := client.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		// 3. ダウンロード速度を計測
		err := s.DownloadTest()
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}

		// 4. アップロード速度を計測
		err = s.UploadTest()
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}

		// 結果を表示
		fmt.Printf("サーバー: %s \n", s.Name)
		fmt.Printf("ダウンロード速度: %.2f Mbps\n", s.DLSpeed)
		fmt.Printf("アップロード速度: %.2f Mbps\n", s.ULSpeed)
	}
}
