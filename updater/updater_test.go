package updater

import (
	"fmt"
	"testing"

	"github.com/Eldius/app-releases-go/github"
	"gopkg.in/h2non/gock.v1"
)

const (
	url = "https://api.github.com"
)

func TestListReleases(t *testing.T) {
	defer gock.Off() // Flush pending mocks after test execution

	repoOwner := "maxPower"
	repoName := "my_awsome_project"

	gock.New(url).
		Get(fmt.Sprintf("/repos/%s/%s/releases", repoOwner, repoName)).
		Reply(200).
		File("../github/sample_data/releases_sample.json")

	results, err := ListReleases(repoOwner, repoName, github.GithubVersionAPI)
	if err != nil {
		t.Errorf("Failed to fetch releases: %s", err.Error())
	}

	if len(results) != 3 {
		t.Errorf("Must return '3' releases, but returned '%d'", len(results))
	}
}
