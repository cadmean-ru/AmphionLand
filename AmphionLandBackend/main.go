package main

import (
	"errors"
	"github.com/cadmean-ru/goRPCKit/rpc"
	"net/http"
)

func main() {
	rpc.Handle("login", func(args ...interface{}) (interface{}, error) {
		var email, password string

		if emailStr, ok := args[0].(string); ok {
			email = emailStr
		}

		if passwordStr, ok := args[1].(string); ok {
			email = passwordStr
		}

		if email == "bruh" && password == "123" {
			return "ok", nil
		}

		return nil, errors.New("breh")
	})

	_ = http.ListenAndServe(":4200", nil)
}
