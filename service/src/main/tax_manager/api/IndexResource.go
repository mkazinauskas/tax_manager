package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func GetIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
