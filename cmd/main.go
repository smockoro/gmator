package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	hdl "github.com/smockoro/gmator/handle"
	rpt "github.com/smockoro/gmator/report"
	rlt "github.com/smockoro/gmator/result"
)

func main() {
	// 引数を処理する
	//   - 実行方式
	//    - 回数
	//    - 並行数
	//    - 間隔
	//    - アクセス先
	//    - (アクセス先(リスト))
	//    - (認証)
	//   - レポート方式
	//    - 標準出力
	//    - ログファイル名
	var (
		times       int
		concurrents int
		interval    time.Duration
		url         string
		format      string
		filename    string
	)

	flag.IntVar(&times, "times", 1, "access times / interval")
	flag.IntVar(&concurrents, "concurrents", 1, "concurrents")
	flag.DurationVar(&interval, "interval", 1*time.Second, "access interval")
	flag.StringVar(&url, "url", "http://localhost", "access url")
	flag.StringVar(&format, "format", "stdout", "output format")
	flag.StringVar(&filename, "filename", "report.log", "log filename")

	flag.Parse()

	// 引数に合わせてhanderを構成
	// - リクエスト先
	// - 結果の出力形式
	// - チャネルでデータ共有
	// TODO:引数に合わせてビルドする処理を作成する
	result := rlt.NewResult()

	handler, err := hdl.NewHandler(
		hdl.SetTimes(times),
		hdl.SetConcurrents(concurrents),
		hdl.SetInterval(interval),
		hdl.SetURL(url),
	)
	if err != nil {
		log.Fatal(err)
	}

	reporter, err := rpt.NewStdoutReporter(
		rpt.SetFormat(format),
	)
	if err != nil {
		log.Fatal(err)
	}

	// 実行開始
	// goroutineとして実行結果をキューに入れる
	go handler.Do(result)

	// 結果出力
	// goroutineとしてキューの結果を出力する
	go reporter.Report(result)

	// 2つのgoroutineの処理が完了するまで待つ
loop:
	for {
		select {
		case <-result.Done:
			break loop
		default:
		}
	}

	// 後片付け（あれば）
	fmt.Println("after care")

}
