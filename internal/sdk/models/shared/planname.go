// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// PlanName - A plan defines how much CPU and memory is required to run an instance of your game server.
//
// `tiny`: shared core, 1gb memory
//
// `small`: 1 core, 2gb memory
//
// `medium`: 2 core, 4gb memory
//
// `large`: 4 core, 8gb memory
type PlanName string

const (
	PlanNameTiny   PlanName = "tiny"
	PlanNameSmall  PlanName = "small"
	PlanNameMedium PlanName = "medium"
	PlanNameLarge  PlanName = "large"
)

func (e PlanName) ToPointer() *PlanName {
	return &e
}
func (e *PlanName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tiny":
		fallthrough
	case "small":
		fallthrough
	case "medium":
		fallthrough
	case "large":
		*e = PlanName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PlanName: %v", v)
	}
}
