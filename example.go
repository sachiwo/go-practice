package main

import (
	"errors"
)

// Example is test function
func Example(code string) (int, error) {
	if code == "hoge" {
		return 1, nil
	}
	return 0, errors.New("code must be hoge")
}
