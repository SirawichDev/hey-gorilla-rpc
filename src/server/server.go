package main

import (
	"fmt"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"net/http"
)

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(HelloService), "")
	s.RegisterBeforeFunc(func(ri *rpc.RequestInfo){
		fmt.Println(ri)
	})
	http.Handle("/hi_rpc", s)
	http.ListenAndServe(":9999", nil)
}

type Args struct {
	H string
}
type HelloService struct {
}

func (s *HelloService) Hello(req *http.Request, args *Args, reply *string) error {
	*reply = args.H
	fmt.Println("Calculating")
	return nil
}
