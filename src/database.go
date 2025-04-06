package main

import "net/http"

type database interface {
	LogRequest(request *http.Request) error
}

func logMessage() {

}
