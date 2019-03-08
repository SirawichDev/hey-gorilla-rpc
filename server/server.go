package main

import (
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")

}

type Args struct {
	H string
}
type HelloService struct {
}

func (s *HelloService) Hello(req *http.Request, args *Args, reply *string) error {
	*reply = args.H
	return nil
}