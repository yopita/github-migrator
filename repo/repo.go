package repo

import (
	"github.com/itchyny/github-migrator/github"
)

// Repo represents a GitHub repository.
type Repo interface {
	Get() (*github.Repo, error)
	ListIssues() github.Issues
}

// New creates a new Repo.
func New(cli github.Client, path string) Repo {
	return &repo{cli: cli, path: path}
}

type repo struct {
	cli  github.Client
	path string
}