package p2p

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/jonggulee/go-coin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)
	openPort := r.URL.Query().Get("openPort")
	result := strings.Split(r.RemoteAddr, ":")
	initPeer(conn, result[0], openPort)
}

func AddPeer(address, port, openPort string) {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)
	initPeer(conn, address, port)
}
