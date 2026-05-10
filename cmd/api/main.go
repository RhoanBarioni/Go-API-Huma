package main

import (
	"net/http"

	"github.com/RhoanBarioni/Go-API-Huma/internal/handlers"
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
// testandfo se a i.a vai me retornar o git commit certo
func main() {
	// Create a new router & API
	router := chi.NewMux()                                            // pega as rotas
	api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0")) // humachi.new vai unir junto com o chi para que funcione junto com o swagger

	// Register GET /greeting/{name} handler.
	// huma.Get(api, "/greeting/{name}", handlers.GreetingEndpoint)
	// GET /alunos
	huma.Get(api, "/alunos", handlers.GetAlunos)
	// GET /aluno/{id}
	huma.Get(api, "/aluno/{id}", handlers.GetAlunosId)
	huma.Post(api, "/aluno", handlers.CreateAluno)
	huma.Put(api, "/aluno/{id}", handlers.UpdateAluno)
	// DELETE /aluno/{id}
	huma.Delete(api, "/aluno/{id}", handlers.DeleteAlunoId)

	// Start the server!
	http.ListenAndServe("127.0.0.1:8888", router)
}

// curl -X GET http://127.0.0.1:8888/greeting/John -H "Content-Type: application/json" -d '{"nome": "John Doe", "notas": [8, 9, 10]}'
