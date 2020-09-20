module github.com/forensicanalysis/custom-collector

go 1.14

require (
	github.com/forensicanalysis/artifactcollector v0.15.1-0.20200810070230-a8479d6f47d8
	github.com/forensicanalysis/artifactlib v0.14.0
	github.com/forensicanalysis/artifactsgo v0.6.6
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.4.0
)

replace github.com/forensicanalysis/fslib => github.com/forensicanalysis/fslib v0.14.6
