<a id="top" name="top"></a>

# Amazon Alexa Skills Kit Events

[<img src="/_asset/misc_home.png" alt="Back to Home" align="right">](/)
[![Go Doc][badge-doc-go]][eawsy-doc]
[![AWS Doc][badge-doc-aws]][aws-doc]

This package allows you to write AWS Lambda functions to process Amazon Alexa Skill Kit Events.

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

	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/alexaskillsevt"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt *alexaskillsevt.Event, ctx *runtime.Context) (interface{}, error) {
	// See details about request types at https://developer.amazon.com/docs/custom-skills/request-types-reference.html
	if evt.Request.Type == "IntentRequest" {
		// sent when the user invokes your skill with a specific intent
		switch evt.Request.Intent.Name {
			case "Intent1":
				log.Println("Handle Intent1")
				// respond with PlainText speech and will keep session open for another request
				resp := alexaskillsevt.NewResponse().WithSpeechPlainText("Im responding to intent 1").WithCardSimple("My title here", "Card body will contain this text for intent 1 response").WithRepromptPlainText("Anything else I can help you with?").Get()
				return resp, nil
			case "Intent2":
				log.Println("Handle Intent2")
				// respond with SSML speech and will end the session
				resp := alexaskillsevt.NewResponse().WithSpeechSSML(`<p>Im <emphasis level="strong">responding</emphasis> to intent 2.</p><p>This will be said after a short pause</p>`).WithCardSimple("My title here", "Card body will contain this text for intent 2 response").WithEndSession().Get()
				return resp, nil
			default:
				log.Println("Unknown intent")
				return nil, errors.New("Unknown intent")
		}
	} else if evt.Request.Type == "LaunchRequest" {
		// sent when the user invokes your skill without providing a specific intent
		resp := alexaskillsevt.NewResponse().WithSpeechPlainText("Welcome to my Alexa Skill").WithCardSimple("Welcome", "Hello there").Get()
		return resp, nil
	}

	return nil, nil
}
```

[eawsy-runtime]: https://github.com/eawsy/aws-lambda-go-shim
[eawsy-doc]: https://godoc.org/github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/alexaskillsevt

[Amazon Alexa Developer]: https://developer.amazon.com/alexa
[Speech Synthesis Markup Language (SSML)]: https://developer.amazon.com/docs/custom-skills/speech-synthesis-markup-language-ssml-reference.html