package safe

import (
	"fmt"

	"github.com/go-zoox/errors"
)

// Do solves panic automatically, converts it to an error and returns it.
func Do(fn func() error) (err error) {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case string:
				err = errors.New(v)
			case error:
				err = v
			default:
				err = fmt.Errorf("%#v", v)
			}
		}
	}()

	err = fn()
	return err
}

// Func returns the method to be executed safely.
func Func(fn func() error) func() error {
	return func() error {
		return Do(fn)
	}
}
