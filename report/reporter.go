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
loop:
	for {
		select {
		case _, ok := <-rlt.ReqChan:
			fmt.Println(ok)
			if !ok {
				close(rlt.RespChan)
				break loop
			}
		default:
			fmt.Println("Go routine resp")
			fmt.Println(<-rlt.RespChan)
			counter++
			fmt.Printf("====response loop counter :%d\n", counter)
		}

	}
	close(rlt.Done)

}
