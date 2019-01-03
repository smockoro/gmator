package handle

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/smockoro/gmator/result"
)

const (
	DefaultTimes    int           = 1
	DefaultConc     int           = 1
	DefaultInterval time.Duration = 1 * time.Second
	DefaultURL      string        = "http://localhost"
)

type Handler interface {
	Do(*result.Result)
}

type handler struct {
	Times       int
	Concurrents int
	Interval    time.Duration
	URL         string
}

type HanlderOptionFunc func(*handler) error

func NewHandler(options ...HanlderOptionFunc) (Handler, error) {
	h := &handler{
		Times:       DefaultTimes,
		Concurrents: DefaultConc,
		Interval:    DefaultInterval,
		URL:         DefaultURL,
	}

	for _, option := range options {
		if err := option(h); err != nil {
			return nil, err
		}
	}

	return h, nil
}

func SetTimes(times int) HanlderOptionFunc {
	return func(h *handler) error {
		h.Times = times
		return nil
	}
}

func SetConcurrents(concurrents int) HanlderOptionFunc {
	return func(h *handler) error {
		h.Concurrents = concurrents
		return nil
	}
}

func SetInterval(interval time.Duration) HanlderOptionFunc {
	return func(h *handler) error {
		h.Interval = interval
		return nil
	}
}

func SetURL(url string) HanlderOptionFunc {
	return func(h *handler) error {
		h.URL = url
		return nil
	}
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
