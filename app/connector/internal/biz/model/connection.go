package model

import (
	"time"

	"github.com/gorilla/websocket"
)

type Connection struct {
	conn *websocket.Conn
	in   chan any
}

func NewConnection(conn *websocket.Conn) *Connection {
	return &Connection{
		conn: conn,
		in:   make(chan any, 64),
	}
}

func (c *Connection) Send(msg any) {
	select {
	case c.in <- msg:
	default:
		// 队列已满则丢弃
		return
	}
}

func (c *Connection) Listen() {
	for msg := range c.in {
		_ = c.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		if err := c.conn.WriteJSON(msg); err != nil {
			return
		}
	}
}

func (c *Connection) Close() {
	_ = c.conn.Close()
}
