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

package apigatewayauthorizerevt

import "encoding/json"

// Identity provides identity information about the API caller.
type Identity struct {
	// The API owner key associated with the API.
	APIKey string

	// The source IP address of the TCP connection making the request to
	// Amazon API Gateway.
	SourceIP string
}

// RequestContext provides contextual information about an Amazon API Gateway
// Proxy Custom Authorizer event.
type RequestContext struct {
	// The path as defined in Amazon API Gateway.
	Path string

	// The AWS account ID associated with the API.
	AccountID string

	// The deployment stage of the API call (for example, Beta or Prod).
	Stage string

	// An automatically generated ID for the API call.
	RequestID string

	// The API caller identification information.
	Identity *Identity

	// The resource path as defined in Amazon API Gateway.
	ResourcePath string

	// The incoming request HTTP method name.
	// Valid values include: DELETE, GET, HEAD, OPTIONS, PATCH, POST, and
	// PUT.
	HTTPMethod string

	// The identifier Amazon API Gateway assigns to the API.
	APIID string

	// The identifier Amazon API Gateway assigns to the resource.
	ResourceID string
}

// Event represents an input to a REQUEST authorizer for an API method with a
// proxy integration.
type Event struct {
	// Type property specifies the authorizer type ("TOKEN" or "REQUEST").
	Type string

	// MethodARN is the ARN of the incoming method request and is populated by
	// API Gateway in accordance with the custom authorizer configuration.
	MethodARN string

	// AuthorizationToken has called supplied token which is allow, deny,
	// unauthorized, or any other string value. An empty string value is the
	// same as unauthorized (only for TOKEN type).
	AuthorizationToken string

	// The resource path with raw placeholders as defined in
	// Amazon API Gateway.
	Resource string

	// The real path corresponding to the path parameters injected into the
	// Resource placeholders.
	Path string

	// The incoming request HTTP method name.
	// Valid values include: DELETE, GET, HEAD, OPTIONS, PATCH, POST, and
	// PUT.
	HTTPMethod string

	// The incoming request HTTP headers.
	// Duplicate entries are not supported.
	Headers map[string]string

	// The incoming request query string parameters.
	// Duplicate entries are not supported.
	QueryStringParameters map[string]string

	// The incoming request path parameters corresponding to the resource
	// path placeholders values as defined in Resource.
	PathParameters map[string]string

	// The name-value pairs defined as configuration attributes associated
	// with the deployment stage of the API.
	StageVariables map[string]string

	// The contextual information associated with the API call.
	RequestContext *RequestContext
}

// Statement represents AWS IAM policy statement.
//
// More detail about IAM policy statement
// http://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html#api-gateway-calling-api-permissions
type Statement struct {
	// Action describes the specific action or actions that will be
	// allowed or denied. Statements must include either an Action or NotAction element.
	Action string

	// Effect is required and specifies whether the statement results
	// in an allow or an explicit deny. Valid values for Effect are Allow and Deny.
	Effect string

	// Resource specifies the object or objects that the statement covers.
	Resource string
}

// PolicyDocument represents an AWS IAM policy document.
type PolicyDocument struct {
	// Version specifies the language syntax rules that are to be used
	// to process this policy. The Version must appear before the Statement element.
	Version string

	// Statement represents an AWS IAM policy statement.
	Statement []Statement
}

// Response represents an output from an Amazon API Gateway Custom Authorizer.
type Response struct {
	// The principal user identification associated with the token sent by the client.
	PrincipalID string `json:"principalId"`

	// PolicyDocument represents an AWS IAM policy document.
	PolicyDocument PolicyDocument `json:"policyDocument"`

	// API Gateway passes the context object from a custom authorizer directly to
	// the backend Lambda function as part of the input event.
	Context map[string]string `json:"context"`
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
