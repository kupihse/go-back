package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

	"github.com/go-chi/chi"
)

var port = ":8080"

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var productStorage = make(map[string]*Product)

func main() {
	router := chi.NewRouter()
	router.Route("/pr", func(r chi.Router) {
		r.Post("/new", func(w http.ResponseWriter, r *http.Request) {
			var p Product
			json.NewDecoder(r.Body).Decode(&p)
			p.Id = uuid.NewV4().String()
			productStorage[p.Id] = &p
		})
		r.Get("/all", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(productStorage)
		})
	})
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, router)
}
