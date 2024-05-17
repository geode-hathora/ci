// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CardPaymentMethod struct {
	Last4 string    `json:"last4"`
	Brand CardBrand `json:"brand"`
}

func (o *CardPaymentMethod) GetLast4() string {
	if o == nil {
		return ""
	}
	return o.Last4
}

func (o *CardPaymentMethod) GetBrand() CardBrand {
	if o == nil {
		return CardBrand("")
	}
	return o.Brand
}
