package handlers

import (
	"log"
	"net/http"

	"imladris/components"
)

type LoginHandler struct {
	http.Handler
}

func (LoginHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		log.Println("Rendering login page")
		err := components.Login().Render(request.Context(), writer)
		if err != nil {
			log.Println("Error rendering login page", err)
		}
		log.Println("Rendered login page")
		return
	case http.MethodPost:
		username := request.FormValue("username")
		password := request.FormValue("password")
		if username == "user" && password == "pass" {
			//
		} else {
			//
		}
		return
	}
}
