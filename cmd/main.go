package main

import (
	"fmt"

	"github.com/smockoro/gmator/handle"
	"github.com/smockoro/gmator/report"
)

func main() {
	// 引数を処理する

	// 引数に合わせてhanderを構成
	// - リクエスト先
	// - 結果の出力形式
	handler := handle.NewHandler()
	reporter := report.NewReporter()

	// 実行
	handler.Do()

	// 結果を出力
	reporter.Report()

	// 後片付け（あれば）
	fmt.Println("after care")

}
