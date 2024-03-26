// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"errors"
	"github.com/livepeer/livepeer-go/internal/utils"
)

type Pinata2 struct {
	// Will be added to the pinata_api_key header.
	APIKey string `json:"apiKey"`
}

func (o *Pinata2) GetAPIKey() string {
	if o == nil {
		return ""
	}
	return o.APIKey
}

type Pinata1 struct {
}

type PinataType string

const (
	PinataTypePinata1 PinataType = "pinata_1"
	PinataTypePinata2 PinataType = "pinata_2"
)

// Pinata - Custom credentials for the Piñata service. Must have either
// a JWT or an API key and an API secret.
type Pinata struct {
	Pinata1 *Pinata1
	Pinata2 *Pinata2

	Type PinataType
}

func CreatePinataPinata1(pinata1 Pinata1) Pinata {
	typ := PinataTypePinata1

	return Pinata{
		Pinata1: &pinata1,
		Type:    typ,
	}
}

func CreatePinataPinata2(pinata2 Pinata2) Pinata {
	typ := PinataTypePinata2

	return Pinata{
		Pinata2: &pinata2,
		Type:    typ,
	}
}

func (u *Pinata) UnmarshalJSON(data []byte) error {

	pinata1 := Pinata1{}
	if err := utils.UnmarshalJSON(data, &pinata1, "", true, true); err == nil {
		u.Pinata1 = &pinata1
		u.Type = PinataTypePinata1
		return nil
	}

	pinata2 := Pinata2{}
	if err := utils.UnmarshalJSON(data, &pinata2, "", true, true); err == nil {
		u.Pinata2 = &pinata2
		u.Type = PinataTypePinata2
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Pinata) MarshalJSON() ([]byte, error) {
	if u.Pinata1 != nil {
		return utils.MarshalJSON(u.Pinata1, "", true)
	}

	if u.Pinata2 != nil {
		return utils.MarshalJSON(u.Pinata2, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type IpfsExportParams struct {
	DollarRef interface{} `json:"$ref,omitempty"`
	// Custom credentials for the Piñata service. Must have either
	// a JWT or an API key and an API secret.
	//
	Pinata *Pinata `json:"pinata,omitempty"`
}

func (o *IpfsExportParams) GetDollarRef() interface{} {
	if o == nil {
		return nil
	}
	return o.DollarRef
}

func (o *IpfsExportParams) GetPinata() *Pinata {
	if o == nil {
		return nil
	}
	return o.Pinata
}
