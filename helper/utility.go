package helper

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func ConnectToWebSocket(url string) (*websocket.Conn, error) {
	var conn *websocket.Conn
	var err error
	st := time.Now()

	for {
		conn, _, err = websocket.DefaultDialer.Dial(url, nil)
		if err == nil {
			return conn, nil
		}

		log.Println("[socket-aggregator] connection error:", err)
		time.Sleep(250 * time.Millisecond) // Retry after 2 seconds

		if latency := time.Since(st).Milliseconds(); latency >= 500 {
			err = errors.New("socket connection failure")
			return nil, err
		}
	}
}

func HandleResponse(rw http.ResponseWriter, statusCode int, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{"data": message})
}
