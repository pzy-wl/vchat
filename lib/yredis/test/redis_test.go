package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/vhaoran/vchat/lib/ylog"

	"github.com/davecgh/go-spew/spew"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yredis"
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
		//panic(" not init ok")
	}
}

func Test_call_back_set(t *testing.T) {
	for i := 0; i < 100; i++ {
		t0 := time.Now()
		k := i % 10
		v, err := yredis.CacheAutoGetH(new(Good), int64(k),
			func() (interface{}, error) {
				time.Sleep(50 * time.Millisecond)

				return &Good{
					ID:   int64(k),
					Name: "whr-test" + fmt.Sprint(int64(k)),
				}, nil

			})
		fmt.Println("------", err, "-----------")
		spew.Dump(v)
		fmt.Println("------", time.Since(t0), "-----------")
	}
}

func Test_CacheClearH(t *testing.T) {
	yredis.CacheClearH(new(Good), 2, 4, 8)
}

func Test_debug(t *testing.T) {
	ylog.Debug("hello")
}
func Test_say(t *testing.T)  {
	fmt.Println("hello, world!")
}