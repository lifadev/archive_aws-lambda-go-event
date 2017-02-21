<a id="top" name="top"></a>

# Amazon Simple Notification Service Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions to process Amazon Simple
Notification Service notifications.

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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/snsevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *snsevt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, rec := range evt.Records {
		fmt.Println(rec)
	}
	return nil, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/snsevt

[aws-doc]: http://docs.aws.amazon.com/sns/latest/dg/welcome.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-efaf27.svg?style=flat-square
