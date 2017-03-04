<a id="top" name="top"></a>

# AWS CodePipeline Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows to write AWS Lambda functions and add them as action in your pipelines to customize the way they 
work.
  
[<img src="/_asset/misc_arrow-up.png" align="right">](#top)
## Quick Hands-On

> For step by step instructions on how to author your AWS Lambda function code in Go, see 
  [eawsy/aws-lambda-go-shim][eawsy-runtime].
  
```sh
go get -u -d github.com/eawsy/aws-lambda-go-event/...
```

```go
package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codepipeline"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/codepipelineevt"
)

func Handle(evt *codepipelineevt.Event, ctx *runtime.Context) (interface{}, error) {
	log.Println(evt)
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	svc := codepipeline.New(sess)
	_, err = svc.PutJobSuccessResult(&codepipeline.PutJobSuccessResultInput{
		JobId: aws.String(evt.Job.ID),
	})
	return nil, err
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/codepipelineevt

[aws-doc]: http://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-3F51B5.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-FF9800.svg?style=flat-square

