package serviceapi

import (
	"context"
	"fmt"
)

type GreetingBodyRequest struct {
	// armazenar o nome dentro da URL
	// é oq o usuario manda para o servidor
	Name string `path:"name" example:"Bruno" doc:"Nome do aluno"`
}

type GreetingBodyResponse struct {
	// é oq o servidor vai devolver para o usuario
	Body struct {
		Name    string `json:"nome" example:"Bruno" doc:"Pega o nome do Aluno dentro JSON"`
		Message string `json:"message" example:"Olá Bruno" doc:"Mensagem com o nome do Aluno"`
		// Notas []float64 `json:"notas" example:"9.5, 5" doc:"Notas do Aluno dentro do JSON"`
	}
}

func GreetingEndpoint(ctx context.Context, input *GreetingBodyRequest) (*GreetingBodyResponse, error) {
	// criar uma var para armazenar os valores da struct de RESPONSE
	res := &GreetingBodyResponse{}
	// pegar o valor dentro do input e jogar para o response
	res.Body.Name = input.Name
	res.Body.Message = "Olá " + input.Name

	return res, nil
}

type AlunoRequest struct {
	Body struct{
		Nome  string    `json:"nome" example:"Bruno" doc:"Pegar o nome do Aluno dentro do JSON /aluno"`
		Notas []float64 `json:"notas" example:"[9.5, 4.5]" doc:"Pegar as notas do Aluno dentro do JSON /aluno"`
	}
}

type AlunoReponse struct {
	Body struct {
		Message string  `json:"message" example:"Aqui está suas notas Bruno: " doc:"Pegar uma mensagem para retornar informações para o Aluno /aluno"`
		Media   float32 `json:"media" example:"10" doc:"Pegar a media do Aluno dentro do JSON"`
		Status  string  `json:"status" example:"Aprovado" doc:"Retornar se o Aluno foi aprovado ou não"`
	}
}

func AlunoEndpoint(ctx context.Context, input *AlunoRequest) (*AlunoReponse, error) {
	fmt.Printf("INPUT: %+v\n", input)
	res := &AlunoReponse{}
	nome := input.Body.Nome
	media := func(notas []float64) float64 {
		if len(input.Body.Notas) == 0 {
			return 0
		}

		var nota float64
		for _, n := range notas {
			nota += n
		}
		return nota / float64(len(notas))
	}(input.Body.Notas)

	res.Body.Media = float32(media)

	if media >= 7 {
		res.Body.Status = "Aprovado"
	} else {
		res.Body.Status = "Reprovado"
	}

	res.Body.Message = fmt.Sprintf("Olá, %v. Aqui está as suas notas: %v \n Sua média final é: %.2f \n você está %s", nome, input.Body.Notas, media, res.Body.Status)

	return res, nil
}
