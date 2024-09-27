package handlers

import (
	"io"
	"log"
	"net/http"
	"serverStateService/models"
	"text/template"
)

func GetServersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		model := models.ServerModel{}

		servers, err := model.GetServers()
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "internal server error")
			log.Fatalln("Error getting servers")
			return
		}

		tmpl, err := template.ParseFiles("views/index.html")
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "internal server error")
			log.Fatalln("Error parsing index.html")
			return
		}

		err = tmpl.Execute(w, servers)
		if err != nil {
			w.WriteHeader(500)
			io.WriteString(w, "internal server error")
			log.Fatalln("Error executing template")
			return
		}
	}
}
