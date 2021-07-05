package updater

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Eldius/app-releases-go/github"
	"github.com/Eldius/app-releases-go/versions"
)

var (
	GithubVersionAPI versions.VersionAPI = github.GithubVersionAPI
)

func GetCurrentBinFile() (string, error) {
	return filepath.Abs(os.Args[0])
}

func FindVersion(version string, releases []versions.Release) versions.Release {
	for _, r := range releases {
		if r.GetName() == version {
			return r
		}
	}
	return nil
}

func ListReleases(repoOwner string, repoName string, api versions.VersionAPI) ([]versions.Release, error) {
	fetcher, err := getFetcher(repoOwner, repoName, api)
	if err != nil {
		return make([]versions.Release, 0), err
	}

	return fetcher.FetchReleases()
}

func NewReleases(currentVersion string, repoOwner string, repoName string, api versions.VersionAPI) error {
	fetcher, err := getFetcher(repoOwner, repoName, api)
	if err != nil {
		log.Println("Failed to get API fetcher:", err.Error())
		return err
	}

	fmt.Println(fetcher.FetchReleases())
	return nil
}

func getFetcher(repoOwner string, repoName string, api versions.VersionAPI) (versions.VersionFetcher, error) {
	switch api {
	case github.GithubVersionAPI:
		return github.NewGithubVersionFetcher(repoOwner, repoName), nil
	default:
		return nil, fmt.Errorf("Invalid API option: '%s'", api)
	}

}
