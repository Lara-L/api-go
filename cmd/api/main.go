package main

import (
	"fmt"
	"github.com/Lara-L/api-go/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Iniciando GO API serviço")
	fmt.Println(`
____ ____ _    ____ _  _ ____    ____ ___  _ 
| __ |  | |    |__| |\ | | __    |__| |__] | 
|__] |__| |___ |  | | \| |__]    |  | |    |`)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
