// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

package matching

import (
	"context"

	"go.temporal.io/server/api/matchingservice/v1"
	"go.temporal.io/server/common/metrics"
)

type (
	// Engine exposes interfaces for clients to interact with the matching engine
	Engine interface {
		Stop()
		AddWorkflowTask(ctx context.Context, addRequest *matchingservice.AddWorkflowTaskRequest) (syncMatch bool, err error)
		AddActivityTask(ctx context.Context, addRequest *matchingservice.AddActivityTaskRequest) (syncMatch bool, err error)
		PollWorkflowTaskQueue(ctx context.Context, request *matchingservice.PollWorkflowTaskQueueRequest, opMetrics metrics.Handler) (*matchingservice.PollWorkflowTaskQueueResponse, error)
		PollActivityTaskQueue(ctx context.Context, request *matchingservice.PollActivityTaskQueueRequest, opMetrics metrics.Handler) (*matchingservice.PollActivityTaskQueueResponse, error)
		QueryWorkflow(ctx context.Context, request *matchingservice.QueryWorkflowRequest) (*matchingservice.QueryWorkflowResponse, error)
		RespondQueryTaskCompleted(ctx context.Context, request *matchingservice.RespondQueryTaskCompletedRequest, opMetrics metrics.Handler) error
		CancelOutstandingPoll(ctx context.Context, request *matchingservice.CancelOutstandingPollRequest) error
		DescribeTaskQueue(ctx context.Context, request *matchingservice.DescribeTaskQueueRequest) (*matchingservice.DescribeTaskQueueResponse, error)
		ListTaskQueuePartitions(ctx context.Context, request *matchingservice.ListTaskQueuePartitionsRequest) (*matchingservice.ListTaskQueuePartitionsResponse, error)
		UpdateWorkerBuildIdCompatibility(ctx context.Context, request *matchingservice.UpdateWorkerBuildIdCompatibilityRequest) (*matchingservice.UpdateWorkerBuildIdCompatibilityResponse, error)
		GetWorkerBuildIdCompatibility(ctx context.Context, request *matchingservice.GetWorkerBuildIdCompatibilityRequest) (*matchingservice.GetWorkerBuildIdCompatibilityResponse, error)
		InvalidateTaskQueueUserData(ctx context.Context, request *matchingservice.InvalidateTaskQueueUserDataRequest) (*matchingservice.InvalidateTaskQueueUserDataResponse, error)
		GetTaskQueueUserData(ctx context.Context, request *matchingservice.GetTaskQueueUserDataRequest) (*matchingservice.GetTaskQueueUserDataResponse, error)
	}
)
