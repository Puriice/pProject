package project

import (
	"errors"
	"log"
	"net/http"

	"github.com/puriice/pProject/internal/types"
	"github.com/puriice/pProject/internal/utils"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Handler struct {
	model types.ProjectModel
}

func NewHandler(model types.ProjectModel) *Handler {
	return &Handler{
		model: model,
	}
}

func (h *Handler) RegisterRoute(router *http.ServeMux) {
	router.HandleFunc("POST /projects", h.handleProjectCreate)
	router.HandleFunc("GET /projects/id/{id}", h.handleProjectQueryByID)
	router.HandleFunc("GET /projects/name/{name}", h.handleProjectQueryByName)
}

func (h *Handler) handleProjectCreate(w http.ResponseWriter, r *http.Request) {
	var payload types.ProjectPayload

	err := utils.ParseJSON(r, &payload)

	if err != nil {
		if errors.Is(err, utils.MissingBody) {
			http.Error(w, "Missing Body", http.StatusBadRequest)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	response, err := h.model.CreateProject(r.Context(), &payload)

	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				w.WriteHeader(http.StatusConflict)
				return
			}
		} else {
			log.Print(err)
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.SendJSON(w, http.StatusCreated, response)
}

func (h *Handler) handleProjectQueryByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project, err := h.model.QueryProjectByID(r.Context(), id)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	utils.SendJSON(w, http.StatusOK, project)
}

func (h *Handler) handleProjectQueryByName(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project, err := h.model.QueryProjectByName(r.Context(), name)

	if err != nil {
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	utils.SendJSON(w, http.StatusOK, project)
}
