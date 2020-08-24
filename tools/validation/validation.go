package validation

import (
	"errors"
	"fmt"
	"reflect"
)

func ValidateInputNotEmpty(data ...interface{}) error {
	for _, value := range data {
		switch value {
		case "":
			return errors.New("make sure input not empty")
		case 0:
			return errors.New("make sure input not a zero")
		case nil:
			return errors.New("make sure input not a nil")
		}
	}
	return nil
}

func ValidateUUID(data ...interface{}) error {
	for _, value := range data {
		if len(value.(string)) != 36 {
			return fmt.Errorf("value '%v' is not a valid UUID", value)
		}
	}
	return nil
}

func ValidateStructNotEmpty(x interface{}) error {
	v := reflect.ValueOf(x)
	values := make([]interface{}, v.NumField())
	for i, val := range values {
		val = v.Field(i).Interface()
		switch {
		case val == "" || val == 0 || val == nil:
			return errors.New("struct fields must be filled")
		default:
			continue
		}
	}
	return nil

}
