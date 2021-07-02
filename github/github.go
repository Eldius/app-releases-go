package github

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Eldius/app-updater-go/versions"
)

const (
	githubReleasesEndpointTemplate = "https://api.github.com/repos/%s/%s/releases"
)

type GithubVersionFetcher struct {
	versions.VersionFetcher
	RepoOwner string
	RepoName  string
}

var (
	NewGithubVersionFetcher versions.VersionFetcherFactory = func(repoOwner string, repoName string) versions.VersionFetcher {
		return GithubVersionFetcher{
			RepoOwner: repoOwner,
			RepoName:  repoName,
		}
	}
)

func (f GithubVersionFetcher) VerifyNewReleases(currentVersion string) versions.Release {
	releases, err := f.FetchReleases()
	if err != nil {
		log.Printf("Failed to fetch releases: %s", err.Error())
		return nil
	}

	var latest *GithubRelease
	var current *GithubRelease
	for _, _r := range releases {
		r := _r.(*GithubRelease)
		log.Println("validating version:", r.Name)
		if r.Name == currentVersion {
			log.Println("changing current version:", r.Name)
			current = r
		}
		if latest == nil || latest.PublishedAt.Before(r.PublishedAt) {
			log.Println("changing latest version:", r.Name)
			latest = r
		}
	}

	if current == nil || current.PublishedAt.Before(latest.PublishedAt) {
		return latest
	}
	return nil
}

func (f GithubVersionFetcher) FetchReleases() ([]versions.Release, error) {
	res, err := http.Get(fmt.Sprintf(githubReleasesEndpointTemplate, f.RepoOwner, f.RepoName))
	if err != nil {
		return make([]versions.Release, 0), err
	}
	defer res.Body.Close()

	return parseReleasesResponse(res.Body)
}

func parseReleasesResponse(body io.ReadCloser) ([]versions.Release, error) {
	var releases []*GithubRelease
	err := json.NewDecoder(body).Decode(&releases)
	if err != nil {
		return make([]versions.Release, 0), err
	}

	var result []versions.Release
	for _, r := range releases {
		result = append(result, r)
	}
	return result, nil
}
