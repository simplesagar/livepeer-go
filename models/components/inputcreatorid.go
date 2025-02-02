// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type InputCreatorIDType string

const (
	InputCreatorIDTypeUnverified InputCreatorIDType = "unverified"
)

func (e InputCreatorIDType) ToPointer() *InputCreatorIDType {
	return &e
}

func (e *InputCreatorIDType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "unverified":
		*e = InputCreatorIDType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for InputCreatorIDType: %v", v)
	}
}

type InputCreatorID1 struct {
	Type  InputCreatorIDType `json:"type"`
	Value string             `json:"value"`
}

func (o *InputCreatorID1) GetType() InputCreatorIDType {
	if o == nil {
		return InputCreatorIDType("")
	}
	return o.Type
}

func (o *InputCreatorID1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type InputCreatorIDUnionType string

const (
	InputCreatorIDUnionTypeInputCreatorID1 InputCreatorIDUnionType = "input-creator-id_1"
	InputCreatorIDUnionTypeStr             InputCreatorIDUnionType = "str"
)

type InputCreatorID struct {
	InputCreatorID1 *InputCreatorID1
	Str             *string

	Type InputCreatorIDUnionType
}

func CreateInputCreatorIDInputCreatorID1(inputCreatorID1 InputCreatorID1) InputCreatorID {
	typ := InputCreatorIDUnionTypeInputCreatorID1

	return InputCreatorID{
		InputCreatorID1: &inputCreatorID1,
		Type:            typ,
	}
}

func CreateInputCreatorIDStr(str string) InputCreatorID {
	typ := InputCreatorIDUnionTypeStr

	return InputCreatorID{
		Str:  &str,
		Type: typ,
	}
}

func (u *InputCreatorID) UnmarshalJSON(data []byte) error {

	inputCreatorID1 := InputCreatorID1{}
	if err := utils.UnmarshalJSON(data, &inputCreatorID1, "", true, true); err == nil {
		u.InputCreatorID1 = &inputCreatorID1
		u.Type = InputCreatorIDUnionTypeInputCreatorID1
		return nil
	}

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = InputCreatorIDUnionTypeStr
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u InputCreatorID) MarshalJSON() ([]byte, error) {
	if u.InputCreatorID1 != nil {
		return utils.MarshalJSON(u.InputCreatorID1, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
