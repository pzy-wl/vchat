package yverify

import (
	"errors"
	"fmt"
	"testing"
)

type XX struct {
}

func (r *XX) Exec() string {
	return "no no no"
}

func Test_aaa(t *testing.T) {
	if err := NewObj(false).
		Gt(3, 20, "fd1").
		Lt(40, 5, "fd2").
		Fn(errors.New("err2")).
		FnBool("我的姓别", false).
		Err()
		err != nil {
		fmt.Println("ret: ", err.Error())
	}

}
