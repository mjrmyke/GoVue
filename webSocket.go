package main

import (
	"fmt"
	"math"

	"github.com/olahol/melody"
)

func handleWebSocketEvents(webSocketManager *melody.Melody) {
	webSocketManager.HandleMessage(func(sess *melody.Session, msg []byte) {
		fmt.Println("websock message received: ", msg)
		webSocketManager.Broadcast(msg)
	})

	webSocketManager.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected")
		sg.Done()
	})
}

func retrieveOrCreateRoom(room string) *melody.Melody {
	var webSocketManager *melody.Melody

	if val, ok := webSocketManagerContainer[room]; ok {
		fmt.Println("Retrieved connection manager")
		webSocketManager = val
	} else {
		fmt.Println("Created connection manager")
		webSocketManager = melody.New()
		webSocketManager.Config.MaxMessageSize = math.MaxInt64 - 1
		webSocketManagerContainer[room] = webSocketManager
	}

	return webSocketManager
}
