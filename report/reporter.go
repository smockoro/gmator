package report

import "fmt"

type Reporter interface {
	Build() error
	Report() error
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

func (r *stdoutReporter) Report() error {
	fmt.Println("stdout report")
	return nil
}
