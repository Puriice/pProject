package project

import "pProject/internal/types"

type Handler struct {
	model *types.ProjectModel
}

func NewHandler(model *types.ProjectModel) *Handler {
	return &Handler{
		model: model,
	}
}
