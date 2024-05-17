// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CardBrand string

const (
	CardBrandAmex            CardBrand = "amex"
	CardBrandCartesBancaires CardBrand = "cartes_bancaires"
	CardBrandDiners          CardBrand = "diners"
	CardBrandDiscover        CardBrand = "discover"
	CardBrandJcb             CardBrand = "jcb"
	CardBrandMastercard      CardBrand = "mastercard"
	CardBrandVisa            CardBrand = "visa"
	CardBrandUnionpay        CardBrand = "unionpay"
	CardBrandCard            CardBrand = "card"
)

func (e CardBrand) ToPointer() *CardBrand {
	return &e
}
func (e *CardBrand) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amex":
		fallthrough
	case "cartes_bancaires":
		fallthrough
	case "diners":
		fallthrough
	case "discover":
		fallthrough
	case "jcb":
		fallthrough
	case "mastercard":
		fallthrough
	case "visa":
		fallthrough
	case "unionpay":
		fallthrough
	case "card":
		*e = CardBrand(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CardBrand: %v", v)
	}
}
