package plants

import (
	"log"
	"net/http"
)

type home struct {
}

func mapRoutes(mux *http.ServeMux) {
	mux.Handle("/", &home{})
}

func (h *home) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.RequestURI != "/" {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	switch request.Method {
	case "GET":
		_, e := writer.Write([]byte("home"))
		if e != nil {
			log.Print(e)
		}
	default:
		writer.WriteHeader(http.StatusNoContent)
	}
}
