// Copyright (c) 2020 Jonas Plum
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
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/forensicanalysis/artifactcollector/collection"
	"github.com/forensicanalysis/artifactcollector/run"
	"github.com/forensicanalysis/artifactsgo"
)

func main() {
	config := collection.Configuration{
		Artifacts: []string{"WindowsEventLogs", "WindowsPrefetchFiles", "MacOSHostsFile", "NTFSMFTFiles"},
	}

	createdCollection := run.Run(&config, artifactsgo.Artifacts, nil)

	zip, err := os.Open(createdCollection.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer zip.Close()

	awsSession := session.Must(session.NewSession())
	uploader := s3manager.NewUploader(awsSession)

	bucketName := "s3.us-west-2.amazonaws.com/bucketname"
	keyName := fmt.Sprintf("%s.forensicstore", createdCollection.Name)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: &bucketName,
		Key:    &keyName,
		Body:   zip,
	})
	if err != nil {
		log.Println(err)
	}
}
