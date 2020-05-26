// Copyright (c) 2020 Siemens AG
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
// the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
// FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
// COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
// IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
// CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
//
// Author(s): Jonas Plum

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
