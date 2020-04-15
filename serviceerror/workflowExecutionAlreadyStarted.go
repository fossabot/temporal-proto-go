// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/failure"
	"go.temporal.io/temporal-proto/primitives"
)

type (
	// WorkflowExecutionAlreadyStarted represents workflow execution already started error.
	WorkflowExecutionAlreadyStarted struct {
		Message        string
		StartRequestId string
		RunId          primitives.UUID
		st             *status.Status
	}
)

// NewWorkflowExecutionAlreadyStarted returns new WorkflowExecutionAlreadyStarted error.
func NewWorkflowExecutionAlreadyStarted(message, startRequestId string, runId primitives.UUID) *WorkflowExecutionAlreadyStarted {
	return &WorkflowExecutionAlreadyStarted{
		Message:        message,
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// Error returns string message.
func (e *WorkflowExecutionAlreadyStarted) Error() string {
	return e.Message
}

func (e *WorkflowExecutionAlreadyStarted) status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&failure.WorkflowExecutionAlreadyStarted{
			StartRequestId: e.StartRequestId,
			RunId:          e.RunId,
		},
	)
	return st
}

func newWorkflowExecutionAlreadyStarted(st *status.Status, failure *failure.WorkflowExecutionAlreadyStarted) *WorkflowExecutionAlreadyStarted {
	return &WorkflowExecutionAlreadyStarted{
		Message: st.Message(),
		StartRequestId: failure.StartRequestId,
		RunId:          failure.RunId,
		st:      st,
	}
}
