package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v28/github"
)

var client *github.Client
var token string
var org string
var includeForks bool
var includeLanguage bool
var listHTTPSURL bool

func buildClient() *github.Client {
	ctx := context.Background()
	if token == "%unset%" {
		// Unauthenticated client
		client = github.NewClient(nil)
	} else {
		tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		tokenClient := oauth2.NewClient(ctx, tokenSource)
		client = github.NewClient(tokenClient)
	}

	return client
}

func pickURL(repo *github.Repository) *string {
	if listHTTPSURL {
		return repo.CloneURL
	} else {
		return repo.SSHURL
	}
}

func pickURLs(repos []*github.Repository) []*string {
	resp := make([]*string, len(repos))
	for i, repo := range repos {
		resp[i] = pickURL(repo)
	}

	return resp
}

func pickMany(repos []*github.Repository) []*string {
	resp := make([]*string, len(repos))
	for i, repo := range repos {
		m := &map[string]*string{
			"url":      pickURL(repo),
			"language": repo.Language,
		}
		bytes, err := json.Marshal(m)
		if err != nil {
			fmt.Printf("Error while marshalling JSON: %s\n", err)
			os.Exit(1)
		}
		json := string(bytes)
		resp[i] = &json
	}

	return resp
}

func getOrgRepos() []*github.Repository {
	var allRepos []*github.Repository
	ctx := context.Background()

	opts := &github.RepositoryListByOrgOptions{
		Type: "all",
		ListOptions: github.ListOptions{
			PerPage: 5000,
		},
	}

	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opts)

		if err != nil {
			fmt.Printf("Error: %s", err)
			os.Exit(1)
		}

		allRepos = append(allRepos, repos...)

		if resp.NextPage == 0 {
			break
		}

		opts.Page = resp.NextPage
	}

	return allRepos
}

func filterRepos(repos []*github.Repository) []*github.Repository {
	if includeForks {
		return repos
	} else {
		filteredRepos := make([]*github.Repository, 0)
		for _, repo := range repos {
			if *repo.Fork {
				continue
			}
			filteredRepos = append(filteredRepos, repo)
		}
		return filteredRepos
	}
}

func formatResponse(repos []*github.Repository) []*string {
	filteredRepos := filterRepos(repos)

	if includeLanguage {
		return pickMany(filteredRepos)
	} else {
		return pickURLs(filteredRepos)
	}
}

func main() {
	flag.StringVar(&token, "token", "%unset%", "GitHub Personal Access token")
	flag.StringVar(&org, "org", "github", "GitHub organization to list repos for")
	flag.BoolVar(&includeForks, "forks", false, "Include forks of repos in output")
	flag.BoolVar(&includeLanguage, "language", false, "Include language in output")
	flag.BoolVar(&listHTTPSURL, "https", false, "List HTTPS URLs instead of SSH")
	flag.Parse()

	client = buildClient()
	repos := getOrgRepos()
	response := formatResponse(repos)

	for _, url := range response {
		fmt.Println(*url)
	}
}
