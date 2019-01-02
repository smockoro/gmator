package result

import "net/http"

// Result :
type Result struct {
	ReqChan  chan *http.Request
	RespChan chan *http.Response
	Done     chan interface{}
}

// NewResult : Resultsインターフェースでresult型を返す
func NewResult() *Result {
	reqChan := make(chan *http.Request)
	respChan := make(chan *http.Response)
	done := make(chan interface{})
	return &Result{ReqChan: reqChan, RespChan: respChan, Done: done}
}
