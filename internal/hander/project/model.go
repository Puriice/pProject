package project

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/puriice/pProject/internal/types"
)

var (
	ErrNotFound = errors.New("Row not found")
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
		"SELECT id, name, description, picture FROM projects WHERE id = $1;",
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
		"SELECT id, name, description, picture FROM projects WHERE name = $1;",
		name,
	).Scan(&project.ID, &project.Name, &project.Description, &project.Picture)

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (m *Model) UpdateProject(context context.Context, id string, payload *types.ProjectPayload) error {
	q := make([]string, 0, 3)
	args := make([]any, 0, 4)
	argn := 1

	if payload.Name != nil {
		q = append(q, fmt.Sprintf("name = $%d", argn))
		args = append(args, payload.Name)
		argn++
	}

	if payload.Description != nil {
		q = append(q, fmt.Sprintf("description = $%d", argn))
		args = append(args, payload.Description)
		argn++
	}

	if payload.Picture != nil {
		q = append(q, fmt.Sprintf("picture = $%d", argn))
		args = append(args, payload.Picture)
		argn++
	}

	sql := fmt.Sprintf(
		"UPDATE projects SET %s WHERE id = $%d;",
		strings.Join(q, ", "),
		argn,
	)

	args = append(args, id)

	cmdTag, err := m.db.Exec(context, sql, args...)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}

func (m *Model) DeleteProject(context context.Context, id string) error {
	cmdTag, err := m.db.Exec(context, "DELETE FROM projects WHERE id = $1;", id)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return ErrNotFound
	}

	return nil
}
