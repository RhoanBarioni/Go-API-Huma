package main

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// o Body é a resposta da API, o huma vai transformar as informações em um JSON
type Greetingoutput struct {
	Body struct { // body é o padrão q os devs usam
		Message string `json:"message" example:"hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a new router & API
	router := chi.NewMux() // pega as rotas
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0")) // humachi.new vai unir junto com o chi para que funcione junto com o swagger

	// TODO: Register operations...

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}
