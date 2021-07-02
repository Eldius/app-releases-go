package versions

import "time"

type Release interface {
	GetName() string
	GetPublishedAt() time.Time
	GetArtifacts() []Artifact
	String() string
}

type Artifact interface {
	GetName() string
	GetArtifactURL() string
	String() string
}

type VersionFetcher interface {
	VerifyNewReleases(current string) Release
	FetchReleases() ([]Release, error)
}

type VersionFetcherFactory func(repoOwner string, repoName string) VersionFetcher
