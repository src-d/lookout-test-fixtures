package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/google/go-github/github"
	fixtures "github.com/src-d/lookout-test-fixtures"
)

func main() {
	tr := github.BasicAuthTransport{
		Username: os.Getenv("GITHUB_USER"),
		Password: os.Getenv("GITHUB_TOKEN"),
	}

	client := github.NewClient(tr.Client())

	fmt.Println("Saving JSON fixtures...\n")

	fixtureList := fixtures.GetAll()

	for _, fixture := range fixtureList {
		filenamePR := fixture.FilenamePR(fixture.CurrentRevision)

		if _, err := os.Stat(filenamePR); err == nil {
			fmt.Printf("%s exists, skip\n", filenamePR)
			continue
		}

		pr, _, err := client.PullRequests.Get(context.Background(),
			fixture.URL.Owner, fixture.URL.Repo, fixture.URL.Number)
		if err != nil {
			log.Fatal(err)
		}

		data, err := json.MarshalIndent(pr, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(
			filenamePR,
			data, 0644)
		if err != nil {
			log.Fatal(err)
		}

		cc, _, err := client.Repositories.CompareCommits(context.Background(),
			fixture.URL.Owner, fixture.URL.Repo,
			pr.GetBase().GetSHA(), pr.GetHead().GetSHA())
		if err != nil {
			log.Fatal(err)
		}

		data, err = json.MarshalIndent(cc, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(
			fixture.FilenameCC(fixture.CurrentRevision),
			data, 0644)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n  PR %s/%s/#%d: %s\n", fixture.Name, fixture.URL.Owner, fixture.URL.Repo, fixture.URL.Number, pr.GetTitle())
	}

	fmt.Println("\n==> Don't forget to run 'go-bindata -modtime 1536310226 -pkg fixtures fixtures/'")
}
