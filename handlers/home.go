package handlers

import (
	"log"
	"net/http"

	"imladris/components"
	"imladris/models"
)

type HomeHandler struct {
	http.Handler
}

func (HomeHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		log.Println("Rendering home page")
		err := components.Home([]models.Person{}).Render(request.Context(), writer)
		if err != nil {
			log.Println("Error rendering home page", err)
		}
		log.Println("Rendered home page")
	default:
		log.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
