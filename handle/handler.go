package handle

import "fmt"

type Handler interface {
	Build() error
	Do() error
}

func NewHandler() Handler {
	return &handler{}
}

type handler struct{}

func (h *handler) Build() error {
	fmt.Println("Http Handler Build")
	return nil
}

func (h *handler) Do() error {
	fmt.Println("Http Handler do rush!")
	return nil
}
