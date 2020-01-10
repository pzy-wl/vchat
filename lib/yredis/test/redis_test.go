package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/weihaoranW/vchat/lib"
	"github.com/weihaoranW/vchat/lib/yredis"
)

type Good struct {
	ID   int64
	Name string
}

func (Good) TableName() string {
	return "good"
}

func init() {
	_, err := lib.InitModulesOfAll()
	if err != nil {
		panic(" not init ok")
	}
}

func Test_call_back_set(t *testing.T) {
	for i := 0; i < 10; i++ {
		t0 := time.Now()
		v, err := yredis.CacheAutoGetH(new(Good), int64(1),
			func(fd interface{}) (interface{}, error) {
				time.Sleep(5 * time.Second)

				return &Good{
					ID:   fd.(int64),
					Name: "whr-test" + fmt.Sprint(fd),
				}, nil

			})
		fmt.Println("------", err, "-----------")
		spew.Dump(v)
		fmt.Println("------", time.Since(t0), "-----------")
	}

}
