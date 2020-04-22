package validation

// Validator validates a request input
type Validator struct {
	errorBag *ErrorBag
}

//Required checks if a required field exists
func (v *Validator) Required(field string, alias string, data *interface{}) {
	//if
}

//ErrorBag is a collection of validation errors
type ErrorBag struct {
	bag map[string][]string
}

func (errorBag *ErrorBag) init() {
	if errorBag.bag == nil {
		errorBag.bag = map[string][]string{}
	}
}

// Put adds an error item for a field
func (errorBag *ErrorBag) Put(field string, error string) {
	errorBag.init()
	errorBag.bag[field] = append(errorBag.bag[field], error)
}

// IsEmpty Checks if ErrorBag is empty
func (errorBag *ErrorBag) IsEmpty() bool {
	errorBag.init()
	return len(errorBag.bag) == 0
}

//AllErrors returns all errors
func (errorBag *ErrorBag) AllErrors() map[string][]string {
	errorBag.init()
	return errorBag.bag
}
