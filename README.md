# ようこそ、Go言語で作るスピードテストツールへ！
Goはシンプルで高速な処理が得意な言語なので、ネットワークツールの作成には最適です。

今回は、実際にスピードテストの仕組みをゼロから作るのは少し複雑なため、世界的に有名なスピードテストサイト「Speedtest.net」の仕組みをGoで利用できるライブラリを使って、ツールを作成してみましょう。

# 今回作成するプログラム概要
このプログラムは、以下の3つのステップで動きます。
1. サーバーの検索: 一番近い（速い）測定用サーバーを見つける。
2. ダウンロード速度の計測: データをダウンロードして速さを測る。
3. アップロード速度の計測: データをアップロードして速さを測る。

# 穴埋め問題：スピードテスト・ツール
コードの 【 】 の部分を予想して埋めてみてください。 

```go
package main

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
)

func main() {
	// 1. スピードテスト用のクライアントを作成
	client := speedtest.【 A 】()

	// 2. 一番近いサーバーを取得
	serverList, _ := client.FetchServers(user)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		// 3. ダウンロード速度を計測
		err := s.【 B 】()
		if err != nil {
			fmt.Println("エラーが発生しました:", err)
			return
		}

		// 4. アップロード速度を計測
		err = s.【 C 】()
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
```

# 初心者のための学習ポイント
1. ```package main```: このファイルが「実行されるプログラム本体」であることを示します。
2. ```import```: 他の人が作った便利な道具（ライブラリ）を読み込みます。
3. ```:=```: Go特有の書き方で、「変数を作る」と同時に「値を代入する」という意味です。
4. ```func main()```: プログラムがスタートする場所です。

# 実行する方法
Goをインストールした後、パソコンのターミナル（コマンドプロンプト）で以下のコマンドを順番に打ってみてください。

1. プロジェクトの初期化
```bash
go mod init speedTest
```

2. 必要なライブラリのインストール
```bash
go get github.com/showwin/speedtest-go/speedtest
```
3. プログラムの実行(コードを ```main.go``` という名前で保存してから)
```bash
go run main.go
```

# 正解
A: ```New```
B: ```DownloadTest```
C: ```UploadTest```