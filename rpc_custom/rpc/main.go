package rpc

import (
	"encoding/gob"
	"fmt"
	"net"
	"testing"
)

type User struct {
	Name string
	Age  int
}

// 测试用户查询方法
func queryUser(uid int) (User, error) {
	user := make(map[int]User)
	// 假数据
	user[0] = User{"kevin", 18}
	user[1] = User{"kevin1", 18}
	user[2] = User{"kevin2", 18}
	// 模拟查询用户
	if u, ok := user[uid]; ok {
		return u, nil
	}
	return User{}, fmt.Errorf("%d err", uid)
}

func TestRPC(t *testing.T) {
	// 编码中有一个字段是interface{}时要注册一下
	gob.Register(User{})
	addr := "127.0.0.1:8080"
	// 创建服务端
	srv := NewServer(addr)
	// 注册服务端方法
	srv.Register("queryUser", queryUser)
	// 服务端等待调用
	go srv.Run()
	// 客户端获取连接
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println(err)
	}
	// 创建客户端对象
	cli := NewClient(conn)
	// 声明函数原型
	var query func(int) (User, error)
	cli.callRpc("queryUser", &query)
	// 得到查询结果
	u, err := query(1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
}
