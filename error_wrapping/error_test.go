package error_wrapping

import (
	"errors"
	"fmt"
    "io/fs"
	"os"
	"testing"
)

// https://pkg.go.dev/errors#Join

func TestErrors_join(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)
	fmt.Println(err)
	if errors.Is(err, err1) {
		fmt.Println("err is err1")
	}
	if errors.Is(err, err2) {
		fmt.Println("err is err2")
	}
}

func TestErrors_unwrap(t *testing.T) {
	err1 := errors.New("apple")
	err2 := fmt.Errorf("banana: %w", err1)
	fmt.Println(err2) // banana: apple
	fmt.Println(errors.Unwrap(err2)) // apple
	fmt.Println(err2) // banana: apple
}

func TestErrors_is(t *testing.T) {
    _, err := os.Open("non-existing")
    if errors.Is(err, fs.ErrNotExist) {
        fmt.Println("file does not exist")
    } else {
        t.Error("wrong error type")
    }
}
