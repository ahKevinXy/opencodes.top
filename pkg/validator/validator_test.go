package validators

import (
	"fmt"
	"reflect"
	"testing"
)

type User struct {
	FirstName string `validate:"required" json:"first_name"`
	Name      string `validate:"required" json:"name"`
	Key       string `validate:"required" json:"key"`
}

func TestInitValidator(t *testing.T) {
	InitValidator()

	user := User{
		FirstName: "1111",
	}

	err := Validate.Struct(user)

	errs := GetErrorMessage(err)

	a := *errs

	fmt.Println(reflect.TypeOf(a))
}
