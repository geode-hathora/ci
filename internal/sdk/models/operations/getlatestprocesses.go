// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetLatestProcessesGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetLatestProcessesGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetLatestProcessesRequest struct {
	AppID  *string                `pathParam:"style=simple,explode=false,name=appId"`
	Status []shared.ProcessStatus `queryParam:"style=form,explode=true,name=status"`
	Region []shared.Region        `queryParam:"style=form,explode=true,name=region"`
}

func (o *GetLatestProcessesRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetLatestProcessesRequest) GetStatus() []shared.ProcessStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *GetLatestProcessesRequest) GetRegion() []shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

type GetLatestProcessesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	ProcessV3s []shared.ProcessV3
}

func (o *GetLatestProcessesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLatestProcessesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLatestProcessesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLatestProcessesResponse) GetProcessV3s() []shared.ProcessV3 {
	if o == nil {
		return nil
	}
	return o.ProcessV3s
}
