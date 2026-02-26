package project

type Handler struct {
	model *ProjectModel
}

func NewHandler(model *ProjectModel) *Handler {
	return &Handler{
		model: model,
	}
}
