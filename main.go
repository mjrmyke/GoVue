package main

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

func main() {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("Could not determine working directory")
	// }

	router := mux.NewRouter().StrictSlash(true)
	router.PathPrefix("/static").Handler(http.FileServer(http.Dir("VueApp/dist")))
	router.PathPrefix("/").HandlerFunc(indexHandler("VueApp/dist/index.html"))

	srv := &http.Server{
		Handler:      handlers.LoggingHandler(os.Stdout, router),
		Addr:         ":80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

func indexHandler(filename string) func(response http.ResponseWriter, request *http.Request) {
	routerLogger()
	realHandler := func(response http.ResponseWriter, request *http.Request) {
		http.ServeFile(response, request, filename)
	}
	return realHandler

}

func routerLogger() func(response http.ResponseWriter, request *http.Request) *http.Response {
	return func(response http.ResponseWriter, request *http.Request) *http.Response {
		log.Infof("%s \"%s %s %s\"", request.RemoteAddr, request.Method, request.RequestURI, request.Proto)
		return nil
	}
}
