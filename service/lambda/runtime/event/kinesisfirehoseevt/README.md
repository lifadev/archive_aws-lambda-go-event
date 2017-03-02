<a id="top" name="top"></a>

# Amazon Kinesis Firehose Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions to transform incoming Amazon Kinesis Firehose source data and
deliver the transformed data to destinations.
  
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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisfirehoseevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(in *kinesisfirehoseevt.Input, ctx *runtime.Context) (kinesisfirehoseevt.Output, error) {
	rcds := make([]*kinesisfirehoseevt.OutputRecord, 0)
	for _, r := range in.Records {
		log.Println(r)
		rcds = append(rcds, &kinesisfirehoseevt.OutputRecord{
			RecordID: r.RecordID,
			Result:   "Ok",
			Data:     r.Data,
		})
	}
	return kinesisfirehoseevt.Output{Records: rcds}, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisfirehoseevt

[aws-doc]: http://docs.aws.amazon.com/firehose/latest/dev/what-is-this-service.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-3F51B5.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-FF9800.svg?style=flat-square

