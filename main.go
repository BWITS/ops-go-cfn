package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/MYOB-Technology/ops-go-cfn/actions"
	"github.com/MYOB-Technology/ops-go-cfn/parse"
	"github.com/MYOB-Technology/ops-go-cfn/stack"
)

var (
	action   = flag.String("a", "provision", "create or delete")
	name     = flag.String("n", "", "Stack name.")
	template = flag.String("t", "templates/test-template.yml", "Template file path.")
	params   = flag.String("p", "templates/test-params.json", "Parameters file path.")
	bucket   = flag.String("b", "", "Bucket containing template.")
	timeout  = flag.Int64("x", 5, "Timeout in minutes.")
)

func main() {
	flag.Parse()

	var (
		svc = stack.Service{
			Context: context.Background(),
			CFAPI:   cf.New(session.Must(session.NewSession())),
			S3API:   s3.New(session.Must(session.NewSession())),
		}
	)

	switch *action {
	case "provision":
		cfg := stack.Config{
			StackName:  *name,
			Parameters: parse.Params(readFile(*params)),
			Template:   readFile(*template),
			BucketName: *bucket,
			Timeout:    *timeout,
		}
		status, err := actions.Provision(&svc, &cfg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Stack - %s\n", status)
	case "delete":
		status, err := actions.Delete(&svc, *name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Stack - %s\n", status)
	default:
		fmt.Printf("Unknown action '%s'\n", *action)
	}
}

func readFile(f string) []byte {
	content, err := ioutil.ReadFile(f)
	if err != nil {
		log.Fatal(err)
	}
	return content
}
