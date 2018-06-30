package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/olahol/melody"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup
var sg sync.WaitGroup

var webSocketManagerContainer map[string]*melody.Melody

// var webSocketManager *melody.Melody

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)
	webSocketManagerContainer = make(map[string]*melody.Melody)

	devPtr := flag.Bool("dev", false, "a bool")
	flag.Parse()

	if !*devPtr {
		fmt.Println("Starting Server, serving on :80")
		go startVueServer(":80")
	} else {
		fmt.Println("Starting Dev Server, serving on localhost:80")
		go startVueServer("localhost:80")
	}

	wg.Add(2)

	wg.Wait()

}

func startVueServer(port string) {
	defer wg.Done()

	router := mux.NewRouter().StrictSlash(true)

	apiRouter := router.PathPrefix("/api/").Subrouter()

	router.HandleFunc("/ws/roll/{room}", startWebSocketConnection)

	apiRouter.HandleFunc("/user/create", createNewUser)

	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("VueApp/dist")))
	router.PathPrefix("/").HandlerFunc(indexHandler("VueApp/dist/index.html"))

	logger := log.New()
	logwriter := logger.Writer()
	defer logwriter.Close()

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(logwriter, router),
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func createNewUser(response http.ResponseWriter, request *http.Request) {

}

func startWebSocketConnection(response http.ResponseWriter, request *http.Request) {
	urlPathVars := mux.Vars(request)
	var webSocketManager *melody.Melody

	if val, ok := webSocketManagerContainer[urlPathVars["room"]]; ok {
		fmt.Println("Retrieved connection manager")
		webSocketManager = val
	} else {
		fmt.Println("Created connection manager")
		webSocketManager = melody.New()
		webSocketManager.Config.MaxMessageSize = math.MaxInt64 - 1
		webSocketManagerContainer[urlPathVars["room"]] = webSocketManager
	}

	fmt.Println("websocket attempted to connect")
	sg.Add(1)

	webSocketManager.HandleMessage(func(sess *melody.Session, msg []byte) {
		fmt.Println("websock message received: ", msg)
		webSocketManager.Broadcast(msg)
	})

	webSocketManager.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("disconnected")
		sg.Done()
	})
	webSocketManager.HandleRequest(response, request)

	sg.Wait()
}

func indexHandler(filename string) func(response http.ResponseWriter, request *http.Request) {
	realHandler := func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, filename)
	}
	return realHandler
}
