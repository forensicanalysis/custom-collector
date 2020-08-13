# Custom Artifactcollector

The [artifactcollector](https://github.com/forensicanalysis/artifactcollector) can be customized to your needs. This repository shows some examples for this.

**Note:** The **master** branch might not be stable. Please use the latest released version (e.g. `git checkout v0.1.0`).

## zip-collector

The [zip-collector](zip-collector/main.go) is an example that stores the collected files in a zip file instead of the database.

## autoruns-collector

The [autoruns-collector](autoruns-collector/main.go) is an example that includes the autoruns tool into the artifactcollector and collects the printed results. 

**Note:** You need to replace the autorunsc.exe with a real one to make this example work. 

