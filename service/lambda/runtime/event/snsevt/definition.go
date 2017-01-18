//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
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

package snsevt

import (
	"encoding/json"
	"time"
)

// MessageAttributes represents an SNS message attribute.
type MessageAttributes struct {
	// The attribute type.
	Type string

	// The attribute value.
	Value string
}

// Record represents the unit of data of an Amazon SNS message.
// See also http://docs.aws.amazon.com/sns/latest/dg/json-formats.html
type Record struct {
	// A Universally Unique Identifier, unique for each message published. For a
	// message that Amazon SNS resends during a retry, the message ID of the
	// original message is used.
	MessageID string

	// The type of the message.
	Type string

	// The time when the notification was published.
	Timestamp time.Time

	// The Subject parameter specified when the notification was published to the
	// topic.
	// Note that this is an optional parameter.
	Subject string

	// The Message value specified when the notification was published to the
	// topic.
	Message string

	// The attributes associated with the message.
	MessageAttributes map[string]*MessageAttributes

	// The Version of the Amazon SNS signature used.
	SignatureVersion string

	// Base64-encoded "SHA1withRSA" signature of the Message, MessageID, Subject
	// (if present), Type, Timestamp, and TopicARN values.
	// See also http://docs.aws.amazon.com/sns/latest/dg/SendMessageToHttp.verify.signature.html
	Signature string

	// The URL to the certificate that was used to sign the message.
	SignatureCertURL string

	// The ARN of the topic that this message was published to.
	TopicARN string

	// An URL that you can use to unsubscribe the endpoint from this topic. If you
	// visit this URL, Amazon SNS unsubscribes the endpoint and stops sending
	// notifications to this endpoint.
	UnsubscribeURL string
}

// EventRecord provides contextual information about an Amazon SNS event.
type EventRecord struct {
	// The event version.
	EventVersion string

	// The event source.
	EventSource string

	// The ARN of the Lambda subscription.
	EventSubscriptionARN string

	// The underlying Amazon SNS record associated with the event.
	SNS *Record
}

// String returns the string representation.
func (e *EventRecord) String() string {
	s, _ := json.MarshalIndent(e, "", "  ")
	return string(s)
}

// GoString returns the string representation.
func (e *EventRecord) GoString() string {
	return e.String()
}

// Event represents an Amazon SNS event.
type Event struct {
	// The list of Amazon SNS event records.
	Records []*EventRecord
}

// String returns the string representation.
func (e *Event) String() string {
	s, _ := json.MarshalIndent(e, "", "  ")
	return string(s)
}

// GoString returns the string representation.
func (e *Event) GoString() string {
	return e.String()
}
