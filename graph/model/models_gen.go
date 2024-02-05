// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"

	"github.com/cble-platform/cble-backend/engine/models"
	"github.com/google/uuid"
)

type BlueprintInput struct {
	Name              string                                  `json:"name"`
	Description       string                                  `json:"description"`
	BlueprintTemplate string                                  `json:"blueprintTemplate"`
	VariableTypes     map[string]models.BlueprintVariableType `json:"variableTypes"`
	ProviderID        uuid.UUID                               `json:"providerId"`
}

type DeploymentInput struct {
	Name string `json:"name"`
}

type ProviderInput struct {
	DisplayName     string `json:"displayName"`
	ProviderGitURL  string `json:"providerGitUrl"`
	ProviderVersion string `json:"providerVersion"`
	ConfigBytes     string `json:"configBytes"`
}

type UserInput struct {
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	FirstName string      `json:"firstName"`
	LastName  string      `json:"lastName"`
	GroupIds  []uuid.UUID `json:"groupIds"`
}

type DeploymentNodeState string

const (
	DeploymentNodeStateAwaiting       DeploymentNodeState = "AWAITING"
	DeploymentNodeStateParentawaiting DeploymentNodeState = "PARENTAWAITING"
	DeploymentNodeStateInprogress     DeploymentNodeState = "INPROGRESS"
	DeploymentNodeStateComplete       DeploymentNodeState = "COMPLETE"
	DeploymentNodeStateTainted        DeploymentNodeState = "TAINTED"
	DeploymentNodeStateFailed         DeploymentNodeState = "FAILED"
	DeploymentNodeStateTodelete       DeploymentNodeState = "TODELETE"
	DeploymentNodeStateDeleted        DeploymentNodeState = "DELETED"
	DeploymentNodeStateTorebuild      DeploymentNodeState = "TOREBUILD"
)

var AllDeploymentNodeState = []DeploymentNodeState{
	DeploymentNodeStateAwaiting,
	DeploymentNodeStateParentawaiting,
	DeploymentNodeStateInprogress,
	DeploymentNodeStateComplete,
	DeploymentNodeStateTainted,
	DeploymentNodeStateFailed,
	DeploymentNodeStateTodelete,
	DeploymentNodeStateDeleted,
	DeploymentNodeStateTorebuild,
}

func (e DeploymentNodeState) IsValid() bool {
	switch e {
	case DeploymentNodeStateAwaiting, DeploymentNodeStateParentawaiting, DeploymentNodeStateInprogress, DeploymentNodeStateComplete, DeploymentNodeStateTainted, DeploymentNodeStateFailed, DeploymentNodeStateTodelete, DeploymentNodeStateDeleted, DeploymentNodeStateTorebuild:
		return true
	}
	return false
}

func (e DeploymentNodeState) String() string {
	return string(e)
}

func (e *DeploymentNodeState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentNodeState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentNodeState", str)
	}
	return nil
}

func (e DeploymentNodeState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type DeploymentState string

const (
	DeploymentStateAwaiting   DeploymentState = "AWAITING"
	DeploymentStateInprogress DeploymentState = "INPROGRESS"
	DeploymentStateComplete   DeploymentState = "COMPLETE"
	DeploymentStateFailed     DeploymentState = "FAILED"
	DeploymentStateDeleted    DeploymentState = "DELETED"
)

var AllDeploymentState = []DeploymentState{
	DeploymentStateAwaiting,
	DeploymentStateInprogress,
	DeploymentStateComplete,
	DeploymentStateFailed,
	DeploymentStateDeleted,
}

func (e DeploymentState) IsValid() bool {
	switch e {
	case DeploymentStateAwaiting, DeploymentStateInprogress, DeploymentStateComplete, DeploymentStateFailed, DeploymentStateDeleted:
		return true
	}
	return false
}

func (e DeploymentState) String() string {
	return string(e)
}

func (e *DeploymentState) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = DeploymentState(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid DeploymentState", str)
	}
	return nil
}

func (e DeploymentState) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PermissionPolicyType string

const (
	PermissionPolicyTypeAllow PermissionPolicyType = "ALLOW"
	PermissionPolicyTypeDeny  PermissionPolicyType = "DENY"
)

var AllPermissionPolicyType = []PermissionPolicyType{
	PermissionPolicyTypeAllow,
	PermissionPolicyTypeDeny,
}

func (e PermissionPolicyType) IsValid() bool {
	switch e {
	case PermissionPolicyTypeAllow, PermissionPolicyTypeDeny:
		return true
	}
	return false
}

func (e PermissionPolicyType) String() string {
	return string(e)
}

func (e *PermissionPolicyType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PermissionPolicyType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PermissionPolicyType", str)
	}
	return nil
}

func (e PermissionPolicyType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
