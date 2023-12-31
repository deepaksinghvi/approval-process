// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/event_type.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Whenever this list of events is changed do change the function shouldBufferEvent in mutableStateBuilder.go to make sure to do the correct event ordering
type EventType int32

const (
	// Place holder and should never appear in a Workflow execution history
	EVENT_TYPE_UNSPECIFIED EventType = 0
	// Workflow execution has been triggered/started
	// It contains Workflow execution inputs, as well as Workflow timeout configurations
	EVENT_TYPE_WORKFLOW_EXECUTION_STARTED EventType = 1
	// Workflow execution has successfully completed and contains Workflow execution results
	EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED EventType = 2
	// Workflow execution has unsuccessfully completed and contains the Workflow execution error
	EVENT_TYPE_WORKFLOW_EXECUTION_FAILED EventType = 3
	// Workflow execution has timed out by the Temporal Server
	// Usually due to the Workflow having not been completed within timeout settings
	EVENT_TYPE_WORKFLOW_EXECUTION_TIMED_OUT EventType = 4
	// Workflow Task has been scheduled and the SDK client should now be able to process any new history events
	EVENT_TYPE_WORKFLOW_TASK_SCHEDULED EventType = 5
	// Workflow Task has started and the SDK client has picked up the Workflow Task and is processing new history events
	EVENT_TYPE_WORKFLOW_TASK_STARTED EventType = 6
	// Workflow Task has completed
	// The SDK client picked up the Workflow Task and processed new history events
	// SDK client may or may not ask the Temporal Server to do additional work, such as:
	// EVENT_TYPE_ACTIVITY_TASK_SCHEDULED
	// EVENT_TYPE_TIMER_STARTED
	// EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES
	// EVENT_TYPE_MARKER_RECORDED
	// EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED
	// EVENT_TYPE_WORKFLOW_EXECUTION_COMPLETED
	// EVENT_TYPE_WORKFLOW_EXECUTION_FAILED
	// EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED
	// EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW
	EVENT_TYPE_WORKFLOW_TASK_COMPLETED EventType = 7
	// Workflow Task encountered a timeout
	// Either an SDK client with a local cache was not available at the time, or it took too long for the SDK client to process the task
	EVENT_TYPE_WORKFLOW_TASK_TIMED_OUT EventType = 8
	// Workflow Task encountered a failure
	// Usually this means that the Workflow was non-deterministic
	// However, the Workflow reset functionality also uses this event
	EVENT_TYPE_WORKFLOW_TASK_FAILED EventType = 9
	// Activity Task was scheduled
	// The SDK client should pick up this activity task and execute
	// This event type contains activity inputs, as well as activity timeout configurations
	EVENT_TYPE_ACTIVITY_TASK_SCHEDULED EventType = 10
	// Activity Task has started executing
	// The SDK client has picked up the Activity Task and is processing the Activity invocation
	EVENT_TYPE_ACTIVITY_TASK_STARTED EventType = 11
	// Activity Task has finished successfully
	// The SDK client has picked up and successfully completed the Activity Task
	// This event type contains Activity execution results
	EVENT_TYPE_ACTIVITY_TASK_COMPLETED EventType = 12
	// Activity Task has finished unsuccessfully
	// The SDK picked up the Activity Task but unsuccessfully completed it
	// This event type contains Activity execution errors
	EVENT_TYPE_ACTIVITY_TASK_FAILED EventType = 13
	// Activity has timed out according to the Temporal Server
	// Activity did not complete within the timeout settings
	EVENT_TYPE_ACTIVITY_TASK_TIMED_OUT EventType = 14
	// A request to cancel the Activity has occurred
	// The SDK client will be able to confirm cancellation of an Activity during an Activity heartbeat
	EVENT_TYPE_ACTIVITY_TASK_CANCEL_REQUESTED EventType = 15
	// Activity has been cancelled
	EVENT_TYPE_ACTIVITY_TASK_CANCELED EventType = 16
	// A timer has started
	EVENT_TYPE_TIMER_STARTED EventType = 17
	// A timer has fired
	EVENT_TYPE_TIMER_FIRED EventType = 18
	// A time has been cancelled
	EVENT_TYPE_TIMER_CANCELED EventType = 19
	// A request has been made to cancel the Workflow execution
	EVENT_TYPE_WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = 20
	// SDK client has confirmed the cancellation request and the Workflow execution has been cancelled
	EVENT_TYPE_WORKFLOW_EXECUTION_CANCELED EventType = 21
	// Workflow has requested that the Temporal Server try to cancel another Workflow
	EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = 22
	// Temporal Server could not cancel the targeted Workflow
	// This is usually because the target Workflow could not be found
	EVENT_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED EventType = 23
	// Temporal Server has successfully requested the cancellation of the target Workflow
	EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_CANCEL_REQUESTED EventType = 24
	// A marker has been recorded.
	// This event type is transparent to the Temporal Server
	// The Server will only store it and will not try to understand it.
	EVENT_TYPE_MARKER_RECORDED EventType = 25
	// Workflow has received a Signal event
	// The event type contains the Signal name, as well as a Signal payload
	EVENT_TYPE_WORKFLOW_EXECUTION_SIGNALED EventType = 26
	// Workflow execution has been forcefully terminated
	// This is usually because the terminate Workflow API was called
	EVENT_TYPE_WORKFLOW_EXECUTION_TERMINATED EventType = 27
	// Workflow has successfully completed and a new Workflow has been started within the same transaction
	// Contains last Workflow execution results as well as new Workflow execution inputs
	EVENT_TYPE_WORKFLOW_EXECUTION_CONTINUED_AS_NEW EventType = 28
	// Temporal Server will try to start a child Workflow
	EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_INITIATED EventType = 29
	// Child Workflow execution cannot be started/triggered
	// Usually due to a child Workflow ID collision
	EVENT_TYPE_START_CHILD_WORKFLOW_EXECUTION_FAILED EventType = 30
	// Child Workflow execution has successfully started/triggered
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_STARTED EventType = 31
	// Child Workflow execution has successfully completed
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_COMPLETED EventType = 32
	// Child Workflow execution has unsuccessfully completed
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_FAILED EventType = 33
	// Child Workflow execution has been cancelled
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_CANCELED EventType = 34
	// Child Workflow execution has timed out by the Temporal Server
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TIMED_OUT EventType = 35
	// Child Workflow execution has been terminated
	EVENT_TYPE_CHILD_WORKFLOW_EXECUTION_TERMINATED EventType = 36
	// Temporal Server will try to Signal the targeted Workflow
	// Contains the Signal name, as well as a Signal payload
	EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_INITIATED EventType = 37
	// Temporal Server cannot Signal the targeted Workflow
	// Usually because the Workflow could not be found
	EVENT_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED EventType = 38
	// Temporal Server has successfully Signaled the targeted Workflow
	EVENT_TYPE_EXTERNAL_WORKFLOW_EXECUTION_SIGNALED EventType = 39
	// Workflow search attributes should be updated and synchronized with the visibility store
	EVENT_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES EventType = 40
	// An update was accepted (i.e. validated)
	EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_ACCEPTED EventType = 41
	// An update was rejected (i.e. failed validation)
	EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_REJECTED EventType = 42
	// An update completed
	EVENT_TYPE_WORKFLOW_EXECUTION_UPDATE_COMPLETED EventType = 43
	// Some property or properties of the workflow as a whole have changed by non-workflow code.
	// The distinction of external vs. command-based modification is important so the SDK can
	// maintain determinism when using the command-based approach.
	EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED_EXTERNALLY EventType = 44
	// Some property or properties of an already-scheduled activity have changed by non-workflow code.
	// The distinction of external vs. command-based modification is important so the SDK can
	// maintain determinism when using the command-based approach.
	EVENT_TYPE_ACTIVITY_PROPERTIES_MODIFIED_EXTERNALLY EventType = 45
	// Workflow properties modified by user workflow code
	EVENT_TYPE_WORKFLOW_PROPERTIES_MODIFIED EventType = 46
)

