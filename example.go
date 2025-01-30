package main

import (
	"encoding/json"
	"log"
)

func myLog(format string, args ...interface{}) {
	const prefix = "[my] "
	log.Printf(prefix+format, args...)
}

func doSomething(msg string) (json.RawMessage, error) {

	m := map[string]string{
		"msg": msg,
	}

	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return json.RawMessage(b), nil

}
