package types

import (
	"context"
)

type ProjectModel interface {
	CreateProject(context context.Context, payload *ProjectPayload) (*Project, error)
	QueryProjectByID(context context.Context, id string) (*Project, error)
	QueryProjectByName(context context.Context, name string) (*Project, error)
}
