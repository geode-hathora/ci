// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type AchPaymentMethod struct {
	Last4    *string `json:"last4,omitempty"`
	BankName *string `json:"bankName,omitempty"`
}

func (o *AchPaymentMethod) GetLast4() *string {
	if o == nil {
		return nil
	}
	return o.Last4
}

func (o *AchPaymentMethod) GetBankName() *string {
	if o == nil {
		return nil
	}
	return o.BankName
}
