package contracts

import "github.com/RhoanBarioni/Go-API-Huma/internal/models"

// O ideal é organizar esses caras em "contracts".
type GetAlunosResquest struct {
	Nome string `query:"nome" example:"Bruno" doc:"Pega o nome do aluno"`
}

type GetAlunosResponse struct {
	Body struct {
		Alunos  []models.Aluno `json:"nome" example:"[]" doc:"Pega o nome do aluno dentro do JSON"`
		Message string         `json:"message" example:"Olá, Bruno. Aqui está as suas notas" doc:"Pega o aluno e suas informações dentro do json e monta uma mensagem para ele"`
	}
}

type GetAlunosIdResquest struct {
	Id string `query:"id" example:"5" doc:"Pega o id do aluno"`
}

type GetAlunosIdResponse struct {
	Body struct {
		Nome    string `json:"nome" example:"Bruno" doc:"Pega o nome do aluno dentro do JSON"`
		Message string `json:"message" example:"Olá, Bruno. Aqui está as suas notas" doc:"Pega o aluno e suas informações dentro do json e monta uma mensagem para ele"`
	}
}

type AlunoRequest struct {
	Body struct {
		Nome      string    `json:"nome" example:"Bruno" doc:"Pegar o nome do Aluno dentro do JSON /aluno"`
		Sobrenome string    `json:"sobrenome" example:"Eita" doc:"Pegar o sobrenome do Aluno dentro do JSON /aluno"`
		Notas     []float32 `json:"notas" example:"[9.5, 4.5]" doc:"Pegar as notas do Aluno dentro do JSON /aluno"`
	}
}

type AlunoResponse struct {
	Body struct {
		Message string  `json:"message" example:"Aqui está suas notas Bruno: " doc:"Pegar uma mensagem para retornar informações para o Aluno /aluno"`
		Media   float32 `json:"media" example:"10" doc:"Pegar a media do Aluno dentro do JSON"`
		Status  string  `json:"status" example:"Aprovado" doc:"Retornar se o Aluno foi aprovado ou não"`
	}
}

type AlunoIdRequest struct {
	Id string `query:"id" example:"5" doc:"Id do Aluno"`

	Body struct {
		Nome      string    `json:"nome" example:"Bruno" doc:"Pega o nome do Aluno dentro JSON"`
		Sobrenome string    `json:"sobrenome" example:"Eita" doc:"Pegar o sobrenome do Aluno dentro do JSON /aluno"`
		Notas     []float32 `json:"notas" example:"[9.5, 4.5]" doc:"Pegar as notas do Aluno dentro do JSON /aluno"`
	}
}

type AlunoIdResponse struct {
	Body struct {
		Message string  `json:"message" example:"Aqui está suas notas Bruno: " doc:"Pegar uma mensagem para retornar informações para o Aluno /aluno"`
		Media   float32 `json:"media" example:"10" doc:"Pegar a media do Aluno dentro do JSON"`
		Status  string  `json:"status" example:"Aprovado" doc:"Retornar se o Aluno foi aprovado ou não"`
	}
}

type AlunoDeleteIdRequest struct {
	Id string `path:"id" example:"5" doc:"Id do Aluno"`
}

type AlunoDeleteIdResponse struct {
	Body struct {
		Message string `json:"message" example:"Aluno deletado com sucesso" doc:"Retorna se foi removido ou deu algum tipo de erro"`
	}
}
