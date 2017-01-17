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

// UnmarshalJSON interprets the data as an int64 being the number of
// milliseconds elapsed since January 1, 1970 00:00:00 UTC. It then sets *t to
// a copy of the interpreted time.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	sec := v / 1000
	nsec := (v - sec*1000) * int64(time.Millisecond)

	t.Time = time.Unix(sec, nsec)

	return nil
}

// UnmarshalJSON interprets data as an map[string][]byte and ungzip the event
// data to an EventRecord. For constistency with other AWS Lambda events a list
// of EventRecords is built with contextual information and the actual log
// event in it.
func (recs *EventRecords) UnmarshalJSON(data []byte) error {
	var logs map[string][]byte

	err := json.Unmarshal(data, &logs)
	if err != nil {
		return err
	}

	zr := bytes.NewReader(logs["data"])
	r, err := gzip.NewReader(zr)
	if err != nil {
		return err
	}
	defer r.Close()

	s, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	aux := &struct {
		*EventRecord
		LogEvents []*LogEvent
	}{}

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
