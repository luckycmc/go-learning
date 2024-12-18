package rpc

import (
	"fmt"
	"net"
	"reflect"
)

type Server struct {
	// 地址
	addr  string
	funcs map[string]reflect.Value
}

func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

// 服务端需要一个注册register
// 第一个参数函数明，第二个传入真正的函数
func (s *Server) Register(rpcName string, f interface{}) {
	// 维护一个map
	// 如果map已经有键了
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	// 如果没有，则将映射加入map，用于调用
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

// 服务端等待调用的方法
func (s *Server) Run() {
	// 监听
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	for {
		// 服务端循环等待调用
		conn, err := lis.Accept()
		if err != nil {
			return
		}
		serSession := NewSession(conn)
		// 使用rpc方式读取数据
		b, err := serSession.Read()
		if err != nil {
			return
		}
		// 数据解码
		rpcData, err := decode(b)
		if err != nil {
			return
		}

		// 根据读到的name，得到要调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			fmt.Printf("rpc function %s not found", rpcData.Name)
			return
		}
		// 遍历解析客户端传来的参数，放切片里
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法
		out := f.Call(inArgs)
		// 遍历out，用户返回给客户端，存到切片里
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		// 数据编码，返回给客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respRPCData)
		if err != nil {
			return
		}
		// 将服务端编码后的数据，写出到客户端
		err = serSession.Write(bytes)
		if err != nil {
			return
		}

	}

}
