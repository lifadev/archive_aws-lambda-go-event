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

package codepipelineevt

import "encoding/json"

// AWSSessionCredentials represents an AWS session credentials object.
// These credentials are temporary credentials that are issued by AWS Secure
// Token Service (STS). They can be used to access input and output artifacts in
// the Amazon S3 bucket used to store artifact for the pipeline in AWS
// CodePipeline.
type AWSSessionCredentials struct {
	// The access key for the session.
	AccessKeyID string

	// The secret access key for the session.
	SecretAccessKey string

	// The token for the session.
	SessionToken string
}

// ActionConfiguration represents information about an action configuration.
type ActionConfiguration struct {
	// The configuration data for the action.
	Configuration *struct {
		// AWS Lambda function name.
		FunctionName string

		// An arbitrary string of user formatted configuration for the
		// AWS Lambda function to complete the job.
		UserParameters string
	}
}

// S3ArtifactLocation represents the location of the Amazon S3 bucket that
// contains a revision.
type S3ArtifactLocation struct {
	// The name of the Amazon S3 bucket.
	BucketName string

	// The key of the object in the Amazon S3 bucket, which uniquely
	// identifies the object in the bucket.
	ObjectKey string
}

// ArtifactLocation represents information about the location of an artifact.
type ArtifactLocation struct {
	// The type of artifact in the location.
	Type string

	// The Amazon S3 bucket that contains the artifact.
	S3Location *S3ArtifactLocation
}

// Artifact represents information about an artifact that will be worked upon
// by actions in the pipeline.
type Artifact struct {
	// The artifact's name.
	Name string

	// The artifact's revision ID.
	// Depending on the type of object, this could be a commit ID (GitHub)
	// or a revision ID (Amazon S3).
	Revision string

	// The location of an artifact.
	Location *ArtifactLocation
}

// Data represents additional information about an AWS CodePipeline job required
// for the AWS Lambda function to complete the job.
type Data struct {
	// Information about an action configuration.
	ActionConfiguration *ActionConfiguration

	// A system-generated token, such as an AWS CodeDeploy deployment ID,
	// that a job requires in order to continue the job asynchronously.
	ContinuationToken string

	// An AWS session credentials object.
	// These credentials are temporary credentials that are issued by AWS
	// Secure Token Service (STS). They can be used to access input and
	// output artifacts in the Amazon S3 bucket used to store artifact for
	// the pipeline in AWS CodePipeline.
	ArtifactCredentials *AWSSessionCredentials

	// The artifact supplied to the job.
	InputArtifacts []*Artifact

	// The output of the job.
	OutputArtifacts []*Artifact
}

// Job represents information about the details of an AWS CodePipeline job.
type Job struct {
	// The unique system-generated ID of the job.
	ID string

	// The ID of the AWS account to use when performing the job.
	AccountID string

	// Additional information about the job.
	Data *Data
}

// Event represents an AWS CodePipeline event.
type Event struct {
	// Information about the details of the job.
	Job *Job
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
