package main

import (
	"fmt"
	"math"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/olahol/melody"
)

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
		fmt.Println("websock message received: ", msg)
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
		webSocketManager.manager.Config.MaxMessageSize = math.MaxInt64 - 1
		webSocketManager.route = room
		webSocketManagerContainer[room] = webSocketManager
	}

	return webSocketManager
}
