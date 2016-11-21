<a id="top" name="top"></a>
[<img src="_asset/powered-by-aws.png" alt="Powered by Amazon Web Services" align="right">][aws-home]
[<img src="_asset/created-by-eawsy.png" alt="Created by eawsy" align="right">][eawsy-home]

# eawsy/aws-lambda-go-event
> A collection of AWS services event definitions for AWS Lambda Go runtime.

[![Runtime][runtime-badge]][eawsy-runtime]
[![Api][api-badge]][eawsy-godoc]
[![Chat][chat-badge]][eawsy-gitter]
![Status][status-badge]
[![License][license-badge]](LICENSE)
<sup>•</sup> <sup>•</sup> <sup>•</sup>
[![Hire us][hire-badge]][eawsy-hire-form]

[AWS Lambda][aws-lambda-home] lets you run code without provisioning or managing servers. 
You can configure [some AWS services][aws-lambda-events] as event sources for AWS Lambda functions. After you
preconfigure the event source mapping, your Lambda function gets invoked automatically when these event sources detect
events.

This project provides a collection of AWS services event definitions and utility functions to streamline 
[AWS Lambda event source mapping][aws-lambda-mapping].

Currently supported AWS services:
  - [Amazon S3](#amazon-s3)
  - [Amazon Kinesis Streams](#amazon-kinesis-streams)
  - [Amazon Simple Notification Service](#amazon-simple-notification-service)

## Preview

> For step by step instructions on how to build and package a Lambda function in Go, see 
  [Deployment Package][eawsy-packaging] on [eawsy/aws-lambda-go][eawsy-runtime].

[<img src="_asset/arrow-up.png" align="right">](#top)
### Amazon S3

[![AWS Doc][aws-badge]][aws-s3-dev]
[![Go Doc][api-badge]][eawsy-s3-dev]

```sh
go get -u -d github.com/eawsy/aws-lambda-go-event/...
```

```go
package main

import (
	"log"

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/s3evt"
	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt *s3evt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, rec := range evt.Records {
		log.Printf(
			"Event: %s, Bucket: %s, Object: %s\n",
			rec.EventName, rec.S3.Bucket.Name, rec.S3.Object.Key,
		)
	}
	return nil, nil
}

func init() {
	runtime.Handle(s3evt.HandlerFunc(handle))
}

func main() {}
```

```sh
docker run --rm -v $GOPATH:/go -v $PWD:/tmp eawsy/aws-lambda-go
```

[<img src="_asset/arrow-up.png" align="right">](#top)
### Amazon Kinesis Streams

[![AWS Doc][aws-badge]][aws-kinesis-dev]
[![Go Doc][api-badge]][eawsy-kinesis-dev]

```sh
go get -u -d github.com/eawsy/aws-lambda-go-event/...
```

```go
package main

import (
	"log"

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisevt"
	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt *kinesisevt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, rec := range evt.Records {
		log.Printf(
			"PartitionKey: %s, SequenceNumber: %s, Size: %dB\n",
			rec.Kinesis.PartitionKey, rec.Kinesis.SequenceNumber, len(rec.Kinesis.Data),
		)
	}
	return nil, nil
}

func init() {
	runtime.Handle(kinesisevt.HandlerFunc(handle))
}

func main() {}
```

```sh
docker run --rm -v $GOPATH:/go -v $PWD:/tmp eawsy/aws-lambda-go
```

[<img src="_asset/arrow-up.png" align="right">](#top)
### Amazon Simple Notification Service

[![AWS Doc][aws-badge]][aws-sns-dev]
[![Go Doc][api-badge]][eawsy-sns-dev]

```sh
go get -u -d github.com/eawsy/aws-lambda-go-event/...
```

```go
package main

import (
	"log"

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/snsevt"
	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

func handle(evt *snsevt.Event, ctx *runtime.Context) (interface{}, error) {
	for _, rec := range evt.Records {
		log.Printf(
			"Subject: %s, Message: %s\n",
			rec.SNS.Subject, rec.SNS.Message,
		)
	}
	return nil, nil
}

func init() {
	runtime.Handle(snsevt.HandlerFunc(handle))
}

func main() {}
```

```sh
docker run --rm -v $GOPATH:/go -v $PWD:/tmp eawsy/aws-lambda-go
```

## About

[![eawsy](_asset/eawsy-logo.png)][eawsy-home]

This project is maintained and funded by Alsanium, SAS.

[We][eawsy-home] :heart: [AWS][aws-home] and open source software. See [our other projects][eawsy-github], or 
[hire us][eawsy-hire-form] to help you build modern applications on AWS.

## License

This product is licensed to you under the Apache License, Version 2.0 (the "License"); you may not use this product 
except in compliance with the License. See [LICENSE](LICENSE) and [NOTICE](NOTICE) for more information.

## Trademark

Alsanium, eawsy, the "Created by eawsy" logo, and the "eawsy" logo are trademarks of Alsanium, SAS. or its affiliates 
in France and/or other countries.

Amazon Web Services, the "Powered by Amazon Web Services" logo, and AWS Lambda are trademarks of Amazon.com, Inc. or 
its affiliates in the United States and/or other countries.

[eawsy-home]: https://eawsy.com
[eawsy-github]: https://github.com/eawsy
[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go
[eawsy-gitter]: https://gitter.im/eawsy/bavardage
[eawsy-godoc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event
[eawsy-wiki]: https://github.com/eawsy/aws-lambda-go-net/wiki
[eawsy-hire-form]: https://docs.google.com/forms/d/e/1FAIpQLSfPvn1Dgp95DXfvr3ClPHCNF5abi4D1grveT5btVyBHUk0nXw/viewform
[eawsy-packaging]: https://github.com/eawsy/aws-lambda-go/wiki/Deployment%20Package
[eawsy-s3-dev]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/s3evt
[eawsy-kinesis-dev]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/kinesisevt
[eawsy-sns-dev]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/snsevt
[aws-home]: https://aws.amazon.com/
[aws-lambda-home]: https://aws.amazon.com/lambda/
[aws-lambda-events]: http://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-function.html
[aws-lambda-mapping]: http://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html
[aws-s3-dev]: http://docs.aws.amazon.com/AmazonS3/latest/dev/Welcome.html
[aws-kinesis-dev]: http://docs.aws.amazon.com/streams/latest/dev/introduction.html
[aws-sns-dev]: http://docs.aws.amazon.com/sns/latest/dg/welcome.html
[runtime-badge]: http://img.shields.io/badge/runtime-go-ef6c00.svg?style=flat-square
[api-badge]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[aws-badge]: http://img.shields.io/badge/api-awsdoc-efaf27.svg?style=flat-square
[chat-badge]: http://img.shields.io/badge/chat-gitter-e91e63.svg?style=flat-square
[status-badge]: http://img.shields.io/badge/status-stable-689f38.svg?style=flat-square
[license-badge]: http://img.shields.io/badge/license-apache-757575.svg?style=flat-square
[hire-badge]: http://img.shields.io/badge/hire-eawsy-2196f3.svg?style=flat-square
