package handlers

import (
	"context"
	"fmt"

	"github.com/RhoanBarioni/Go-API-Huma/internal/contracts"
	service "github.com/RhoanBarioni/Go-API-Huma/internal/service"
)

// Esses caras são handlers.

func GetAlunos(ctx context.Context, input *contracts.GetAlunosResquest) (*contracts.GetAlunosResponse, error) {
	res := &contracts.GetAlunosResponse{}
	res.Body.Nome = input.Nome
	res.Body.Message = "Olá, " + input.Nome
	return res, nil
}

// func GetAlunosDB(db *sql.DB) func(context.Context, *contracts.GetAlunosRequest) (*contracts.GetAlunosResponse, error) {
// 	return func(ctx context.Context, input *contracts.GetAlunosRequest) (*contracts.GetAlunosResponse, error) {
// 		alunos, err := repository.GetAlunosDB(ctx, db)
// 		if err != nil {
// 			return nil, err
// 		}
// 		res := &contracts.GetAlunosResponse{}
// 		res.Body.Alunos = alunos
// 		if input.Nome != "" {
// 			res.Body.Message = fmt.Sprintf("Resultado para a busca %s", input.Nome)
// 		} else {
// 			res.Body.Message = "Listagem completa de alunos"
// 		}
// 		return res, nil
// 	}
// }

func GetAlunosId(ctx context.Context, input *contracts.GetAlunosIdResquest) (*contracts.GetAlunosIdResponse, error) {

	res := &contracts.GetAlunosIdResponse{}
	res.Body.Nome = "Bruno"
	res.Body.Message = "Olá, " + res.Body.Nome + ". Seu ID: " + input.Id
	return res, nil
}

func CreateAluno(ctx context.Context, input *contracts.AlunoRequest) (*contracts.AlunoReponse, error) {
	res := &contracts.AlunoReponse{}
	nome := input.Body.Nome
	// dentro do parametro, ele so vai criar a var temporaria dentro da funcao
	media := service.CalcMedia(input.Body.Notas)

	res.Body.Media = float32(media)

	if media >= 7 {
		res.Body.Status = "Aprovado"
	} else {
		res.Body.Status = "Reprovado"
	}

	res.Body.Message = fmt.Sprintf("Olá, %v. Aqui está as suas notas: %v \n Sua média final é: %.2f \n você está %s", nome, input.Body.Notas, media, res.Body.Status)

	return res, nil
}

func UpdateAluno(ctx context.Context, input *contracts.AlunoIdRequest) (*contracts.AlunoIdResponse, error) {
	res := &contracts.AlunoIdResponse{}

	media := service.CalcMedia(input.Body.Notas)

	res.Body.Media = float32(media)

	if media >= 7 {
		res.Body.Status = "Aprovado"
	} else {
		res.Body.Status = "Reprovado"
	}

	res.Body.Message = fmt.Sprintf("Aluno ID: %s atualizado: %s | Notas: %v | Média: %.2f | Status: %s", input.Id, input.Body.Name, input.Body.Notas, media, res.Body.Status)

	return res, nil
}

func DeleteAlunoId(ctx context.Context, input *contracts.AlunoDeleteIdRequest) (*contracts.AlunoDeleteIdResponse, error) {
	res := &contracts.AlunoDeleteIdResponse{}
	alunoId := "5"
	if alunoId != input.Id {
		res.Body.Message = "Aluno não encontrado ou não existe"
		return res, nil
	}

	// usar a funcao do go "delete" para deletar

	res.Body.Message = fmt.Sprintf("Aluno com o ID: %v foi deletado com sucesso", input.Id)

	return res, nil
}
