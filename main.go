package fixtures

import (
	"strconv"
	"strings"

	"github.com/src-d/lookout"
)

var fixtures = Fixtures{
	{
		// FIXME, use another branch and create a PR for it
		Name: "too-long-line",
		CommitRevision: lookout.CommitRevision{
			Base: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/master",
				Hash:                  "c6e7417ff3b84560f44fe940876494f58e6d68f1",
			},
			Head: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/master",
				Hash:                  "6a92946068897d0a6f6ffa6457f889163dcc51b5",
			},
		},
		PR: PullRequest{
			URL: "",
		},
	},
	{
		Name: "file-delete",
		CommitRevision: lookout.CommitRevision{
			Base: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/master",
				Hash:                  "c6e7417ff3b84560f44fe940876494f58e6d68f1",
			},
			Head: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/remove_file",
				Hash:                  "d0b400aab40eff88066fbe9e27f8aa64f5530538",
			},
		},
		PR: PullRequest{
			URL: "https://github.com/src-d/lookout-test-fixtures/pull/2",
		},
	},
	{
		Name: "file-rename",
		CommitRevision: lookout.CommitRevision{
			Base: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/master",
				Hash:                  "6a92946068897d0a6f6ffa6457f889163dcc51b5",
			},
			Head: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/rename_file",
				Hash:                  "6221d2fe0bc2148debfa8d3c8c92b8c15451920d",
			},
		},
		PR: PullRequest{
			URL: "https://github.com/src-d/lookout-test-fixtures/pull/3",
		},
	},
	{
		Name: "bblfsh-unknown-language",
		CommitRevision: lookout.CommitRevision{
			Base: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/master",
				Hash:                  "6a92946068897d0a6f6ffa6457f889163dcc51b5",
			},
			Head: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/bblfsh_unknown_lang",
				Hash:                  "8bc4f2b51be093d261a5ea58e390bc94c1ca3401",
			},
		},
		PR: PullRequest{
			URL: "https://github.com/src-d/lookout-test-fixtures/pull/4",
		},
	},
	{
		Name: "not-ff-merge",
		CommitRevision: lookout.CommitRevision{
			Base: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/i197-base",
				Hash:                  "1f5664bfe6a04a33b6de17a4df0b051d7a43b918",
			},
			Head: lookout.ReferencePointer{
				InternalRepositoryURL: "https://github.com/src-d/lookout-test-fixtures",
				ReferenceName:         "refs/heads/i197-head",
				Hash:                  "5fe468b62112e69bae390051e990271f7b1cc294",
			},
		},
		PR: PullRequest{
			URL: "https://github.com/src-d/lookout-test-fixtures/pull/1",
		},
	},
}

// PullRequest is a struct with information about pull request
type PullRequest struct {
	URL string
	// add recorded responses here later?
}

// Fixture is struct for a test case
type Fixture struct {
	Name           string
	CommitRevision lookout.CommitRevision
	PR             PullRequest
}

// ReviewEvent creates fake review event with information from fixture
func (f *Fixture) ReviewEvent(provider string) lookout.ReviewEvent {
	var number uint32 = 1
	if f.PR.URL != "" {
		parts := strings.Split(f.PR.URL, "/")
		n, _ := strconv.ParseUint(parts[len(parts)-1], 10, 32)
		number = uint32(n)
	}

	return lookout.ReviewEvent{
		Provider:       provider,
		InternalID:     f.Name,
		IsMergeable:    true,
		Number:         number,
		CommitRevision: f.CommitRevision,
	}
}

// PushEvent creates fake push event with information from fixture
func (f *Fixture) PushEvent(provider string) lookout.PushEvent {
	return lookout.PushEvent{
		Provider:       provider,
		InternalID:     f.Name,
		CommitRevision: f.CommitRevision,
	}
}

// Fixtures is a list of fixtures
type Fixtures []*Fixture

// GetByName returns fixture by name
func (g Fixtures) GetByName(name string) *Fixture {
	for _, f := range g {
		if f.Name == name {
			return f
		}
	}

	return nil
}

// GetAll returns all fixtures
func GetAll() []*Fixture {
	return fixtures
}

// GetByName returns fixture by name
func GetByName(name string) *Fixture {
	return fixtures.GetByName(name)
}
