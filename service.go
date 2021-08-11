package main

import (
	"context"
	"fmt"
	"github.com/longjoy/gRpc/proto"
	"google.golang.org/grpc"
	"net"
)

type UserInfoService struct {

}
var u = UserInfoService{}
func (s *UserInfoService)GetUserInfo(ctx context.Context, req *proto.UserRequest)(res *proto.UserResponse, err error)  {
	name := req.Name
	if name == "zs" {
		res = &proto.UserResponse{
			Id: 123,
			Age: 29,
			Name: "张三",
			Hobby: []string{"music", "movie"},
		}

	}
	return
}

func main()  {
	addr := "127.0.0.1:8888"
	listener, err := net.Listen("tcp", addr)
	if err !=nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Println("监听地址 %s\n",addr)
	newServer := grpc.NewServer()
	proto.RegisterUserInfoServiceServer(newServer, &u)
	newServer.Serve(listener)


}