var EventType_name = map[int32]string{
	0:  "Unspecified",
	1:  "WorkflowExecutionStarted",
	2:  "WorkflowExecutionCompleted",
	3:  "WorkflowExecutionFailed",
	4:  "WorkflowExecutionTimedOut",
	5:  "WorkflowTaskScheduled",
	6:  "WorkflowTaskStarted",
	7:  "WorkflowTaskCompleted",
	8:  "WorkflowTaskTimedOut",
	9:  "WorkflowTaskFailed",
	10: "ActivityTaskScheduled",
	11: "ActivityTaskStarted",
	12: "ActivityTaskCompleted",
	13: "ActivityTaskFailed",
	14: "ActivityTaskTimedOut",
	15: "ActivityTaskCancelRequested",
	16: "ActivityTaskCanceled",
	17: "TimerStarted",
	18: "TimerFired",
	19: "TimerCanceled",
	20: "WorkflowExecutionCancelRequested",
	21: "WorkflowExecutionCanceled",
	22: "RequestCancelExternalWorkflowExecutionInitiated",
	23: "RequestCancelExternalWorkflowExecutionFailed",
	24: "ExternalWorkflowExecutionCancelRequested",
	25: "MarkerRecorded",
	26: "WorkflowExecutionSignaled",
	27: "WorkflowExecutionTerminated",
	28: "WorkflowExecutionContinuedAsNew",
	29: "StartChildWorkflowExecutionInitiated",
	30: "StartChildWorkflowExecutionFailed",
	31: "ChildWorkflowExecutionStarted",
	32: "ChildWorkflowExecutionCompleted",
	33: "ChildWorkflowExecutionFailed",
	34: "ChildWorkflowExecutionCanceled",
	35: "ChildWorkflowExecutionTimedOut",
	36: "ChildWorkflowExecutionTerminated",
	37: "SignalExternalWorkflowExecutionInitiated",
	38: "SignalExternalWorkflowExecutionFailed",
	39: "ExternalWorkflowExecutionSignaled",
	40: "UpsertWorkflowSearchAttributes",
	41: "WorkflowExecutionUpdateAccepted",
	42: "WorkflowExecutionUpdateRejected",
	43: "WorkflowExecutionUpdateCompleted",
	44: "WorkflowPropertiesModifiedExternally",
	45: "ActivityPropertiesModifiedExternally",
	46: "WorkflowPropertiesModified",
}

