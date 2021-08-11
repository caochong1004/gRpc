package main

import (
	"context"
	"fmt"
	"github.com/longjoy/gRpc/proto"
	"google.golang.org/grpc"
)

func main()  {
	conn, err := grpc.Dial("127.0.0.1:8888", grpc.WithInsecure())
	if err !=nil {
		fmt.Printf("grpc 连接失败: %s\n",err)
	}
	defer conn.Close()
	client := proto.NewUserInfoServiceClient(conn)
	req := new(proto.UserRequest)
	req.Name = "zs"
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Printf("返回值失败%s\n",err)
	}
	fmt.Printf("返回值%v\n",response.Id)
}
