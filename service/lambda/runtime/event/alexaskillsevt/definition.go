//
// Copyright 2016 Alsanium, SAS. or its affiliates. All rights reserved.
// Modifications copyright 2018 Kyle Vanek. All rights reserved.
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

package alexaskillsevt

import (
	"encoding/json"
	"time"
)

// EventSession provides session details for an Amazon Alexa Skills Events
type EventSession struct {
	Sessionid   string `json:"sessionId"`
	Application struct {
		Applicationid string `json:"applicationId"`
	} `json:"application"`
	Attributes map[string]interface{} `json:"attributes"`
	User       struct {
		Userid      string                 `json:"userId"`
		Accesstoken interface{}            `json:"accessToken"`
		Permissions map[string]interface{} `json:"permissions"`
	} `json:"user"`
	New bool `json:"new"`
}

// EventRequest provides request infomration for an Amazon Alexa Skills Event
// https://developer.amazon.com/docs/custom-skills/request-types-reference.html
type EventRequest struct {
	Type        string    `json:"type"`
	Requestid   string    `json:"requestId"`
	Timestamp   time.Time `json:"timestamp"`
	DialogState string    `json:"dialogState"`
	Locale      string    `json:"locale"`
	Intent      struct {
		Name  string `json:"name"`
		Slots map[string]struct {
			Name               string `json:"name"`
			Value              string `json:"value"`
			ConfirmationStatus string `json:"confirmationStatus"`
			Resolutions        struct {
				ResolutionsPerAuthority []struct {
					Authority string `json:"authority"`
					Status    struct {
						Code string `json:"code"`
					} `json:"status"`
					Values map[string]struct {
						Name string `json:"name"`
						ID   string `json:"id"`
					} `json:"values"`
				} `json:"resolutionsPerAuthority"`
			} `json:"resolutions"`
		} `json:"slots"`
	} `json:"intent"`
	Reason interface{} `json:"reason"`
	Error  struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}

// Event represents an Amazon Alexa Skills Event.
type Event struct {
	// Skill Event Session
	Session *EventSession `json:"session"`

	// Skill Event Request
	Request *EventRequest `json:"request"`
}

// String returns the string representation.
func (e *Event) String() string {
	s, _ := json.Marshal(e)
	return string(s)
}

// GoString returns the string representation.
func (e *Event) GoString() string {
	return e.String()
}
