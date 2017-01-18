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
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"
)

type logEventAlias LogEvent

type timestamp struct {
	time.Time
}

// UnmarshalJSON interprets the data as an int64 being the number of
// milliseconds elapsed since January 1, 1970 00:00:00 UTC. It then sets *t to
// a copy of the interpreted time.
func (t *timestamp) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	sec := v / 1000
	nsec := (v - sec*1000) * int64(time.Millisecond)

	t.Time = time.Unix(sec, nsec)

	return nil
}

type jsonLogEvent struct {
	*logEventAlias
	Timestamp timestamp
}

// UnmarshalJSON interprets data as a LogEvent with a special timestamp. It then
// leverages type aliasing and struct embedding to fill LogEvent with an usual
// time.Time.
func (le *LogEvent) UnmarshalJSON(data []byte) error {
	var jle jsonLogEvent
	if err := json.Unmarshal(data, &jle); err != nil {
		return err
	}

	*le = *(*LogEvent)(jle.logEventAlias)
	le.Timestamp = jle.Timestamp.Time

	return nil
}

// MarshalJSON reverts the effect of type aliasing and struct embedding used
// during the marshalling step to make the pattern seamless.
func (le *LogEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonLogEvent{
		(*logEventAlias)(le),
		timestamp{le.Timestamp},
	})
}

// UnmarshalJSON interprets data as an map[string][]byte and ungzip the event
// data to an EventRecord. For constistency with other AWS Lambda events a list
// of EventRecords is built with contextual information and the actual log
// event in it.
func (recs *EventRecords) UnmarshalJSON(data []byte) error {
	var logs struct {
		Data []byte
	}
	err := json.Unmarshal(data, &logs)
	if err != nil {
		return err
	}

	zr := bytes.NewReader(logs.Data)
	r, err := gzip.NewReader(zr)
	if err != nil {
		return err
	}
	defer r.Close()

	s, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	var aux struct {
		*EventRecord
		LogEvents []*LogEvent
	}
	err = json.Unmarshal(s, &aux)
	if err != nil {
		return err
	}

	*recs = make(EventRecords, len(aux.LogEvents))
	for i, evt := range aux.LogEvents {
		(*recs)[i] = &EventRecord{
			Owner:               aux.Owner,
			LogGroup:            aux.LogGroup,
			LogStream:           aux.LogStream,
			LogEvent:            evt,
			MessageType:         aux.MessageType,
			SubscriptionFilters: aux.SubscriptionFilters,
		}
	}

	return nil
}
