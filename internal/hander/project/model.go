package project

import "github.com/jackc/pgx/v5/pgxpool"

type ProjectModel struct {
	db *pgxpool.Pool
}

func NewModel(db *pgxpool.Pool) *ProjectModel {
	return &ProjectModel{
		db: db,
	}
}

func (m *ProjectModel) createProject() {

}
