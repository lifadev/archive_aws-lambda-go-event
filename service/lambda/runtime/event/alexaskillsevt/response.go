//
// Copyright 2018 Kyle Vanek. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package alexaskillsevt

var (
	// Version specifies the response version number. Defaults to "1.0" if not defined
	Version string
)

// NewResponse creates a new response object
func NewResponse() *CreateResponse {
	if Version == "" {
		Version = "1.0"
	}
	return &CreateResponse{
		resp: &EventResponse{Version: Version},
	}
}

// CreateResponse object type
type CreateResponse struct {
	resp *EventResponse
}

// Get returns the actual EventResponse for Lambda to return
func (r *CreateResponse) Get() *EventResponse {
	if r.resp.Version == "" {
		r.resp.Version = "1.0"
	}
	return r.resp
}

// WithEndSession specifies that the session will end after this response
func (r *CreateResponse) WithEndSession() *CreateResponse {
	r.resp.Response.ShouldEndSession = true
	return r
}

// WithSpeechPlainText sets the speechlet response to the specified PlainText text
func (r *CreateResponse) WithSpeechPlainText(text string) *CreateResponse {
	r.resp.Response.OutputSpeech = &ResponseOutputSpeech{
		Type: "PlainText",
		Text: text,
	}
	return r
}

// WithSpeechSSML sets the speechlet response to the specified SSML text
func (r *CreateResponse) WithSpeechSSML(text string) *CreateResponse {
	r.resp.Response.OutputSpeech = &ResponseOutputSpeech{
		Type: "SSML",
		SSML: `<speak>` + text + `</speak>`,
	}
	return r
}

// WithCardSimple sets the card to a Simple card type with a title and text
func (r *CreateResponse) WithCardSimple(title string, text string) *CreateResponse {
	r.resp.Response.Card = &ResponseCard{
		Type:    "Simple",
		Title:   title,
		Content: text,
	}

	return r
}

// WithCardStandard sets the card to a Standard card type with a title, text, and either/both a small and large image using specified URLs
func (r *CreateResponse) WithCardStandard(title string, text string, imageSmallURL string, imageLargeURL string) *CreateResponse {
	r.resp.Response.Card = &ResponseCard{
		Type:  "Standard",
		Title: title,
		Text:  text,
		Image: &ResponseCardImage{
			SmallImageURL: imageSmallURL,
			LargeImageURL: imageLargeURL,
		},
	}

	return r
}

// WithCardLinkAccount sets the card to a LinkAccount card type
func (r *CreateResponse) WithCardLinkAccount() *CreateResponse {
	r.resp.Response.Card = &ResponseCard{
		Type: "LinkAccount",
	}

	return r
}

// WithRepromptPlainText sets the reprompt to the specified PlainText text in the response
func (r *CreateResponse) WithRepromptPlainText(text string) *CreateResponse {
	r.resp.Response.Reprompt = &ResponseReprompt{
		OutputSpeech: &ResponseOutputSpeech{
			Type: "PlainText",
			Text: text,
		},
	}

	return r
}

// WithRepromptSSML sets the reprompt to the specified SSML text in the response
func (r *CreateResponse) WithRepromptSSML(text string) *CreateResponse {
	r.resp.Response.Reprompt = &ResponseReprompt{
		OutputSpeech: &ResponseOutputSpeech{
			Type: "SSML",
			SSML: text,
		},
	}

	return r
}

// WithSessionAttribute adds a session attribute to the response
func (r *CreateResponse) WithSessionAttribute(key string, item interface{}) *CreateResponse {
	if r.resp.SessionAttributes == nil {
		r.resp.SessionAttributes = map[string]interface{}{}
	}
	r.resp.SessionAttributes[key] = item

	return r
}

// WithDirective adds a directive to the response
func (r *CreateResponse) WithDirective(directive map[string]interface{}) *CreateResponse {
	if r.resp.Response.Directives == nil {
		r.resp.Response.Directives = make([]interface{}, 1)
	}
	r.resp.Response.Directives = append(r.resp.Response.Directives, directive)

	return r
}

// EventResponse provides information about an Amazon Alexa Skills event response
type EventResponse struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes"`
	Response          ResponseValue          `json:"response"`
}

// ResponseValue provides specific response information in a EventResponse
type ResponseValue struct {
	OutputSpeech     *ResponseOutputSpeech `json:"outputSpeech,omitempty"`
	Card             *ResponseCard         `json:"card,omitempty"`
	Reprompt         *ResponseReprompt     `json:"reprompt,omitempty"`
	Directives       []interface{}         `json:"directives,omitempty"` // See directives for interface information at https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#card-object
	ShouldEndSession bool                  `json:"shouldEndSession"`
}

// ResponseOutputSpeech provides OutputSpeech information in an EventResponse
type ResponseOutputSpeech struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
}

// ResponseCard provides Card information in an EventResponse
type ResponseCard struct {
	Type    string             `json:"type,omitempty"`
	Title   string             `json:"title,omitempty"`
	Content string             `json:"content,omitempty"`
	Text    string             `json:"text,omitempty"`
	Image   *ResponseCardImage `json:"image,omitempty"`
}

// ResponseCardImage provides image URLs for cards
type ResponseCardImage struct {
	SmallImageURL string `json:"smallImageUrl,omitempty"`
	LargeImageURL string `json:"largeImageUrl,omitempty"`
}

// ResponseReprompt provides OutputSpeech information for a reprompt in an EventResponse
type ResponseReprompt struct {
	OutputSpeech *ResponseOutputSpeech `json:"outputSpeech,omitempty"`
}
