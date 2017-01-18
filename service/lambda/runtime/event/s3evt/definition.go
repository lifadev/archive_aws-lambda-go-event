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

package s3evt

import (
	"encoding/json"
	"time"
)

// Object provides information about the Amazon S3 object that caused the event.
type Object struct {
	// The object key provides information about the bucket and object involved in
	// the event.
	// Note that the object keyname value is URL encoded. For example
	// "red flower.jpg" becomes "red+flower.jpg".
	Key string

	// The object size in bytes. Provided for "ObjectCreated" event, otherwise 0.
	Size int

	// The object ETag. Provided for "ObjectCreated" event, otherwise empty.
	ETag string

	// The object version if bucket is versioning-enabled, otherwise empty.
	VersionID string

	// A string representation of a hexadecimal value used to determine event
	// sequence, only used with PUTs and DELETEs.
	// It provides a way to determine the sequence of events. Event notifications
	// are not guaranteed to arrive in the order that the events occurred.
	// However, notifications from events that create objects (PUTs) and delete
	// objects contain a sequencer, which can be used to determine the order of
	// events for a given object key.
	// If you compare the sequencer strings from two event notifications on the
	// same object key, the event notification with the greater sequencer
	// hexadecimal value is the event that occurred later. If you are using event
	// notifications to maintain a separate database or index of your Amazon S3
	// objects, you will probably want to compare and store the sequencer values
	// as you process each event notification.
	// Note that:
	// - The sequencers cannot be used to determine order for events on different
	//   object keys.
	// - The sequencers can be of different lengths. So to compare these values,
	//   you first right pad the shorter value with zeros and then do
	//   lexicographical comparison.
	Sequencer string
}

// UserIdentity represents details about an IAM identity,
type UserIdentity struct {
	// A unique identifier for the entity that made the call. For requests made
	// with temporary security credentials, this value includes the session name
	// that is passed to the AssumeRole, AssumeRoleWithWebIdentity, or
	// GetFederationToken API call.
	PrincipalID string
}

// Bucket provides information about the Amazon S3 bucket from which the event
// originated.
type Bucket struct {
	// The bucket name.
	Name string

	// The bucket ARN.
	ARN string

	// The IAM identity of the Amazon S3 bucket owner.
	OwnerIdentity *UserIdentity
}

// Record represents the unit of data of an Amazon S3 notification.
type Record struct {
	// The schema version for the record.
	S3SchemaVersion string

	// The Amazon S3 notification feature enables you to receive notifications
	// when certain events happen in your bucket. To enable notifications, you
	// must first add a notification configuration identifying the events you want
	// Amazon S3 to publish, and the destinations where you want Amazon S3 to send
	// the event notifications. You store this configuration in the notification
	// subresource associated with a bucket. ConfigurationID represents the ID of
	// this configuration.
	ConfigurationID string

	// Information about the Amazon S3 bucket from which the event originated.
	Bucket *Bucket

	// Information about the Amazon S3 object that caused the event.
	Object *Object
}

// RequestParameters provides information about the request that caused the
// event.
type RequestParameters struct {
	// The IP address of the requester.
	SourceIPAddress string
}

// ResponseElements provides tracing information returned by Amazon S3 in the
// response to the request that caused the event. It is useful if you want to
// trace the request by following up with Amazon S3 support.
type ResponseElements struct {
	// The Amazon S3 host that processed the request.
	AMZID2 string `json:"x-amz-id-2"`

	// The Amazon S3 generated request ID.
	AMZRequestID string `json:"x-amz-request-id"`
}

// EventRecord provides contextual information about an Amazon S3 event.
type EventRecord struct {
	// The Amazon S3 event name.
	// See also http://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html#notification-how-to-event-types-and-destinations
	EventName string

	// The time when Amazon S3 finished processing the request.
	EventTime time.Time

	// The event version.
	EventVersion string

	// The event Source.
	EventSource string

	// The IAM identity of the user who caused the event.
	UserIdentity *UserIdentity

	// Information about the request that caused the event.
	RequestParameters *RequestParameters

	// Tracing information returned by Amazon S3 in the response to the request
	// that caused the event.
	ResponseElements *ResponseElements

	// The AWS region where the event originated.
	AWSRegion string

	// The underlying Amazon S3 record associated with the event.
	S3 *Record
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

// Event represents an Amazon S3 event.
// See also http://docs.aws.amazon.com/AmazonS3/latest/dev/notification-content-structure.html
type Event struct {
	// The list of Amazon S3 event records.
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
