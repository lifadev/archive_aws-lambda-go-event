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
	"strconv"
	"time"
)

type recordAlias InputRecord

type timestamp struct {
	time.Time
}

// UnmarshalJSON interprets the data as a int64 number, with milliseconds
// precision. It then sets *t to a copy of the interpreted time.
func (t *timestamp) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	t.Time = time.Unix(0, v*int64(time.Millisecond))
	return nil
}

type jsonRecord struct {
	*recordAlias
	ApproximateArrivalTimestamp timestamp
}

// UnmarshalJSON interprets data as a Record with a special timestamp. It then
// leverages type aliasing and struct embedding to fill Record with an usual
// time.Time.
func (r *InputRecord) UnmarshalJSON(data []byte) error {
	var jr jsonRecord
	if err := json.Unmarshal(data, &jr); err != nil {
		return err
	}

	*r = *(*InputRecord)(jr.recordAlias)
	r.ApproximateArrivalTimestamp = jr.ApproximateArrivalTimestamp.Time

	return nil
}

// MarshalJSON reverts the effect of type aliasing and struct embedding used
// during the marshalling step to make the pattern seamless.
func (r *InputRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonRecord{
		(*recordAlias)(r),
		timestamp{r.ApproximateArrivalTimestamp},
	})
}
