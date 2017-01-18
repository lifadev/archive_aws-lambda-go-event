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

package cloudwatchlogsevt

import (
	"encoding/json"
	"time"
)

// LogEvent represents a log event, which is a record of activity that was
// recorded by the application or resource being monitored.
type LogEvent struct {
	// A unique identifier for every log event.
	ID string

	// The time the event occurred.
	Timestamp time.Time `json:"-"`

	// The data contained in the log event.
	Message string
}

// EventRecord provides contextual information about an Amazon CloudWatch Logs
// event.
type EventRecord struct {
	// The AWS Account ID of the originating log data.
	Owner string

	// The log group name of the originating log data.
	LogGroup string

	// The log stream name of the originating log data.
	LogStream string

	// The actual log data.
	LogEvent *LogEvent

	// Message type can be one of:
	// - DATA_MESSAGE: Usual message type for payload.
	// - CONTROL_MESSAGE: Excpetional message type for ping.
	MessageType string

	// The list of subscription filter names that matched with the originating log
	// data.
	SubscriptionFilters []string
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

// EventRecords represents a list of EventRecord.
type EventRecords []*EventRecord

// Event represents an Amazon CloudWatch Logs event.
type Event struct {
	// The list of Amazon CloudWatch Logs event records.
	Records EventRecords `json:"awslogs"`
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
