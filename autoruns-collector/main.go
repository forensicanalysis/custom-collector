package main

import (
	"github.com/forensicanalysis/artifactcollector/collection"
	"github.com/forensicanalysis/artifactcollector/run"
	"github.com/forensicanalysis/artifactlib/goartifacts"
	"github.com/forensicanalysis/custom-collector/autoruns-collector/assets"
)

//go:generate go get golang.org/x/tools/cmd/goimports github.com/cugu/go-resources/cmd/resources@v0.3.0
//go:generate go mod tidy
//go:generate resources -package assets -output assets/bin.generated.go pack/bin/*

func main() {
	autorunsArtifacts := goartifacts.ArtifactDefinition{
		Name: "Autoruns",
		Sources: []goartifacts.Source{{
			Type: "COMMAND",
			Attributes: goartifacts.Attributes{
				Cmd:  "autorunsc.exe",
				Args: []string{"-x"},
			},
		}},
		SupportedOs: []string{"Windows"},
	}

	config := collection.Configuration{
		Artifacts: []string{"Autoruns"},
	}

	run.Run(&config, []goartifacts.ArtifactDefinition{autorunsArtifacts}, assets.FS)
}
