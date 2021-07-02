package updater

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Eldius/app-updater-go/github"
	"github.com/Eldius/app-updater-go/versions"
)

func GetCurrentBinFile() (string, error) {
	return filepath.Abs(os.Args[0])
}

func FindVersion(version string, releases []versions.Release) versions.Release {
	for _, r := releases {
		if r.GetName() == version {
			return r
		}
	}
	return nil
}

func NewReleases(currentVersion string, repoOwner string, repoName string) {
	var fetcher versions.VersionFetcher = github.NewGithubVersionFetcher(repoOwner, repoName)

	

	fmt.Println(fetcher.FetchReleases())
}