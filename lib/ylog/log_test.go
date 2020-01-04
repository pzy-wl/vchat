package ylog

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/weihaoranW/vchat/common/g"
	"github.com/weihaoranW/vchat/common/ytime"
)

func Test_log_test(t *testing.T) {
	obj := &LogWorker{
		today:      ytime.Today(),
		BackupPath: "./log/backup",
		LogPath:    "./log",
		FileName:   "vchat",
		Ext:        ".log",
	}

	h := 100000
	bean := g.NewWaitGroupN(100)

	for i := 100; i < 100+h; i++ {
		j := i
		bean.Call(func() error {
			obj.Println(j, j)
			return nil
		})
	}
	bean.Wait()

}

func Test_moveTo_(t *testing.T) {
	src := "./log/"
	dst := "./log/backup"
	//
	err := MoveTo("a.txt", src, dst)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}

}

func Test_log_single(t *testing.T) {
	obj := &LogWorker{
		today:      ytime.Today(),
		BackupPath: "./log/backup",
		LogPath:    "./log",
		FileName:   "vchat",
		Ext:        ".log",
	}

	h := 10000 * 100
	t0 := time.Now()
	defer obj.Close()
	for i := 100; i < 100+h; i++ {
		j := i
		obj.Println("hello", j)
	}
	log.Println("time:", time.Since(t0))
}