var EventType_value = map[string]int32{
	"Unspecified":                                     0,
	"WorkflowExecutionStarted":                        1,
	"WorkflowExecutionCompleted":                      2,
	"WorkflowExecutionFailed":                         3,
	"WorkflowExecutionTimedOut":                       4,
	"WorkflowTaskScheduled":                           5,
	"WorkflowTaskStarted":                             6,
	"WorkflowTaskCompleted":                           7,
	"WorkflowTaskTimedOut":                            8,
	"WorkflowTaskFailed":                              9,
	"ActivityTaskScheduled":                           10,
	"ActivityTaskStarted":                             11,
	"ActivityTaskCompleted":                           12,
	"ActivityTaskFailed":                              13,
	"ActivityTaskTimedOut":                            14,
	"ActivityTaskCancelRequested":                     15,
	"ActivityTaskCanceled":                            16,
	"TimerStarted":                                    17,
	"TimerFired":                                      18,
	"TimerCanceled":                                   19,
	"WorkflowExecutionCancelRequested":                20,
	"WorkflowExecutionCanceled":                       21,
	"RequestCancelExternalWorkflowExecutionInitiated": 22,
	"RequestCancelExternalWorkflowExecutionFailed":    23,
	"ExternalWorkflowExecutionCancelRequested":        24,
	"MarkerRecorded":                                  25,
	"WorkflowExecutionSignaled":                       26,
	"WorkflowExecutionTerminated":                     27,
	"WorkflowExecutionContinuedAsNew":                 28,
	"StartChildWorkflowExecutionInitiated":            29,
	"StartChildWorkflowExecutionFailed":               30,
	"ChildWorkflowExecutionStarted":                   31,
	"ChildWorkflowExecutionCompleted":                 32,
	"ChildWorkflowExecutionFailed":                    33,
	"ChildWorkflowExecutionCanceled":                  34,
	"ChildWorkflowExecutionTimedOut":                  35,
	"ChildWorkflowExecutionTerminated":                36,
	"SignalExternalWorkflowExecutionInitiated":        37,
	"SignalExternalWorkflowExecutionFailed":           38,
	"ExternalWorkflowExecutionSignaled":               39,
	"UpsertWorkflowSearchAttributes":                  40,
	"WorkflowExecutionUpdateAccepted":                 41,
	"WorkflowExecutionUpdateRejected":                 42,
	"WorkflowExecutionUpdateCompleted":                43,
	"WorkflowPropertiesModifiedExternally":            44,
	"ActivityPropertiesModifiedExternally":            45,
	"WorkflowPropertiesModified":                      46,
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b482d2737d9259e4, []int{0}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.EventType", EventType_name, EventType_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/event_type.proto", fileDescriptor_b482d2737d9259e4)
}

