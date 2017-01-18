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

/*
Package cloudwatchschedevt allows you to write AWS Lambda functions executed on
a regular, scheduled basis using the schedule event capability in Amazon
CloudWatch Events.

This package works only when forwarding the whole event to your Lambda function.
Otherwise you will have to handle the event manually.
*/
package cloudwatchschedevt
