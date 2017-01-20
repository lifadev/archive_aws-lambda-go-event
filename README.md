<a id="top" name="top"></a>
[<img src="_asset/logo_powered-by-aws.png" alt="Powered by Amazon Web Services" align="right">][aws-home]
[<img src="_asset/logo_created-by-eawsy.png" alt="Created by eawsy" align="right">][eawsy-home]

# eawsy/aws-lambda-go-event

> Type definitions and helpers for AWS Lambda event sources.

[![Api][badge-api]][eawsy-godoc]
[![Status][badge-status]](#top)
[![License][badge-license]](LICENSE)
[![Chat][badge-chat]][eawsy-chat]

[AWS Lambda][aws-lambda-home] lets you run code without provisioning or managing servers. With 
[eawsy/aws-lambda-go-shim][eawsy-runtime], you can author your Lambda function code in Go. This project provides type 
definitions and helpers to deal with [AWS Lambda event source mapping][aws-lambda-mapping].

[<img src="_asset/misc_arrow-up.png" align="right">](#top)
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

[<img src="_asset/misc_arrow-up.png" align="right">](#top)
## About

[![eawsy](_asset/logo_eawsy.png)][eawsy-home]

This project is maintained and funded by Alsanium, SAS.

[We][eawsy-home] :heart: [AWS][aws-home] and open source software. See [our other projects][eawsy-github], or 
[hire us][eawsy-hire] to help you build modern applications on AWS.

[<img src="_asset/misc_arrow-up.png" align="right">](#top)
## Contact

We want to make it easy for you, users and contributers, to talk with us and connect with each others, to share ideas, 
solve problems and make help this project awesome. Here are the main channels we're running currently and we'd love to 
hear from you on them.

### Twitter 
  
[eawsyhq][eawsy-twitter] 

Follow and chat with us on Twitter. 

Share stories!

### Gitter 

[eawsy/bavardage][eawsy-chat]

This is for all of you. Users, developers and curious. You can find help, links, questions and answers from all the 
community including the core team.

Ask questions!

### GitHub

[pull requests][eawsy-pr] & [issues][eawsy-issues]

You are invited to contribute new features, fixes, or updates, large or small; we are always thrilled to receive pull 
requests, and do our best to process them as fast as we can.

Before you start to code, we recommend discussing your plans through the [eawsy/bavardage channel][eawsy-chat], 
especially for more ambitious contributions. This gives other contributors a chance to point you in the right direction, 
give you feedback on your design, and help you find out if someone else is working on the same thing.

Write code!

[<img src="_asset/misc_arrow-up.png" align="right">](#top)
## License

This product is licensed to you under the Apache License, Version 2.0 (the "License"); you may not use this product 
except in compliance with the License. See [LICENSE](LICENSE) and [NOTICE](NOTICE) for more information.

[<img src="_asset/misc_arrow-up.png" align="right">](#top)
## Trademark

Alsanium, eawsy, the "Created by eawsy" logo, and the "eawsy" logo are trademarks of Alsanium, SAS. or its affiliates in 
France and/or other countries.

Amazon Web Services, the "Powered by Amazon Web Services" logo, and AWS Lambda are trademarks of Amazon.com, Inc. or its 
affiliates in the United States and/or other countries.


[eawsy-home]: https://eawsy.com
[eawsy-github]: https://github.com/eawsy
[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-chat]: https://gitter.im/eawsy/bavardage
[eawsy-twitter]: https://twitter.com/@eawsyhq
[eawsy-godoc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event
[eawsy-hire]: https://docs.google.com/forms/d/e/1FAIpQLSfPvn1Dgp95DXfvr3ClPHCNF5abi4D1grveT5btVyBHUk0nXw/viewform
[eawsy-pr]: https://github.com/eawsy/aws-lambda-go-event/issues?q=is:pr%20is:open
[eawsy-issues]: https://github.com/eawsy/aws-lambda-go-event/issues?q=is:issue%20is:open

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

[aws-home]: https://aws.amazon.com/
[aws-lambda-home]: https://aws.amazon.com/lambda/
[aws-lambda-events]: http://docs.aws.amazon.com/lambda/latest/dg/invoking-lambda-function.html
[aws-lambda-mapping]: http://docs.aws.amazon.com/lambda/latest/dg/intro-invocation-modes.html

[badge-api]: http://img.shields.io/badge/api-godoc-7986cb.svg?style=flat-square
[badge-chat]: http://img.shields.io/badge/chat-gitter-e91e63.svg?style=flat-square
[badge-status]: http://img.shields.io/badge/status-stable-689f38.svg?style=flat-square
[badge-license]: http://img.shields.io/badge/license-apache-757575.svg?style=flat-square
