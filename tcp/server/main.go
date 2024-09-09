func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed,err: ", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据: ", recvStr)
		conn.Write([]byte(recvStr))
	}
}