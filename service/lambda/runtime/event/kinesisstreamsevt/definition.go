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

package kinesisstreamsevt

import (
	"encoding/json"
	"time"
)

// Record represents the unit of data of an Amazon Kinesis stream.
type Record struct {
	// The schema version for the record.
	KinesisSchemaVersion string

	// The unique identifier of the record in the stream.
	// See also http://docs.aws.amazon.com/streams/latest/dev/kinesis-record-processor-duplicates.html
	SequenceNumber string

	// The approximate time that the record was inserted into the stream.
	// This is set when a stream successfully receives and stores a record and
	// is commonly referred to as a server-side timestamp.
	// It has millisecond precision and there are no guarantees about its
	// accuracy, or that the it is always increasing. For example, records in
	// a shard or across a stream might have timestamps that are out of order.
	ApproximateArrivalTimestamp time.Time

	// Identifies which shard in the stream the data record is assigned to.
	PartitionKey string

	// The data blob. The data in the blob is both opaque and immutable to the
	// Amazon Kinesis service, which does not inspect, interpret, or change
	// the data in the blob in any way. When the data blob (the payload before
	// base64-encoding) is added to the partition key size, the total size
	// must not exceed the maximum record size (1 MB).
	// Data is automatically base64 encoded/decoded by the SDK.
	Data []byte
}

// EventRecord provides contextual information about an Amazon Kinesis streams
// event.
type EventRecord struct {
	// The event id.
	EventID string

	// The name of the event.
	EventName string

	// The event version.
	EventVersion string

	// The source of the event.
	EventSource string

	// The ARN of the event source.
	EventSourceARN string

	// The ARN for the identity used to invoke the Lambda Function.
	InvokeIdentityARN string

	// The AWS region where the event originated.
	AWSRegion string

	// The underlying Amazon Kinesis streams record associated with the event.
	Kinesis *Record
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

// Event represents an Amazon Kinesis streams event.
type Event struct {
	// The list of Amazon Kinesis streams event records.
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
