package errcode

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	var err = errors.New("hahah")
	err = NewError(DefaultServerError, err)
	if err != nil {
		fmt.Println(err)
	}
}
