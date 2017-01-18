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

package sesevt

import (
	"encoding/json"
	"time"
)

type commonHeadersAlias CommonHeaders

type timestamp struct {
	time.Time
}

// UnmarshalJSON interprets the data as a RFC322 date and time. It then sets *t
// to a copy of the interpreted time.
// See https://tools.ietf.org/search/rfc5322#section-3.3 for more informations.
func (t *timestamp) UnmarshalJSON(data []byte) error {
	v, err := time.Parse(`"Mon, _2 Jan 2006 15:04:05 -0700"`, string(data))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}

type jsonCommonHeaders struct {
	*commonHeadersAlias
	Date timestamp
}

// UnmarshalJSON interprets data as a CommonHeaders with a special timestamp.
// It then leverages type aliasing and struct embedding to fill CommonHeaders
// with an usual time.Time.
func (ch *CommonHeaders) UnmarshalJSON(data []byte) error {
	var jch jsonCommonHeaders
	if err := json.Unmarshal(data, &jch); err != nil {
		return err
	}

	*ch = *(*CommonHeaders)(jch.commonHeadersAlias)
	ch.Date = jch.Date.Time

	return nil
}

// MarshalJSON reverts the effect of type aliasing and struct embedding used
// during the marshalling step to make the pattern seamless.
func (ch *CommonHeaders) MarshalJSON() ([]byte, error) {
	return json.Marshal(&jsonCommonHeaders{
		(*commonHeadersAlias)(ch),
		timestamp{ch.Date},
	})
}
