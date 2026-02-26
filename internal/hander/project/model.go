package project

import (
	"context"
	"pProject/internal/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Model struct {
	db *pgxpool.Pool
}

func NewModel(db *pgxpool.Pool) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) CreateProject(context context.Context, payload *types.ProjectPayload) (*types.Project, error) {
	var id string

	err := m.db.QueryRow(
		context,
		"INSERT INTO projects (name, description, picture) VALUES ($1, $2, $3) RETURNING id",
		payload.Name,
		payload.Description,
		payload.Picture,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &types.Project{
		ID:          &id,
		Name:        payload.Name,
		Description: payload.Description,
		Picture:     payload.Picture,
	}, nil
}
