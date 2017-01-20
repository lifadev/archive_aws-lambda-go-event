<a id="top" name="top"></a>

# Amazon Cognito Sync Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions to process Amazon Cognito 
Sync events published each time a dataset is synchronized.

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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cognitosyncevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *cognitosyncevt.Event, ctx *runtime.Context) (interface{}, error) {
	fmt.Println(evt)
	return nil, nil
}
```

[<img src="/_asset/misc_arrow-up.png" align="right">](#top)
## Supported Events

  - [Amazon API Gateway Proxy Events][eawsy-apigatewayproxyevt]
  - [Amazon CloudWatch Logs Events][eawsy-cloudwatchlogsevt]
  - [Amazon CloudWatch Scheduled Events][eawsy-cloudwatchschedevt]
  - [Amazon Cognito Sync Events][eawsy-cognitosyncevt]
  - [Amazon DynamoDB Streams Events][eawsy-dynamodbstreamsevt]
  - [Amazon Kinesis Streams Events][eawsy-kinesisstreamsevt]
  - [Amazon S3 Events][eawsy-s3evt]
  - [Amazon Simple Email Service Events][eawsy-sesevt]
  - [Amazon Simple Notification Service Events][eawsy-snsevt]
  - [AWS CloudFormation Events][eawsy-cloudformationevt]

[<img src="/_asset/misc_arrow-up.png" align="right">](#top)
## License

This product is licensed to you under the Apache License, Version 2.0 (the "License"); you may not use this product 
except in compliance with the License. See [LICENSE](/LICENSE) and [NOTICE](/NOTICE) for more information.

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/cognitosyncevt

[aws-doc]: http://docs.aws.amazon.com/cognito/latest/developerguide/what-is-amazon-cognito.html

[badge-doc-go]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[badge-doc-aws]: http://img.shields.io/badge/api-awsdoc-efaf27.svg?style=flat-square

[eawsy-apigatewayproxyevt]: /service/lambda/runtime/event/apigatewayproxyevt
[eawsy-cloudwatchlogsevt]: /service/lambda/runtime/event/cloudwatchlogsevt
[eawsy-cloudwatchschedevt]: /service/lambda/runtime/event/cloudwatchschedevt 
[eawsy-cognitosyncevt]: /service/lambda/runtime/event/cognitosyncevt
[eawsy-dynamodbstreamsevt]: /service/lambda/runtime/event/dynamodbstreamsevt
[eawsy-kinesisstreamsevt]: /service/lambda/runtime/event/kinesisstreamsevt
[eawsy-s3evt]: /service/lambda/runtime/event/s3evt
[eawsy-sesevt]: /service/lambda/runtime/event/sesevt
[eawsy-snsevt]: /service/lambda/runtime/event/snsevt
[eawsy-cloudformationevt]: /service/lambda/runtime/event/cloudformationevt
