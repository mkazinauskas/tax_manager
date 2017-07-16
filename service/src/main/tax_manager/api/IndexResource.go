package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func GetIndex(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Tax Calculator!\nApi documentation url: https://github.com/modestukasai/tax_manager")
}
