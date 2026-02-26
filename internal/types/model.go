package types

import "context"

type ProjectModel interface {
	CreateProject(context context.Context, payload *ProjectPayload) (*Project, error)
}
