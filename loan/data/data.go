package data

import (
	"net/url"

	"github.com/thedevsaddam/govalidator"
)

const (
	//PostDataType is the type of request body data in a POST request
	PostDataType int = iota
	//PutDataType is the type of request body datd in a PUT request
	PutDataType
)

//LoanData holds request body data when requesting a loan
type LoanData struct {
	LoanID     string  `json:"loanId"`
	BorrowerID string  `json:"borrowerId"`
	Amount     int64   `json:"amount"`
	Term       int     `json:"term"`
	Interest   float64 `json:"interest"`
}

//Validate validates LoanData
func (data *LoanData) Validate(dataType int) url.Values {
	rules := govalidator.MapData{}

	if dataType == PostDataType {
		rules = govalidator.MapData{
			"borrowerId": []string{"required", "uuid"},
			"amount":     []string{"required", "numeric"},
			"term":       []string{"required", "numeric"},
			"interest":   []string{"required"},
		}
	} else { //Update request
		rules = govalidator.MapData{
			"loanId":   []string{"required", "uuid"},
			"amount":   []string{"numeric"},
			"term":     []string{"numeric"},
			"interest": []string{},
		}
	}

	opts := govalidator.Options{
		Data:  data,
		Rules: rules,
	}

	v := govalidator.New(opts)

	return v.ValidateStruct()
}
