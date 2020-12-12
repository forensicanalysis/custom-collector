module github.com/forensicanalysis/custom-collector

go 1.14

require (
	github.com/forensicanalysis/artifactcollector v0.15.1
	github.com/forensicanalysis/artifactlib v0.14.2
	github.com/forensicanalysis/artifactsgo v0.6.6
	github.com/pkg/errors v0.9.1
	github.com/spf13/afero v1.4.1
)

replace github.com/forensicanalysis/fslib => github.com/forensicanalysis/fslib v0.14.7
