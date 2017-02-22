//
// Copyright 2017 Alsanium, SAS. or its affiliates. All rights reserved.
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

package kinesisfirehoseevt

import (
	"encoding/json"
	"time"
)

// OutputRecord represents the transformed Amazon Kinesis Firehose record
type OutputRecord struct {
	// The record ID is passed from Amazon Kinesis Firehose to AWS Lambda
	// during the invocation. The transformed record must contain the same
	// record ID. Any mismatch between the ID of the original record and the
	// ID of the transformed record is treated as a data transformation
	// failure.
	RecordID string `json:"recordId"`

	// The status of the data transformation of the record. The possible
	// values are:
	// - "Ok": the record was transformed successfully.
	// - "Dropped": the record was dropped intentionally by your processing
	//   logic.
	// - "ProcessingFailed": the record could not be transformed.
	// If a record has a status of "Ok" or "Dropped",
	// Amazon Kinesis firehose considers it successfully processed.
	// Otherwise, Amazon Kinesis Firehose considers it unsuccessfully
	// processed.
	Result string `json:"result"`

	// The transformed data payload, after base64-encoding.
	Data []byte `json:"data"`
}

// Output represents the result of the processing of Amazon Kinesis Firehose
// input records.
type Output struct {
	// Transformed Amazon Kinesis Firehose records.
	Records []*OutputRecord `json:"records"`
}

// InputRecord represents the unit of data of an Amazon Kinesis Firehose event.
type InputRecord struct {
	// The unique identifier of the record passed from
	// Amazon Kinesis Firehose to AWS Lambda.
	RecordID string

	// The approximate time that the record was inserted into the
	// Amazon Kinesis Firehose delivery stream. This is set when a delivery
	// stream successfully receives and stores a record and is commonly
	// referred to as a server-side timestamp. It has millisecond precision
	// and there are no guarantees about its accuracy, or that it is always
	// increasing. For example, records in a specific
	// Amazon Kinesis Firehose delivery stream might have timestamps that
	// are out of order.
	ApproximateArrivalTimestamp time.Time

	// The data blob. The data in the blob is both opaque and immutable to
	// the Amazon Kinesis Firehose service, which does not inspect,
	// interpret, or change the data in the blob in any way. The data blob
	// consists of any kind of data and the total size must not exceed the
	// maximum record size (1 MB).
	// Data is automatically base64 encoded/decoded by the SDK.
	Data []byte
}

// Input represents an Amazon Kinesis Firehose delivery stream event and
// provides contextual information.
type Input struct {
	// The invocation ID.
	InvocationID string

	// The ARN of the Amazon Kinesis Firehose.
	DeliveryStreamARN string

	// The AWS region where the event originated.
	Region string

	// The list of Amazon Kinesis Firehose event records.
	Records []*InputRecord
}

// String returns the string representation.
func (e *Input) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}

// GoString returns the string representation.
func (e *Input) GoString() string {
	return e.String()
}

// String returns the string representation.
func (e *Output) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}

// GoString returns the string representation.
func (e *Output) GoString() string {
	return e.String()
}
