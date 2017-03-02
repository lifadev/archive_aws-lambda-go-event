<a id="top" name="top"></a>

# AWS CloudFormation Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you write AWS Lambda functions as the back end of your API 
through the AWS Lambda proxy integration.

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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudformationevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *cloudformationevt.Event, ctx *runtime.Context) (interface{}, error) {
	log.Println(evt)
	return nil, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudformationevt

[aws-doc]: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-3F51B5.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-FF9800.svg?style=flat-square
