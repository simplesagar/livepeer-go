// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-go/models/components"
)

type GetTaskRequest struct {
	// ID of the task
	TaskID string `pathParam:"style=simple,explode=false,name=taskId"`
}

func (o *GetTaskRequest) GetTaskID() string {
	if o == nil {
		return ""
	}
	return o.TaskID
}

type GetTaskResponse struct {
	HTTPMeta components.HTTPMetadata
	// Success
	Task *components.Task
}

func (o *GetTaskResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetTaskResponse) GetTask() *components.Task {
	if o == nil {
		return nil
	}
	return o.Task
}
