package main

import (
	"net/http"

	serviceapi "github.com/RhoanBarioni/Go-API-Huma/service"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// o Body é a resposta da API, o huma vai transformar as informações em um JSON
type GreetingOutput struct {
	Body struct { // body é o padrão q os devs usam
		Message string `json:"message" example:"hello, world!" doc:"Greeting message"`
	}
}

func main() {
	// Create a new router & API
	router := chi.NewMux()                                            // pega as rotas
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0")) // humachi.new vai unir junto com o chi para que funcione junto com o swagger

	// Register GET /greeting/{name} handler.
	huma.Get(api, "/greeting/{name}", serviceapi.GreetingEndpoint)
	huma.Post(api, "/aluno", serviceapi.AlunoEndpoint)
	huma.Put(api, "/aluno/{id}", serviceapi.AlunoIdEndpoint)

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}

// curl -X GET http://127.0.0.1:8888/greeting/John -H "Content-Type: application/json" -d '{"nome": "John Doe", "notas": [8, 9, 10]}'