var fileDescriptor_b482d2737d9259e4 = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0x5f, 0x4f, 0xd3, 0x5e,
	0x18, 0x5e, 0xf9, 0xfd, 0x44, 0x39, 0x2a, 0x96, 0xa3, 0xfc, 0x9b, 0x50, 0xfe, 0x0f, 0x18, 0xd0,
	0x31, 0x20, 0x62, 0x86, 0x89, 0x96, 0xf6, 0x9d, 0x1c, 0xd9, 0xda, 0x7a, 0x7a, 0x3a, 0xc0, 0x9b,
	0x06, 0x93, 0xc5, 0x2c, 0x11, 0xd6, 0xe0, 0x24, 0xe1, 0xce, 0x4f, 0x60, 0xfc, 0x18, 0xc6, 0x8f,
	0xe0, 0x27, 0xf0, 0x92, 0x4b, 0x2e, 0x65, 0xdc, 0x18, 0xaf, 0xf8, 0x08, 0xa6, 0xa5, 0x5b, 0x4b,
	0xd9, 0xba, 0xea, 0xdd, 0x92, 0xf3, 0x3c, 0xcf, 0x79, 0x9f, 0xf7, 0x3c, 0xef, 0xfa, 0xa2, 0x54,
	0xad, 0x7c, 0x60, 0x57, 0x8f, 0xf6, 0xdf, 0x67, 0xf6, 0xed, 0x4a, 0xa6, 0x7c, 0xf8, 0xf1, 0xe0,
	0x43, 0xe6, 0x38, 0x9b, 0x29, 0x1f, 0x97, 0x0f, 0x6b, 0x56, 0xed, 0xc4, 0x2e, 0x8b, 0xf6, 0x51,
	0xb5, 0x56, 0xc5, 0xfd, 0x0d, 0x9c, 0xb8, 0x6f, 0x57, 0x44, 0x17, 0x27, 0x1e, 0x67, 0xd3, 0x9f,
	0xfb, 0x50, 0x0f, 0x38, 0x58, 0x76, 0x62, 0x97, 0x71, 0x12, 0x0d, 0x40, 0x09, 0x54, 0x66, 0xb1,
	0x3d, 0x1d, 0x2c, 0x53, 0x35, 0x74, 0x90, 0x49, 0x9e, 0x80, 0xc2, 0x27, 0xf0, 0x3c, 0x9a, 0x09,
	0x9c, 0xed, 0x68, 0x74, 0x3b, 0x5f, 0xd0, 0x76, 0x2c, 0xd8, 0x05, 0xd9, 0x64, 0x44, 0x53, 0x2d,
	0x83, 0x49, 0x94, 0x81, 0xc2, 0x73, 0x78, 0x01, 0xcd, 0x46, 0x43, 0x65, 0xad, 0xa8, 0x17, 0xc0,
	0x01, 0x77, 0xe1, 0x39, 0x34, 0x1d, 0x0d, 0xce, 0x4b, 0xa4, 0x00, 0x0a, 0xff, 0x5f, 0x67, 0x59,
	0x46, 0x8a, 0xa0, 0x58, 0x9a, 0xc9, 0xf8, 0xff, 0x71, 0x0a, 0x4d, 0xb6, 0x02, 0x33, 0xc9, 0xd8,
	0xb6, 0x0c, 0x79, 0x0b, 0x14, 0xd3, 0x11, 0xbd, 0x85, 0xa7, 0xd1, 0x78, 0x7b, 0x9c, 0xe7, 0xa8,
	0x3b, 0x52, 0xcd, 0x37, 0x73, 0x3b, 0x12, 0xe7, 0x57, 0x77, 0x07, 0x4f, 0xa1, 0xb1, 0xb6, 0x38,
	0xcf, 0x6f, 0x4f, 0x48, 0x4c, 0x92, 0x19, 0x29, 0x11, 0xb6, 0x17, 0xb6, 0x80, 0x42, 0x16, 0x42,
	0x38, 0xcf, 0xc2, 0xdd, 0x48, 0x35, 0xdf, 0xc2, 0xbd, 0x50, 0x69, 0xd7, 0x71, 0x5e, 0x69, 0xf7,
	0x23, 0xc5, 0x7c, 0x9f, 0xbd, 0x78, 0x09, 0xcd, 0xb7, 0xbf, 0x54, 0x52, 0x65, 0x28, 0x58, 0x14,
	0x5e, 0x9b, 0x60, 0x38, 0x77, 0x3f, 0xc0, 0x33, 0x68, 0xa2, 0x03, 0x1c, 0x14, 0x9e, 0xc7, 0x23,
	0x68, 0x28, 0x00, 0x73, 0xee, 0xa3, 0x4d, 0xa3, 0x7d, 0xa1, 0x10, 0x5f, 0x9d, 0xe6, 0x09, 0x05,
	0x85, 0xc7, 0x78, 0x14, 0x0d, 0xdf, 0x38, 0x6b, 0x0a, 0x3f, 0xc4, 0x2b, 0x48, 0xec, 0x10, 0xdc,
	0x70, 0xcd, 0x8f, 0x70, 0x1a, 0xa5, 0xe2, 0x70, 0x40, 0xe1, 0xfb, 0xb1, 0x8c, 0x9e, 0x07, 0xb0,
	0x9e, 0x4a, 0x43, 0x14, 0x76, 0x19, 0x50, 0x55, 0x2a, 0xb4, 0xd2, 0x20, 0x2a, 0x61, 0x44, 0x72,
	0x2e, 0x1c, 0xc0, 0x2f, 0xd0, 0xb3, 0x7f, 0x13, 0xf1, 0x5e, 0x6f, 0x10, 0x6f, 0xa0, 0xf5, 0x80,
	0x42, 0x14, 0xe5, 0x86, 0xdf, 0x21, 0x2c, 0xa0, 0x64, 0x80, 0x5c, 0x94, 0xe8, 0x36, 0x50, 0x8b,
	0x82, 0xac, 0x51, 0x05, 0x14, 0x7e, 0xb8, 0x73, 0x3f, 0x0c, 0xf2, 0x52, 0x95, 0x9c, 0x42, 0x92,
	0x78, 0x11, 0xcd, 0x75, 0x98, 0x68, 0xa0, 0x45, 0xa2, 0xba, 0xc6, 0x1f, 0xc7, 0x78, 0x1d, 0x4d,
	0x65, 0x44, 0x35, 0x41, 0xb1, 0x24, 0xc3, 0x52, 0x61, 0x87, 0x1f, 0xc1, 0xeb, 0x68, 0x35, 0xc0,
	0x71, 0x43, 0x62, 0xc9, 0x5b, 0xa4, 0xa0, 0x44, 0x77, 0x79, 0x14, 0xaf, 0xa1, 0xe5, 0xf8, 0x44,
	0xaf, 0xb3, 0x02, 0xce, 0xa0, 0x85, 0x00, 0xab, 0x2d, 0xbe, 0x11, 0xd6, 0x31, 0x9c, 0x45, 0x4b,
	0x71, 0x08, 0xfe, 0x80, 0x8e, 0x63, 0x11, 0xa5, 0xe3, 0x50, 0xbc, 0x9a, 0x26, 0xf0, 0x32, 0x5a,
	0x8c, 0x75, 0x45, 0x23, 0xa6, 0x93, 0x71, 0x8b, 0xf2, 0x07, 0x7d, 0x2a, 0xf4, 0x36, 0xed, 0x29,
	0xfe, 0x7b, 0x4e, 0x87, 0x62, 0x78, 0x15, 0x8b, 0x98, 0x53, 0x30, 0x83, 0x9f, 0xa2, 0xb5, 0xbf,
	0x23, 0x7b, 0xfd, 0x48, 0xe1, 0x55, 0x94, 0x89, 0x99, 0xfe, 0x66, 0x52, 0x67, 0x43, 0x4d, 0x34,
	0x75, 0x03, 0x28, 0xf3, 0x29, 0x06, 0x48, 0x54, 0xde, 0xb2, 0x24, 0xc6, 0x28, 0xd9, 0x34, 0x19,
	0x18, 0xfc, 0x5c, 0xa8, 0x89, 0x2d, 0xd4, 0x4d, 0x5d, 0x91, 0x98, 0xf3, 0x2f, 0x27, 0x83, 0xee,
	0x78, 0x9a, 0x8f, 0x4d, 0xa1, 0xf0, 0x0a, 0x64, 0x87, 0x92, 0xee, 0x3c, 0x13, 0x1e, 0xc5, 0x0f,
	0xd0, 0x02, 0x7e, 0x82, 0x56, 0x5a, 0x71, 0x74, 0xaa, 0xe9, 0x40, 0x19, 0x01, 0xc3, 0x2a, 0x6a,
	0x8a, 0xfb, 0xd9, 0x6f, 0x76, 0xa7, 0xb0, 0xc7, 0x2f, 0x86, 0x78, 0xcd, 0x7f, 0xe7, 0x0e, 0xbc,
	0xa5, 0x76, 0xdf, 0xed, 0x16, 0x3c, 0x5e, 0xdc, 0xfc, 0xce, 0x9d, 0x9e, 0x0b, 0x89, 0xb3, 0x73,
	0x21, 0x71, 0x79, 0x2e, 0x70, 0x9f, 0xea, 0x02, 0xf7, 0xb5, 0x2e, 0x70, 0x3f, 0xea, 0x02, 0x77,
	0x5a, 0x17, 0xb8, 0x9f, 0x75, 0x81, 0xfb, 0x55, 0x17, 0x12, 0x97, 0x75, 0x81, 0xfb, 0x72, 0x21,
	0x24, 0x4e, 0x2f, 0x84, 0xc4, 0xd9, 0x85, 0x90, 0x40, 0x43, 0x95, 0xaa, 0xd8, 0x72, 0xc3, 0xd9,
	0xec, 0x6d, 0xae, 0x37, 0xba, 0xb3, 0x08, 0xe9, 0xdc, 0x9b, 0x89, 0x77, 0x01, 0x6c, 0xa5, 0x7a,
	0x6d, 0x71, 0xda, 0x70, 0x7f, 0x7c, 0xeb, 0x1a, 0x64, 0x1e, 0xa0, 0x52, 0x15, 0x25, 0xbb, 0x22,
	0x82, 0x2b, 0x57, 0xca, 0xfe, 0xee, 0x4a, 0xfa, 0x27, 0xb9, 0x9c, 0x64, 0x57, 0x72, 0x39, 0xf7,
	0x2c, 0x97, 0x2b, 0x65, 0xdf, 0x76, 0xbb, 0xbb, 0xd6, 0xea, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbd, 0xfa, 0x22, 0x03, 0x95, 0x09, 0x00, 0x00,
}

func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
