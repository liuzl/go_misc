package main

import (
	"fmt"

	"github.com/juju/errors"
)

func main() {
	err := errors.ErrorStack(fmt.Errorf("this is cool"))
	fmt.Println(err)
}
