package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func mapRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", home)
	mux.HandleFunc("/greet", greet)
}

func home(writer http.ResponseWriter, request *http.Request) {
	if request.RequestURI != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodGet:
		_, e := writer.Write([]byte("home"))
		if e != nil {
			log.Print(e)
		}
	default:
		writer.WriteHeader(http.StatusNoContent)
	}
}

func greet(writer http.ResponseWriter, request *http.Request) {
	if request.RequestURI != "/greet" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case http.MethodPost:
		// deserialize request
		var c Client
		e := json.NewDecoder(request.Body).Decode(&c)
		if e != nil {
			sendError(writer, e.Error())
			return
		} else if c.Id == "" || c.Endpoint == "" {
			sendError(writer, "request json is in the incorrect format")
			return
		}

		// add to clients map
		e = addClient(&c)
		if e != nil {
			sendError(writer, e.Error())
			return
		}

		writer.WriteHeader(http.StatusOK)
		reply := fmt.Sprintf("%s client registered", c.Id)
		writer.Write([]byte(reply))

	default:
		writer.WriteHeader(http.StatusNoContent)
	}
}

func sendError(writer http.ResponseWriter, errorStr string) {
	response := &ErrorResponse{Error: errorStr}
	bytes, e := response.Serialize()
	if e != nil {
		panic(e)
	}

	http.Error(writer, string(bytes), http.StatusBadRequest)
}
