package main

import (
	"flag"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

var wg sync.WaitGroup

func main() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stdout)

	devPtr := flag.Bool("dev", false, "a bool")
	flag.Parse()

	if *devPtr {
		prepareProdServer()
	} else {
		go startVueServer(":80")
	}
	wg.Add(1)

	wg.Wait()
}

func startVueServer(port string) {
	defer wg.Done()

	router := mux.NewRouter().StrictSlash(true)
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

func indexHandler(filename string) func(response http.ResponseWriter, request *http.Request) {
	realHandler := func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, filename)
	}
	return realHandler
}
