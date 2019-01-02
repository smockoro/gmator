package main

import (
	"fmt"

	"github.com/smockoro/gmator/result"
)

func main() {
	// 引数を処理する

	// 引数に合わせてhanderを構成
	// - リクエスト先
	// - 結果の出力形式
	// - チャネルでデータ共有
	result := result.NewResult()
	//defer close(result.ReqChan)
	//defer close(result.RespChan)

	//handler := handle.NewHandler()
	//reporter := report.NewStdoutReporter()

	// 実行開始
	// goroutineとして実行結果をキューに入れる
	//go handler.Do(result)

	// 結果出力
	// goroutineとしてキューの結果を出力する
	//go reporter.Report(result)

	// 実験
	ch := make(chan interface{})

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("write %d\n", i)
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer close(result.Done)
		for v := range ch {
			fmt.Printf("read %d\n", v)
		}
	}()

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
