package testme

import (
	"encoding/json"
	"errors"
	"fmt"
)

// this is a sentinal error that comes from the standard [errors] package
var Sentinal = errors.New("ok")

// this function uses [encoding/json] from the standard library
func SomeJson() error {
	j, err := json.Marshal(5)

	fmt.Printf("%s\n\n", j)

	return fmt.Errorf("then error is: %w", err)

}
