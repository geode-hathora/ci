// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/hathora/ci/internal/sdk/models/shared"
	"net/http"
)

type ListActivePublicLobbiesGlobals struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
}

func (o *ListActivePublicLobbiesGlobals) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

type ListActivePublicLobbiesRequest struct {
	AppID *string `pathParam:"style=simple,explode=false,name=appId"`
	// If omitted, active public lobbies in all regions will be returned.
	Region *shared.Region `queryParam:"style=form,explode=true,name=region"`
}

func (o *ListActivePublicLobbiesRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *ListActivePublicLobbiesRequest) GetRegion() *shared.Region {
	if o == nil {
		return nil
	}
	return o.Region
}

type ListActivePublicLobbiesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Ok
	LobbyV3s []shared.LobbyV3
}

func (o *ListActivePublicLobbiesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListActivePublicLobbiesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListActivePublicLobbiesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListActivePublicLobbiesResponse) GetLobbyV3s() []shared.LobbyV3 {
	if o == nil {
		return nil
	}
	return o.LobbyV3s
}
