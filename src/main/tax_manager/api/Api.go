package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)

func Initialize() {
	router := httprouter.New()
	router.GET("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}
