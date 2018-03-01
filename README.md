# eas3y
[![Build Status](https://travis-ci.org/beeceej/eas3y.svg?branch=master)](https://travis-ci.org/beeceej/eas3y)


Example usage:

```go
package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/beeceej/eas3y"
)

var svc = s3.New(session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")})))

type ExampleCustom struct {
	Name string
}

type ExampleDefault struct {
	Name string
}

// SaveConfig builds a save configuration for eas3y to use when putting to s3
func (e *ExampleCustom) SaveConfig() *eas3y.Config {
	return eas3y.NewConfig(
		eas3y.WithBucket("eas3y"),
		eas3y.WithKey(e.Name),
		eas3y.WithContentType("text/xml"),
		eas3y.AsXML(),
	)
}

// SaveConfig builds a save configuration for eas3y to use when putting to s3
func (e *ExampleDefault) SaveConfig() *eas3y.Config {
	return eas3y.NewDefaultConfig("eas3y", e.Name)
}

func main() {
	var (
		out *s3.PutObjectOutput
		err error
	)

	updateReadme := parseFlags()

	eDefault := &ExampleDefault{Name: "Default"}
	out, err = eas3y.Save(svc, eDefault)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(out.GoString())
	}

	eCustom := &ExampleCustom{Name: "Custom"}
	out, err = eas3y.Save(svc, eCustom)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(out.GoString())
	}
}

func parseFlags() (shouldUpdate bool) {
	b := flag.Bool("update-readme", false, " try --update-readme=true if you'd like to update the documentation")
	flag.Parse()
	return *b
}

```

Will result in the following files in S3

-> `[bucket=eas3y]/Custom.xml`

-> `[bucket=eas3y]/Default.json`
