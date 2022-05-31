package safe

import (
	"testing"

	"github.com/go-zoox/errors"
)

type customError struct {
	msg string
}

func (s *customError) Error() string {
	return s.msg
}

type throwData struct {
	value string
}

func TestDo(t *testing.T) {
	var err error
	err = Do(func() error {
		return nil
	})
	if err != nil {
		t.Errorf("Expected nil, got %#v", err)
	}

	err = Do(func() error {
		return &customError{"custom error"}
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "custom error" {
		t.Errorf("Expected 'custom error', got %#v", err)
	}

	err = Do(func() error {
		panic("throw")
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "throw" {
		t.Errorf("Expected 'throw', got %#v", err)
	}

	err = Do(func() error {
		panic(errors.New("throw"))
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "throw" {
		t.Errorf("Expected 'throw', got %#v", err)
	}

	err = Do(func() error {
		panic(&throwData{"throw"})
	})
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestFunc(t *testing.T) {
	fn := Func(func() error {
		return nil
	})
	err := fn()
	if err != nil {
		t.Errorf("Expected nil, got %#v", err)
	}

	fn = Func(func() error {
		return &customError{"custom error"}
	})
	err = fn()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "custom error" {
		t.Errorf("Expected 'custom error', got %#v", err)
	}

	fn = Func(func() error {
		panic("throw")
	})
	err = fn()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "throw" {
		t.Errorf("Expected 'throw', got %#v", err)
	}

	fn = Func(func() error {
		panic(errors.New("throw"))
	})
	err = fn()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if err.Error() != "throw" {
		t.Errorf("Expected 'throw', got %#v", err)
	}

	fn = Func(func() error {
		panic(&throwData{"throw"})
	})
	err = fn()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
