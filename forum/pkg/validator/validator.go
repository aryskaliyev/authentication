package validator

import "fmt"

type Rule func(key string, value interface{}) error

type Validator struct {
	rules []Rule
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) AddRule(rule Rule) {
	v.rules = append(v.rules, rule)
}

func (v *Validator) Validate(data map[string]interface{}) map[string]string {
	errors := make(map[string]string)

	for _, rule := range v.rules {
		for key, value := range data {
			if err := rule(key, value); err != nil {
				errors[key] = fmt.Sprintf("%v", err)
			}
		}
	}

	return errors
}
