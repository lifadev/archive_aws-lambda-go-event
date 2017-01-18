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

// Header represents an header value
type Header struct {
	// The name of the header
	Name string

	// The value of the header
	Value string
}

// CommonHeaders represents a list of headers common to all emails.
// See https://tools.ietf.org/search/rfc5322 for more information.
type CommonHeaders struct {
	// The message identifier is a globally unique identifier for a message. The
	// generator of the message identifier guarantee that the id is unique.
	// Semantically, the angle bracket characters are not part of the MessageID;
	// the MessageID is what is contained between the two angle bracket
	// characters.
	// Format: <id-left@id-right>
	MessageID string

	// Specifies the date and time at which the creator of the message indicated
	// that the message was complete and ready to enter the mail delivery system.
	// For instance, this might be the time that a user pushes the "send" or
	// "submit" button in an application program.
	Date time.Time `json:"-"`

	// Contains a short string identifying the topic of the message. When used in
	// a reply, the field body MAY start with the string "Re: " (an abbreviation
	// of the Latin "in re", meaning "in the matter of") followed by the contents
	// of the "Subject:" field body of the original message.
	Subject string

	// Specifies the author(s) of the message, that is, the mailbox(es) of the
	// person(s) or system(s) responsible for the writing of the message.
	From []string

	// Contains the address(es) of the primary recipient(s) of the message.
	To []string

	// Intended to show the envelope address of the real sender as opposed to the
	// sender used for replying (the From: and Reply-To: headers).
	ReturnPath string
}

// Verdict encapsulates information about the check that was executed.
type Verdict struct {
	// The status of the verdict, can be one of "PASS", "FAIL", "GRAY" or
	// "PROCESSING_FAILED". See each check to know more about the signification
	// of each status.
	Status string
}

// Action encapsulates information about the action that was executed.
type Action struct {
	// Indicates the type of action that was executed.
	// In our case value is always "Lambda".
	Type string

	// Contains the invocation type of the AWS Lambda function.
	// Possible values are "RequestResponse"" and "Event". Present only for the
	// AWS Lambda action type.
	InvocationType string

	// Contains the ARN of the AWS Lambda function that was triggered.
	// Present only for the AWS Lambda action type.
	FunctionARN string
}

// Mail contains information about the email to which the notification pertains.
type Mail struct {
	// The unique ID assigned to the email by Amazon SES. If the email was
	// delivered to Amazon S3, the message ID is also the Amazon S3 object key
	// that was used to write the message to your Amazon S3 bucket.
	// Note that this message ID was assigned by Amazon SES. You can find the
	// message ID of the original email in the headers and CommonHeaders object of
	// the Mail object.
	MessageID string

	// The time at which the email was received.
	Timestamp time.Time

	// The email address from which the email was sent (the envelope MAIL FROM
	// address).
	Source string

	// A list of email addresses that were recipients of the original mail.
	Destination []string

	// A list of headers common to all emails.
	// Note that any message ID within the CommonHeaders object is from the
	// original message that you passed to Amazon SES. The message ID that Amazon
	// SES subsequently assigned to the message is in the MessageID field of the
	// Mail object.
	CommonHeaders *CommonHeaders

	// Specifies whether the headers were truncated in the notification, which
	// will happen if the headers are larger than 10 KB.
	// Possible values are true or false.
	HeadersTruncated bool

	// A list of the email's original headers. Each header in the list has a name
	// field and a value field.
	// Note that any message ID within the headers field is from the original
	// message that you passed to Amazon SES. The message ID that Amazon SES
	// subsequently assigned to the message is in the MessageID field of the Mail
	// object.
	// See https://tools.ietf.org/search/rfc5322 for more information about
	// headers.
	Headers []*Header
}

// Receipt contains information about the email delivery.
type Receipt struct {
	// A list of the recipient addresses for this delivery. This list might be a
	// subset of the recipients to which the mail was addressed.
	Recipients []string

	// Specifies when the action was triggered.
	Timestamp time.Time

	// Specifies the period, in milliseconds, from the time Amazon SES received
	// the message to the time it triggered the action.
	ProcessingTimeMillis int

	// Encapsulates information about the action that was executed.
	Action *Action

	// Indicates whether the message is spam. Possible values are as follows:
	// - PASS: the check succeeded.
	// - FAIL: the check failed.
	// - GRAY: Amazon SES scanned the email but could not determine with
	//   confidence whether it is spam.
	// - PROCESSING_FAILED: Amazon SES is unable to scan the content of the
	//   email. For example, the email is not a valid MIME message.
	SpamVerdict *Verdict

	// Indicates whether the DomainKeys Identified Mail (DKIM) check passed.
	// Possible values are as follows:
	// - PASS: the check succeeded.
	// - FAIL: the check failed.
	// - GRAY: the message is not DKIM-signed.
	// - PROCESSING_FAILED: there is an issue that prevents Amazon SES from
	//   checking the DKIM signature. For example, DNS queries are failing or the
	//   DKIM signature header is not formatted properly.
	DKIMVerdict *Verdict

	// Indicates whether the Sender Policy Framework (SPF) check passed. Possible
	// values are as follows:
	// - PASS: the check succeeded.
	// - FAIL: the check failed.
	// - GRAY: there is no SPF policy under the domain used in the MAIL FROM
	//   command.
	// - PROCESSING_FAILED: there is an issue that prevents Amazon SES from
	//   checking the SPF record. For example, DNS queries are failing.
	SPFVerdict *Verdict

	// Indicates whether the message contains a virus. Possible values are as
	// follows:
	// - PASS: the check succeeded.
	// - FAIL: the check failed.
	// - GRAY: Amazon SES scanned the email but could not determine with
	//   confidence whether it contains a virus.
	// - PROCESSING_FAILED: Amazon SES is unable to scan the content of the email.
	//   For example, the email is not a valid MIME message.
	VirusVerdict *Verdict
}

// Record represents the unit of data of an Amazon SES message.
// See also http://docs.aws.amazon.com/ses/latest/DeveloperGuide/notification-contents.html
type Record struct {
	// Contains information about the email to which the notification pertains.
	Mail *Mail

	// Contains information about the email delivery.
	Receipt *Receipt
}

// EventRecord provides contextual information about an Amazon SES event.
type EventRecord struct {
	// The event version.
	EventVersion string

	// The event source.
	EventSource string

	// The underlying Amazon SES record associated with the event.
	SES *Record
}

// String returns the string representation.
func (e *EventRecord) String() string {
	s, _ := json.MarshalIndent(e, "", "  ")
	return string(s)
}

// GoString returns the string representation.
func (e *EventRecord) GoString() string {
	return e.String()
}

// Event represents an Amazon SES event.
type Event struct {
	// The list of Amazon SES event records.
	Records []*EventRecord
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
