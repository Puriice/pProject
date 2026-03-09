package pproject_test

import (
	"testing"

	"github.com/puriice/pproject/pkg/sdk/pproject"
)

func TestGetProjectInfo(t *testing.T) {
	projectService := pproject.NewService("http://localhost:8081", nil)

	project, err := projectService.GetProjectInfo("b5e50b5a-9234-44e3-af27-054b88b20b3a")

	if err != nil {
		t.Error(err)
	}

	t.Log(project)
}
