package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

type rollRoomMessage struct {
	From    string `json:"from"`
	Message string `json:"message"`
	System  string `json:"system"`
	Data    string `json:"data"`
	Room    string `json:"room"`
}

type roomInfo struct {
	manager *melody.Melody
	route   string
	rollers map[string]socketInfo
}

type socketInfo struct {
	host        string
	lastmessage time.Time
	userName    string
}

func startWebSocketConnection(response http.ResponseWriter, request *http.Request) {
	var sg sync.WaitGroup
	sg.Add(1)
	var connection socketInfo

	urlPathVars := mux.Vars(request)
	websocketManager := retrieveOrCreateRoom(urlPathVars["room"])

	handleWebSocketEvents(websocketManager, connection, &sg)
	websocketManager.manager.HandleRequest(response, request)
	sg.Wait()
}

func handleWebSocketEvents(webSocketManager roomInfo, connection socketInfo, sg *sync.WaitGroup) {
	webSocketManager.manager.HandleMessage(func(sess *melody.Session, msg []byte) {
		var message rollRoomMessage
		err := json.Unmarshal(msg, &message)

		if err != nil {
			fmt.Println("Failed to decode json: ", err)
		}

		fmt.Println("websock message received: ", message)

		// if len(message.System) > 0 {
		// 	if !(webSocketManagerContainer[message.Room].rollers[message.From]) {
		// 		webSocketManagerContainer[message.Room].rollers[message.From] = true
		// 		webSocketManagerContainer[message.Room] = webSocketManager
		// 	}
		// }

		webSocketManager.manager.Broadcast(msg)
	})

	webSocketManager.manager.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected")
		sg.Done()
	})
}

func retrieveOrCreateRoom(room string) roomInfo {
	var webSocketManager roomInfo

	if val, ok := webSocketManagerContainer[room]; ok {
		fmt.Println("Retrieved connection manager")
		webSocketManager = val
	} else {
		fmt.Println("Created connection manager")
		webSocketManager.manager = melody.New()
		webSocketManager.rollers = make(map[string]socketInfo)
		webSocketManager.manager.Config.MaxMessageSize = math.MaxInt64 - 1
		webSocketManager.route = room
		webSocketManagerContainer[room] = webSocketManager
	}

	return webSocketManager
}
