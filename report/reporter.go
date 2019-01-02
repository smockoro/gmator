package report

import (
	"fmt"

	"github.com/smockoro/gmator/result"
)

type Reporter interface {
	Build() error
	Report(*result.Result)
}

type stdoutReporter struct{}

func NewStdoutReporter() Reporter {
	return &stdoutReporter{}
}

func (r *stdoutReporter) Build() error {
	// ほんとはオプションをストラクトが持っていてそれを適用していくのがやりたい
	fmt.Println("stdout reporter build")
	return nil
}

func (r *stdoutReporter) Report(rlt *result.Result) {
	fmt.Println("stdout report")

	counter := 0

	for resp := range rlt.RespChan {
		fmt.Println("Go routine resp")
		fmt.Printf("====response loop counter :%d\n", counter)
		fmt.Println(resp)
		counter++
	}
	close(rlt.Done)

}
