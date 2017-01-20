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

package cloudformationevt

import "encoding/json"

// Event represents an AWS CloudFormation event.
// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/crpg-ref-requests.html
type Event struct {
	// A unique ID for the request.
	// Combining the StackID with the RequestID forms a value that can be used to
	// uniquely identify a request on a particular custom resource.
	RequestID string

	// The ARN that identifies the stack containing the custom resource.
	// Combining the StackID with the RequestID forms a value that can be used to
	// uniquely identify a request on a particular custom resource.
	StackID string

	// The request type is set by the AWS CloudFormation stack operation
	// (create-stack, update-stack, or delete-stack) that was initiated by the
	// template developer for the stack that contains the custom resource.
	// Must be one of: Create, Update, or Delete.
	// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/crpg-ref-requesttypes.html
	RequestType string

	// The template developer-chosen resource type of the custom resource in the
	// AWS CloudFormation template.
	// Custom resource type names can be up to 60 characters long and can include
	// alphanumeric and the following characters: _@-.
	ResourceType string

	// The template developer-chosen name (logical ID) of the custom resource in
	// the AWS CloudFormation template.
	// This is provided to facilitate communication between the custom resource
	// provider and the template developer.
	LogicalResourceID string

	// A required custom resource provider-defined physical ID that is unique for
	// that provider.
	// Always sent with Update and Delete requests; never sent with Create.
	// See http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cfn-customresource.html#d0e41385
	PhysicalResourceID string

	// This field contains the contents of the Properties object sent by the
	// template developer. Its contents are defined by the custom resource
	// provider.
	ResourceProperties json.RawMessage

	// Used only for Update requests. Contains the resource properties that were
	// declared previous to the update request.
	OldResourceProperties json.RawMessage

	// The response URL identifies a pre-signed Amazon S3 bucket that receives
	// responses from the custom resource provider to AWS CloudFormation.
	ResponseURL string

	// The service token (AWS Lambda function ARN) that is obtained from the
	// custom resource provider to access the service.
	ServiceToken string
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
