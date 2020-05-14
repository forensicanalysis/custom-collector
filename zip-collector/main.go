package main

import (
	"log"
	"os"

	"github.com/forensicanalysis/artifactcollector/collection"
	"github.com/forensicanalysis/artifactcollector/run"
	"github.com/forensicanalysis/artifactsgo"
	"github.com/forensicanalysis/custom-collector/zip-collector/zipwrite"
)

func main() {
	f, err := os.Create("collection.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fs := zipwrite.New(f)
	defer fs.Close()

	config := collection.Configuration{
		Artifacts: []string{"WindowsEventLogs", "WindowsPrefetchFiles", "MacOSHostsFile"},
		FS:        fs,
	}

	run.Run(&config, artifactsgo.Artifacts, nil)
}
