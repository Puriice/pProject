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
	id := new(string)

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
		ID:          id,
		Name:        payload.Name,
		Description: payload.Description,
		Picture:     payload.Picture,
	}, nil
}

func (m *Model) QueryProjectByID(context context.Context, id string) (*types.Project, error) {
	project := new(types.Project)

	err := m.db.QueryRow(
		context,
		"SELECT * FROM projects WHERE id = $1;",
		id,
	).Scan(&project.ID, &project.Name, &project.Description, &project.Picture)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m *Model) QueryProjectByName(context context.Context, name string) (*types.Project, error) {
	project := new(types.Project)

	err := m.db.QueryRow(
		context,
		"SELECT * FROM projects WHERE name = $1;",
		name,
	).Scan(&project.ID, &project.Name, &project.Description, &project.Picture)

	if err != nil {
		return nil, err
	}

	return project, nil
}
