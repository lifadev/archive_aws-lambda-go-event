<a id="top" name="top"></a>

# Amazon DynamoDB Streams Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions to perform custom actions 
in response to updates made to DynamoDB tables.

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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/dynamodbstreamsevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *dynamodbstreamsevt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, rec := range evt.Records {
		log.Println(rec)
	}
	return nil, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/dynamodbstreamsevt

[aws-doc]: http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-3F51B5.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-FF9800.svg?style=flat-square
