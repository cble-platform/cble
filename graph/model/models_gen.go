// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type BlueprintInput struct {
	Name              string `json:"name"`
	BlueprintTemplate string `json:"blueprintTemplate"`
	ParentGroupID     string `json:"parentGroupId"`
	ProviderID        string `json:"providerId"`
}

type ProviderInput struct {
	DisplayName     string `json:"displayName"`
	ProviderGitURL  string `json:"providerGitUrl"`
	ProviderVersion string `json:"providerVersion"`
	ConfigBytes     string `json:"configBytes"`
}

type UserInput struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CommandStatus string

const (
	CommandStatusQueued     CommandStatus = "QUEUED"
	CommandStatusFailed     CommandStatus = "FAILED"
	CommandStatusSucceeded  CommandStatus = "SUCCEEDED"
	CommandStatusInprogress CommandStatus = "INPROGRESS"
)

var AllCommandStatus = []CommandStatus{
	CommandStatusQueued,
	CommandStatusFailed,
	CommandStatusSucceeded,
	CommandStatusInprogress,
}

func (e CommandStatus) IsValid() bool {
	switch e {
	case CommandStatusQueued, CommandStatusFailed, CommandStatusSucceeded, CommandStatusInprogress:
		return true
	}
	return false
}

func (e CommandStatus) String() string {
	return string(e)
}

func (e *CommandStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommandStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommandStatus", str)
	}
	return nil
}

func (e CommandStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CommandType string

const (
	CommandTypeConfigure CommandType = "CONFIGURE"
	CommandTypeDeploy    CommandType = "DEPLOY"
	CommandTypeDestroy   CommandType = "DESTROY"
)

var AllCommandType = []CommandType{
	CommandTypeConfigure,
	CommandTypeDeploy,
	CommandTypeDestroy,
}

func (e CommandType) IsValid() bool {
	switch e {
	case CommandTypeConfigure, CommandTypeDeploy, CommandTypeDestroy:
		return true
	}
	return false
}

func (e CommandType) String() string {
	return string(e)
}

func (e *CommandType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CommandType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CommandType", str)
	}
	return nil
}

func (e CommandType) MarshalGQL(w io.Writer) {
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
