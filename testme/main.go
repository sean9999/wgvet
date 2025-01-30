package testme

import (
	"encoding/json"
	"errors"
	"fmt"
)

var Sentinal = errors.New("ok")

func SomeJson() error {
	j, err := json.Marshal(5)

	fmt.Printf("%s\n\n", j)

	return fmt.Errorf("then error is: %w", err)

}
