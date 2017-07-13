package utils

import (
	"errors"
	"fmt"
	"log"
)

func Check(err error) {
	if err != nil {
		log.Panicln(err)
		panic(err)
	}
}

func CheckError(err error, messageTemplate string, arguments ... interface{}) {
	if err != nil {
		log.Panicln(err)
		panic(errors.New(fmt.Sprintf(messageTemplate, arguments)))
	}
}

func Error(messageTemplate string, arguments ... interface{}) {
	panic(errors.New(fmt.Sprintf(messageTemplate, arguments)))
}
