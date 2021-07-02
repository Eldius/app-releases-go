package main

import (
	"fmt"

	"github.com/Eldius/app-updater-go/github"
	"github.com/Eldius/app-updater-go/versions"
)

func main() {

	var artifact versions.Artifact = &github.Asset{}
	var release versions.Release = &github.GithubRelease{}
	fmt.Println(artifact)
	fmt.Println(release)

	var factory versions.VersionFetcherFactory = github.NewGithubVersionFetcher

	var fetcher versions.VersionFetcher = factory("eldius", "gvm")
	var releases []versions.Release
	var err error
	releases, err = fetcher.FetchReleases()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(&releases)
}
