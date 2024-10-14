package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "mysql/grpc/user"
)

func main() {
	// 连接
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("连接异常：%s\n", err)
	}
	defer conn.Close()
	// 实例化grpc客户端
	client := pb.NewUserInfoServiceClient(conn)
	// 组装请求参数
	req := new(pb.UserRequest)
	req.Name = "kevin"
	// 调用接口
	response, err := client.GetUserInfo(context.Background(), req)
	if err != nil {
		fmt.Println("响应异常 %s\n", err)
	}
	fmt.Printf("响应结果%v \n", response)
}
