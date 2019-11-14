package repo

import "github.com/itchyny/github-migrator/github"

// Repo represents a GitHub repository.
type Repo interface {
	NewPath(string) Repo
	Get() (*github.Repo, error)
	ListMembers() github.Members
	Update(*github.UpdateRepoParams) (*github.Repo, error)
	ListLabels() github.Labels
	CreateLabel(*github.CreateLabelParams) (*github.Label, error)
	UpdateLabel(string, *github.UpdateLabelParams) (*github.Label, error)
	ListIssues() github.Issues
	GetIssue(int) (*github.Issue, error)
	ListComments(int) github.Comments
	ListPullReqs() github.PullReqs
	GetPullReq(int) (*github.PullReq, error)
	ListPullReqCommits(int) github.Commits
	GetDiff(string) (string, error)
	GetCompare(string, string) (string, error)
	ListReviews(int) github.Reviews
	ListReviewComments(int) github.ReviewComments
	Import(*github.Import) (*github.ImportResult, error)
	GetImport(int) (*github.ImportResult, error)
}

// New creates a new Repo.
func New(cli github.Client, path string) Repo {
	return &repo{cli: cli, path: path}
}

type repo struct {
	cli  github.Client
	path string
}

// NewPath creates a new Repo with the same client.
func (r *repo) NewPath(path string) Repo {
	return New(r.cli, path)
}
