package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// ConnClients socket客户
type ConnClients struct {
	conn *websocket.Conn // 连接对象
	send chan []byte     // 发送消息通道
}

// ConnServer socket容器
type ConnServer struct {
	broadcast  chan []byte                   // 广播消息通道
	register   chan *ConnClients             // 注册的客户
	unregister chan *ConnClients             // 未注册的客户
	clients    map[*ConnClients]*ConnClients // 客户集合
}

// 定义全局控制器
var connServer = &ConnServer{
	broadcast:  make(chan []byte),
	register:   make(chan *ConnClients),
	unregister: make(chan *ConnClients),
	clients:    make(map[*ConnClients]*ConnClients),
}

// 升级配置
var socketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func websocketServer(w http.ResponseWriter, r *http.Request) {
	// 将http协议升级为websocket协议
	conn, err := socketUpgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// 新客户
	connClient := &ConnClients{
		conn: conn,
		send: make(chan []byte, 256),
	}
	// 加入注册列表
	connServer.register <- connClient

	// 开启单独线程用于收发消息
	go connClient.receiveMsg()
	go connClient.sendMsg()
}

type WebsocketData struct {
	Action string
	Id     string
}

func (c *ConnClients) receiveMsg() {
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		var msg map[string]interface{}
		err = json.Unmarshal(message, &msg)
		// 解析socket消息，如果是close，则注销当前客户，如果是心跳信息ping，则返回pong
		if action, ok := msg["Action"]; ok {
			if action == "close" {
				connServer.unregister <- c
				break
			}
			if action == "ping" {
				j, _ := json.Marshal(&map[string]string{"Action": "pong"})
				c.send <- j
			}
		} else {
			// c.send <- message
			AddMsg(message)
		}
	}
}
func (c *ConnClients) sendMsg() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		// 循环遍历发送消息通道
		case message, ok := <-c.send:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (m *ConnServer) run() {
	for {
		select {
		case client := <-m.register: // 新增客户
			m.clients[client] = client
		case client := <-m.unregister: // 注销
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.send)
			}
		case message := <-m.broadcast: // 广播消息到每个客户
			for client := range m.clients {
				client.send <- message
			}
		}
	}
}

// AddMsg 公用方法 推送消息给客户端
func AddMsg(t []byte) {
	connServer.broadcast <- t
}
