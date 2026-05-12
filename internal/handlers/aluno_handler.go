package handlers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/RhoanBarioni/Go-API-Huma/internal/contracts"
	"github.com/RhoanBarioni/Go-API-Huma/internal/models"
	"github.com/RhoanBarioni/Go-API-Huma/internal/repository"
	service "github.com/RhoanBarioni/Go-API-Huma/internal/service"
)

type Handler struct {
	DB *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// Esses caras são handlers.

func (h *Handler) GetAlunos(ctx context.Context, input *contracts.GetAlunosResquest) (*contracts.GetAlunosResponse, error) {
	alunosDB, err := repository.GetAlunosDB(ctx, h.DB)
	if err != nil {
		return nil, err
	}
	res := &contracts.GetAlunosResponse{}
	res.Body.Alunos = alunosDB
	res.Body.Message = "Consulta feita no Banco de Dados"
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

func (h *Handler) GetAlunosId(ctx context.Context, input *contracts.GetAlunosIdResquest) (*contracts.GetAlunosIdResponse, error) {
	alunosDB, err := repository.GetAlunoIDNameDB(ctx, h.DB, input.Id)
	if err != nil {
		return nil, err
	}
	res := &contracts.GetAlunosIdResponse{}
	res.Body.Nome = alunosDB.Nome
	res.Body.Message = fmt.Sprintf("Olá, %s %s (ID: %d). Sua Média é: %0.2f", alunosDB.Nome, alunosDB.Sobrenome, alunosDB.Id, alunosDB.Media)
	return res, nil
}

func (h *Handler) CreateAluno(ctx context.Context, input *contracts.AlunoRequest) (*contracts.AlunoResponse, error) {
	res := &contracts.AlunoResponse{}
	// dentro do parametro, ele so vai criar a var temporaria dentro da funcao
	media := service.CalcMedia(input.Body.Notas)

	res.Body.Media = float32(media)

	if media >= 7 {
		res.Body.Status = "Aprovado"
	} else {
		res.Body.Status = "Reprovado"
	}

	aluno := models.Aluno{
		Nome:      input.Body.Nome,
		Sobrenome: input.Body.Sobrenome,
		Media:     float64(res.Body.Media),
	}

	err := repository.PostAlunoDB(ctx, h.DB, &aluno)
	if err != nil {
		return nil, err
	}

	res.Body.Message = fmt.Sprintf("Olá, %v %v. Aqui está as suas notas: %v \n Sua média final é: %.2f \n você está %s", aluno.Nome, aluno.Sobrenome, input.Body.Notas, aluno.Media, res.Body.Status)

	return res, nil
}

func (h *Handler) UpdateAluno(ctx context.Context, input *contracts.AlunoIdRequest) (*contracts.AlunoIdResponse, error) {
	res := &contracts.AlunoIdResponse{}

	media := service.CalcMedia(input.Body.Notas)

	res.Body.Media = float32(media)

	if media >= 7 {
		res.Body.Status = "Aprovado"
	} else {
		res.Body.Status = "Reprovado"
	}

	aluno := models.Aluno{
		Id: input.Id,
		Nome:      input.Body.Nome,
		Sobrenome: input.Body.Sobrenome,
		Media:     float64(res.Body.Media),
	}

	err := repository.PutAlunoDB(ctx, h.DB, &aluno)
	if err != nil{
		return nil, err
	}

	res.Body.Message = fmt.Sprintf("Aluno ID: %s atualizado: %s %s | Notas: %v | Média: %.2f | Status: %s", aluno.Id, aluno.Nome, aluno.Sobrenome, input.Body.Notas, aluno.Media, res.Body.Status)

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
