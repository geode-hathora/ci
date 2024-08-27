// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type GetDeploymentsV2DeprecatedGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsV2DeprecatedGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentsV2DeprecatedRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *GetDeploymentsV2DeprecatedRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type GetDeploymentsV2DeprecatedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	DeploymentV2s []shared.DeploymentV2
}

func (o *GetDeploymentsV2DeprecatedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetDeploymentsV2DeprecatedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetDeploymentsV2DeprecatedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetDeploymentsV2DeprecatedResponse) GetDeploymentV2s() []shared.DeploymentV2 {
	if o == nil {
		return nil
	}
	return o.DeploymentV2s
}
