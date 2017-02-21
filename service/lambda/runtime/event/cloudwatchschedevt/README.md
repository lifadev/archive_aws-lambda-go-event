<a id="top" name="top"></a>

# Amazon CloudWatch Scheduled Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions executed on a regular,
scheduled basis using the schedule event capability in Amazon CloudWatch Events.
  
[<img src="/_asset/misc_arrow-up.png" align="right">](#top)
## Quick Hands-On

> For step by step instructions on how to author your AWS Lambda function code in Go, see 
  [eawsy/aws-lambda-go-shim][eawsy-runtime].
  
```sh
go get -u -d github.com/eawsy/aws-lambda-go-event/...
```

```go
package main

// /* Required, but no C code needed. */
import "C"

import (
	"fmt"

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudwatchschedevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *cloudwatchschedevt.Event, ctx *runtime.Context) (interface{}, error) {
	fmt.Println(evt)
	return nil, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cloudwatchschedevt

[aws-doc]: http://docs.aws.amazon.com/lambda/latest/dg/with-scheduled-events.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-efaf27.svg?style=flat-square
