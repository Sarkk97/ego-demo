package data

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

//FundWalletData holds the request body of
//a transaction verification request
type FundWalletData struct {
	Reference string `json:"reference"`
	UserID    string `json:"userId"`
}

//Validate validates a FundWalletData
func (fundWalletData *FundWalletData) Validate() url.Values {
	rules := govalidator.MapData{
		"reference": []string{"required"},
		"userId":    []string{"required"},
	}

	opts := govalidator.Options{
		Data:  fundWalletData,
		Rules: rules,
	}

	v := govalidator.New(opts)

	return v.ValidateStruct()
}
