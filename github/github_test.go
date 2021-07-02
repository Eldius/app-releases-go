package github

import (
	"fmt"
	"os"
	"testing"

	"gopkg.in/h2non/gock.v1"
)

const (
	url = "https://api.github.com"
)

func TestParseReleasesResponse(t *testing.T) {

	f, err := os.Open("sample_data/releases_sample.json")
	if err != nil {
		t.Errorf("Failed to open sample file: %s", err.Error())
	}

	r, err := parseReleasesResponse(f)
	if err != nil {
		t.Errorf("Failed to parse release payload: %s", err.Error())
	}

	if len(r) != 3 {
		t.Errorf("Releases list lenth must have 3 elements, but has %d", len(r))
	}
}

func TestFetchReleases(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "homeJSimpson"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("sample_data/releases_sample.json")

	fetcher := NewGithubVersionFetcher(repoOwner, repoName)
	releases, err := fetcher.FetchReleases()
	if err != nil {
		t.Errorf("Failed to fetch releases: %s", err.Error())
	}

	if len(releases) != 3 {
		t.Errorf("Must return 3 releases, but returned '%d'", len(releases))
	}
}

func TestVerifyNewReleasesWithTwoNewReleases(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "homeJSimpson"
	currentRelease := "v0.0.1"
	latestRelease := "v0.0.3"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("sample_data/releases_sample.json")

	fetcher := NewGithubVersionFetcher(repoOwner, repoName)
	latest := fetcher.VerifyNewReleases(currentRelease)
	if latest == nil {
		t.Errorf("Should return a non nil release")
	}

	if latest != nil && latest.GetName() != latestRelease {
		t.Errorf("Latest release must be 'v0.0.3', but was '%s'", latest.GetName())
	}
}

func TestVerifyNewReleasesWithOneNewRelease(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "homeJSimpson"
	currentRelease := "v0.0.2"
	latestRelease := "v0.0.3"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("sample_data/releases_sample.json")

	fetcher := NewGithubVersionFetcher(repoOwner, repoName)
	latest := fetcher.VerifyNewReleases(currentRelease)
	if latest == nil {
		t.Errorf("Should return a non nil release")
	}

	if latest != nil && latest.GetName() != latestRelease {
		t.Errorf("Latest release must be 'v0.0.3', but was '%s'", latest.GetName())
	}
}

func TestVerifyNewReleasesWithNoNewRelease(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "homeJSimpson"
	currentRelease := "v0.0.3"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("sample_data/releases_sample.json")

	fetcher := NewGithubVersionFetcher(repoOwner, repoName)
	latest := fetcher.VerifyNewReleases(currentRelease)
	if latest != nil {
		t.Errorf("Should return a nil release")
	}
}

func TestVerifyNewReleasesWithNoCurrentRelease(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "homeJSimpson"
	currentRelease := ""
	latestRelease := "v0.0.3"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("sample_data/releases_sample.json")

	fetcher := NewGithubVersionFetcher(repoOwner, repoName)
	latest := fetcher.VerifyNewReleases(currentRelease)
	if latest == nil {
		t.Errorf("Should return a non nil release")
	}

	if latest != nil && latest.GetName() != latestRelease {
		t.Errorf("Latest release must be 'v0.0.3', but was '%s'", latest.GetName())
	}
}
