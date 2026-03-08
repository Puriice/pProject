package project

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/puriice/pProject/pkg/model"
)

const slashAPISlashVersionOneSlashProjectLength = 16

type ProjectService struct {
	url string
}

func NewService(url string) *ProjectService {
	var sb strings.Builder

	sb.Grow(len(url) + slashAPISlashVersionOneSlashProjectLength)
	sb.WriteString(url)

	if !strings.HasSuffix(url, "/") {
		sb.WriteRune('/')
	}

	sb.WriteString("api/v1/projects")

	return &ProjectService{
		url: sb.String(),
	}
}

const slashIDSlashUUIDLength = 40

func (s *ProjectService) GetProjectInfo(id string) (*model.Project, error) {
	var url strings.Builder

	url.Grow(len(s.url) + slashIDSlashUUIDLength)

	url.WriteString(s.url)
	url.WriteString("/id/")
	url.WriteString(id)

	resp, err := http.Get(url.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var project *model.Project = new(model.Project)

	err = json.NewDecoder(resp.Body).Decode(project)

	return project, err
}
