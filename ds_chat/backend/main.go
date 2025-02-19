package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool) // 存储所有客户端连接
var broadcast = make(chan Message)           // 消息广播通道

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func main() {
	r := gin.Default()

	r.GET("/ws", handleWebSocket) // WebSocket 路由
	go handleMessages()           // 处理消息广播

	r.Run(":8080")
}

// 处理 WebSocket 连接
func handleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	clients[conn] = true

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			delete(clients, conn)
			break
		}
		broadcast <- msg
	}
}

// 广播消息给所有客户端
func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
