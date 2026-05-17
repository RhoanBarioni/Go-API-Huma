package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/RhoanBarioni/Go-API-Huma/internal/models"
)

func GetAlunosDB(ctx context.Context, db *sql.DB) ([]models.Aluno, error) {
	alunos := []models.Aluno{}
	rows, err := db.QueryContext(ctx, "SELECT * FROM alunos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		aluno := models.Aluno{}
		rows.Scan(&aluno.Id, &aluno.Nome, &aluno.Sobrenome, &aluno.Media)
		alunos = append(alunos, aluno)
	}

	return alunos, nil
}

func GetAlunoIDNameDB(ctx context.Context, db *sql.DB, id string) (models.Aluno, error) {
	aluno := models.Aluno{}
	err := db.QueryRowContext(ctx, "select * from alunos where id = ?", id).Scan(
		&aluno.Id,
		&aluno.Nome,
		&aluno.Sobrenome,
		&aluno.Media,
	)
	if err != nil {
		log.Println("erro ao buscar aluno:", err)
		return aluno, err
	}

	return aluno, nil
}

func PostAlunoDB(ctx context.Context, db *sql.DB, aluno *models.Aluno) error {
	_, err := db.ExecContext(ctx, "insert into alunos (nome, sobrenome, media) values (?, ?, ?)", aluno.Nome, aluno.Sobrenome, aluno.Media)
	if err != nil {
		return err
	}

	return nil
}

func PutAlunoDB(ctx context.Context, db *sql.DB, aluno *models.Aluno) error {
	_, err := db.ExecContext(ctx, "update alunos set nome = ?, sobrenome = ?, media = ? where id = ?", aluno.Nome, aluno.Sobrenome, aluno.Media, aluno.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAlunoId(ctx context.Context, db *sql.DB, aluno *models.Aluno) error {
	_, err := db.ExecContext(ctx, "DELETE from alunos where id = ?", aluno.Id)
	if err != nil {
		return err
	}
	return nil
}