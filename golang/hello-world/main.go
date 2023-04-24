package main

import "hello-world/handler"

func init() {
	a := HostImpl{}
	handler.SetHandler(a)
}

type HostImpl struct {
}

func (e HostImpl) Handle(req handler.HandlerRequest) handler.HandlerResponse {
	return handler.HandlerResponse{Status: 200, Body: "Hello from GO!\n"}
}

//go:generate wit-bindgen tiny-go ./wit --out-dir=handler
func main() {}
