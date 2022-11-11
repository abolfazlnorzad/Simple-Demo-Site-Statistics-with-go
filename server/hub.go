package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type hub struct {
	Conn *websocket.Conn
}

type Message struct {
	Count    int    `json:"Count"`
	SiteName string `json:"SiteName"`
}

func newHub(cn *websocket.Conn) *hub {
	return &hub{
		Conn: cn,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *hub) run() {
	rand.Seed(time.Now().UnixNano())
	for {
		_, message, err := h.receive()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		SiteName := string(message)
		for {
			time.Sleep(time.Second)
			msg := Message{
				Count:    rand.Intn(501),
				SiteName: SiteName,
			}
			h.send(msg)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
}

func (h *hub) receive() (messageType int, message []byte, err error) {
	messageType, message, err = h.Conn.ReadMessage()
	if err != nil {
		fmt.Println(err)
		return 0, nil, err
	}
	return
}

func (h *hub) send(msg Message) error {
	err := h.Conn.WriteJSON(msg)
	if err != nil {
		return err
	}
	return nil
}
