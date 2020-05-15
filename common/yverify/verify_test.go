package yverify

import (
	"errors"
	"fmt"
	"testing"
)

func Test_aaa(t *testing.T) {
	if err := NewObj(false).
		Gt("fd1", 3, 20).
		Lt("fd2", 40, 5).
		Fn(errors.New("err2")).
		Err()
		err != nil {
		fmt.Println("ret: ", err.Error())
	}
}
