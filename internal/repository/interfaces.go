package repository

import (
	"context"

	"github.com/puriice/pProject/internal/types"
	"github.com/puriice/pProject/pkg/model"
)

type ProjectRepository interface {
	CreateProject(context context.Context, payload *types.ProjectPayload) (*model.Project, error)
	QueryProjectByID(context context.Context, id string) (*model.Project, error)
	QueryProjectByName(context context.Context, name string) (*model.Project, error)
	UpdateProject(context context.Context, id string, payload *types.ProjectPayload) error
	DeleteProject(context context.Context, id string) error
}
