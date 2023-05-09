package client

import (
	"github.com/gorilla/websocket"
	"log"
)

type WsClient struct {
	*websocket.Conn
	send    chan []byte
	session *SessionHandler
}

func Dial(addr string, s *SessionHandler) *WsClient {
	// 连接到WebSocket服务器
	conn, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return &WsClient{
		Conn:    conn,
		send:    make(chan []byte, 100),
		session: s,
	}
}
func (w *WsClient) SendMsg(msg []byte) {
	w.send <- msg
}

func (w *WsClient) Startup() {
	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := w.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			w.session.OnRecv(message)
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case data := <-w.send:
				// 向WebSocket服务器发送消息
				err := w.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		}
	}()
}
