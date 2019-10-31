package github

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Repo represents a repository.
type Repo struct {
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	HTMLURL       string `json:"html_url"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
}

type repoOrError struct {
	Repo
	Message string `json:"message"`
}

func getRepoPath(repo string) string {
	return newPath("/repos/" + repo).
		String()
}

func (c *client) GetRepo(path string) (*Repo, error) {
	res, err := c.get(c.url(getRepoPath(path)))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r repoOrError
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	if r.Message != "" {
		return nil, fmt.Errorf("%s: %s", r.Message, path)
	}

	return &r.Repo, nil
}

// UpdateRepoParams represents a parameter on updating a repository.
type UpdateRepoParams struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	Private       bool   `json:"private"`
	DefaultBranch string `json:"default_branch"`
}

func updateRepoPath(repo string) string {
	return newPath("/repos/" + repo).
		String()
}

// UpdateRepo updates a repository.
func (c *client) UpdateRepo(path string, params *UpdateRepoParams) (*Repo, error) {
	bs, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(bs)
	res, err := c.patch(c.url(updateRepoPath(path)), body)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var r repoOrError
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	if r.Message != "" {
		return nil, fmt.Errorf("%s: %s", r.Message, path)
	}

	return &r.Repo, nil
}
