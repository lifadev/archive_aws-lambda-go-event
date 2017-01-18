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

package cloudwatchschedevt

import (
	"encoding/json"
	"time"
)

// Event represents an Amazon CloudWatch Scheduled event.
type Event struct {
	// A unique value generated for every event.
	ID string

	// The event timestamp, which is the time when the event have been actually
	// triggered.
	Time time.Time

	// The 12-digit number identifying an AWS account.
	Account string

	// Identifies the AWS region where the event originated.
	Region string

	// A JSON object, whose content is at the discretion of the service
	// originating the event.
	// In our case the value is always an empty object.
	Detail json.RawMessage

	// Identifies, in combination with the source field, the fields and values
	// that will appear in the detail field.
	// In our case the value is always "Scheduled Event".
	DetailType string `json:"detail-type"`

	// Identifies the service that sourced the event. All events sourced from
	// within AWS will begin with "aws.". Customer-generated events can have any
	// value here as long as it doesn't begin with "aws.".
	// In our case the value is always "aws.events".
	Source string

	// This JSON array contains ARNs that identify resources that are involved in
	// the event.
	// In our case the value is always an Amazon CloudWatch Rule ARN.
	Resources []string
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
