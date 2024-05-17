// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type InviteStatusRescindedType string

const (
	InviteStatusRescindedTypeRescinded InviteStatusRescindedType = "rescinded"
)

func (e InviteStatusRescindedType) ToPointer() *InviteStatusRescindedType {
	return &e
}
func (e *InviteStatusRescindedType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rescinded":
		*e = InviteStatusRescindedType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InviteStatusRescindedType: %v", v)
	}
}

type InviteStatusRescinded struct {
	// System generated unique identifier for a user. Not guaranteed to have a specific format.
	RescindedBy string `json:"rescindedBy"`
	// System generated unique identifier for a user. Not guaranteed to have a specific format.
	UserID *string                   `json:"userId,omitempty"`
	Type   InviteStatusRescindedType `json:"type"`
}

func (o *InviteStatusRescinded) GetRescindedBy() string {
	if o == nil {
		return ""
	}
	return o.RescindedBy
}

func (o *InviteStatusRescinded) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *InviteStatusRescinded) GetType() InviteStatusRescindedType {
	if o == nil {
		return InviteStatusRescindedType("")
	}
	return o.Type
}
