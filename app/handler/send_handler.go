package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/mindwingx/graph-coordinator/helper"
	"log"
	"net/http"
	"sync"
)

type (
	wsItem struct {
		State   string `json:"state"`
		Message string `json:"message"`
	}
)

var mx = sync.RWMutex{}

func SendHandler(rw http.ResponseWriter, r *http.Request) {
	var msg, resp wsItem

	res, ok := r.Context().Value("decodedMsg").(string)
	if !ok {
		helper.HandleResponse(rw, http.StatusUnprocessableEntity,
			"[api] Failed to retrieve decoded data from context",
		)
		return
	}

	ws := wsConn()
	if ws == nil {
		helper.HandleResponse(rw, http.StatusUnprocessableEntity, "[api][ws] socket failed")
		return
	} else {
		defer func() { _ = ws.Close() }()
	}

	msg = wsItem{
		State:   "api",
		Message: res,
	}

	mx.Lock()

	err := ws.WriteJSON(msg)
	if err != nil {
		fmt.Println("[api][ws] socket write-json error:", err)
		return
	}
	mx.Unlock()

	err = ws.ReadJSON(&resp)
	if err != nil {
		helper.HandleResponse(rw, http.StatusUnprocessableEntity, "[api][ws] socket failed")
		return
	}

	helper.HandleResponse(rw, http.StatusCreated, "[api][ws] "+resp.Message)
}

// HELPER METHODS

func wsConn() *websocket.Conn {
	conn, err := helper.ConnectToWebSocket(helper.AggregatorSocketUrl)
	if err != nil {
		log.Println("[api-coordinator]", err)
		return nil
	}

	return conn
}
