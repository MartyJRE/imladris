package handlers

import (
	"log"
	"net/http"

	"imladris/components"
)

type SettingsHandler struct {
	http.Handler
}

func (SettingsHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		log.Println("Rendering settings page")
		err := components.Settings().Render(request.Context(), writer)
		if err != nil {
			log.Println("Error rendering settings page", err)
		}
		log.Println("Rendered settings page")
	default:
		log.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
	}
}
