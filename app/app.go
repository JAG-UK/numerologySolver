package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/JAG-UK/numerologySolver/app/handler"
	"github.com/JAG-UK/numerologySolver/config"
)

func InitAndRun(conf config.Config) {

	// Initialise resources and services

	router := mux.NewRouter()

	handler.InitBasicCharMap()
	handler.PrecompWordList()

	// Register management functions
	//router.HandleFunc("/wordlist", handler.initDB).Methods("PUT")
	//router.HandleFunc("/wordlist", handler.eraseDB).Methods("DELETE")

	// Register operational functions
	router.HandleFunc("/word", handler.GetValueOfWord).Methods("PUT")
	router.HandleFunc("/number", handler.GetAllWordsOfValue).Methods("PUT")

	log.Printf("Listening on port %v", conf.ListenOn)
	log.Fatal(http.ListenAndServe(conf.ListenOn, router))
}
