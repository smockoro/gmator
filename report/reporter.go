package report

import (
	"fmt"

	"github.com/smockoro/gmator/result"
)

const (
	DefaultFormat   string = "stdout"
	DefaultFilename string = "report.log"
)

type Reporter interface {
	Report(*result.Result)
}

type stdoutReporter struct {
	Format string
}

type StdoutReporterOptionFunc func(*stdoutReporter) error

func NewStdoutReporter(options ...StdoutReporterOptionFunc) (Reporter, error) {
	r := &stdoutReporter{}

	for _, option := range options {
		if err := option(r); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func SetFormat(format string) StdoutReporterOptionFunc {
	return func(r *stdoutReporter) error {
		r.Format = format
		return nil
	}
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
