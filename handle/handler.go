package handle

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/smockoro/gmator/result"
)

type Handler interface {
	Build() error
	Do(*result.Result)
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Build() error {
	fmt.Println("Http Handler Build")
	return nil
}

func (h *handler) Do(rlt *result.Result) {
	fmt.Println("Http Handler do rush!")

	// とりあえず3つのアクセス先に向けてリクエストを送ることができるようになったので、1つに搾って機能を追加しよう
	go func() {
		req1, _ := http.NewRequest("GET", "http://www.google.com", nil)
		req2, _ := http.NewRequest("GET", "http://www.yahoo.com", nil)
		req3, _ := http.NewRequest("GET", "http://www.microsoft.com", nil)
		for i := 0; i < 10; i++ {
			fmt.Printf("====reqest loop counter :%d\n", i)
			rlt.ReqChan <- req1
			rlt.ReqChan <- req2
			rlt.ReqChan <- req3
			time.Sleep(1 * time.Second) // アクセスするリクエストを作るタイミングをここで調整している
		}
		close(rlt.ReqChan)
	}()

loop:
	for {
		select {
		case req, ok := <-rlt.ReqChan:
			if ok {
				go func(req *http.Request) {
					fmt.Println("Go routine")
					c := &http.Client{}
					resp, err := c.Do(req)
					if err != nil {
						log.Print(err)
					}
					rlt.RespChan <- resp
				}(req)
			}
			if !ok {
				close(rlt.RespChan)
				break loop
			}
		}
	}
}
