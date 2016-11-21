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

package s3evt

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go/service/lambda/runtime"
)

// HandlerFunc type is an adapter to allow the use of ordinary functions as
// Amazon S3 events handlers. If f is a function with the appropriate
// signature, HandlerFunc(f) is a Handler that calls f after unmarshaling the
// raw json event.
type HandlerFunc func(*Event, *runtime.Context) (interface{}, error)

// HandleLambda calls f(evt, ctx)
func (f HandlerFunc) HandleLambda(revt json.RawMessage, rctx *runtime.Context) (interface{}, error) {
	evt := &Event{}
	if err := json.Unmarshal(revt, evt); err != nil {
		return nil, err
	}
	return f(evt, rctx)
}
