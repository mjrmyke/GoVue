package main

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"

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

func startWebSocketConnection(response http.ResponseWriter, request *http.Request) {
	sg.Add(1)
	urlPathVars := mux.Vars(request)
	websocketManager := retrieveOrCreateRoom(urlPathVars["room"])

	handleWebSocketEvents(websocketManager)
	websocketManager.manager.HandleRequest(response, request)
	sg.Wait()
}

func handleWebSocketEvents(webSocketManager roomInfo) {
	webSocketManager.manager.HandleMessage(func(sess *melody.Session, msg []byte) {
		var message rollRoomMessage
		err := json.Unmarshal(msg, &message)

		if err != nil {
			fmt.Println("Failed to decode json: ", err)
		}

		fmt.Println("websock message received: ", message)

		if len(message.System) > 0 {
			if !(webSocketManagerContainer[message.Room].rollers[message.From]) {
				webSocketManagerContainer[message.Room].rollers[message.From] = true
				webSocketManagerContainer[message.Room] = webSocketManager
			}
		}

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
		webSocketManager.rollers = make(map[string]bool)
		webSocketManager.manager.Config.MaxMessageSize = math.MaxInt64 - 1
		webSocketManager.route = room
		webSocketManagerContainer[room] = webSocketManager
	}

	return webSocketManager
}
