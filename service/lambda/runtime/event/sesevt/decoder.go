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
	"bytes"
	"time"
)

// UnmarshalJSON interprets the data as a RFC322 date and time. It then sets *t
// to a copy of the interpreted time.
// See https://tools.ietf.org/search/rfc5322#section-3.3 for more informations.
func (t *MailTimestamp) UnmarshalJSON(data []byte) error {
	v, err := time.Parse("Mon, _2 Jan 2006 15:04:05 -0700", string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}

// UnmarshalJSON interprets the data as a RFC3339Nano time. It then sets *t to a
// copy of the interpreted time.
func (t *Timestamp) UnmarshalJSON(data []byte) error {
	v, err := time.Parse(time.RFC3339Nano, string(bytes.Trim(data, `"`)))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}
