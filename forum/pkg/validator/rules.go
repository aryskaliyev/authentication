package validator

import (
	"fmt"
	"strings"
)

func Required(key string, value interface{}) error {
	if strings.TrimSpace(value.(string)) == "" {
		return ErrMissingField
	}

	return nil
}

func MaxLength(maxLength int) Rule {
	return func(key string, value interface{}) error {
		str, ok := value.(string)
		if !ok {
			return fmt.Errorf("%s is not a string", key)
		}

		if len(str) > maxLength {
			return ErrStringLengthMax
		}

		return nil
	}
}

func MinLength(minLength int) Rule {
	return func(key string, value interface{}) error {
		str, ok := value.(string)

		if !ok {
			return fmt.Errorf("%s is not a string", key)
		}

		if len(str) < minLength {
			return ErrStringLengthMin
		}

		return nil
	}
}
