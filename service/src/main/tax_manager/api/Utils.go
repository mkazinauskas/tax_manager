package api

import (
	"encoding/json"
)

func Marshal(param interface{}) (string) {
	marshaled, _ := json.Marshal(param)
	return string(marshaled)
}
