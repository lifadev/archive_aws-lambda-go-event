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

package dynamodbstreamsevt

import (
	"encoding/json"
	"strconv"
	"time"
)

type recordAlias Record

type timestamp struct {
	time.Time
}

// UnmarshalJSON interprets the data as a float64 number, with the integer parts
// being the number of seconds elapsed since January 1, 1970 00:00:00 UTC and
// the fractional part being millisecond offset within the second. It then sets
// *t to a copy of the interpreted time.
func (t *timestamp) UnmarshalJSON(data []byte) error {
	v, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}

	sec := int64(v)
	nsec := int64((v - float64(sec)) * float64(time.Second))

	t.Time = time.Unix(sec, nsec)
	return nil
}

type jsonRecord struct {
	*recordAlias
	ApproximateCreationDateTime timestamp
}

// UnmarshalJSON interprets data as a Record with a special timestamp. It then
// leverages type aliasing and struct embedding to fill Record with an usual
// time.Time.
func (r *Record) UnmarshalJSON(data []byte) error {
	var jr jsonRecord
	if err := json.Unmarshal(data, &jr); err != nil {
		return err
	}

	*r = *(*Record)(jr.recordAlias)
	r.ApproximateCreationDateTime = jr.ApproximateCreationDateTime.Time

	return nil
}

// MarshalJSON reverts the effect of type aliasing and struct embedding used
// during the marshalling step to make the pattern seamless.
func (r *Record) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonRecord{
		(*recordAlias)(r),
		timestamp{r.ApproximateCreationDateTime},
	})
}
