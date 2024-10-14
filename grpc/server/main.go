package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "mysql/grpc/user"
	"net"
)

// UserInfoService 定义空接口
type UserInfoService struct {
	pb.UnimplementedUserInfoServiceServer
}

var u = UserInfoService{}

// 实现方法
func (s *UserInfoService) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
	// 通过用户名查询用户信息
	name := req.Name
	if name == "kevin" {
		resp = &pb.UserResponse{
			Id:    1,
			Name:  name,
			Age:   22,
			Hobby: []string{"Sing", "run"},
		}
	}
	return
}

func main() {
	// 地址
	addr := "127.0.0.1:8080"
	// 监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常：%s\n", err)
	}
	fmt.Printf("监听端口:%s\n", addr)
	// 实例化grpc
	s := grpc.NewServer()
	// 在grpc上注册微服务
	pb.RegisterUserInfoServiceServer(s, &u)
	// 启动服务端
	s.Serve(listener)
}
