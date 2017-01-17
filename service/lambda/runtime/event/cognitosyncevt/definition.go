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

package cognitosyncevt

import "encoding/json"

// DatasetRecord contains information about each record in a data set.
type DatasetRecord struct {
	// Old value of the record.
	OldValue string

	// New value of the record.
	NewValue string

	// The operation associated with the record:
	// - replace: if a record is created or updated.
	// - remove: if a record is deleted.
	OP string
}

// Event represents an Amazon Cognito Sync event.
type Event struct {
	// The version of the event.
	Version int

	// The identity pool ID associated with the dataset.
	IdentityPoolID string

	// The actual identity ID from Amazon Cognito.
	IdentityID string

	// The region in which dataset resides.
	Region string

	// The dataset name of the event.
	DatasetName string

	// The map of dataset records for the event.
	DatasetRecords map[string]*DatasetRecord

	// The event type.
	// In our case the value is always "SyncTrigger".
	EventType string
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
