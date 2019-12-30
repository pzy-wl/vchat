package model

import (
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"vchat/lib"
	"vchat/lib/ypg"
)

type Abc struct {
	ypg.BaseModel
	ID    int
	CName string
}

func (r *Abc) TableName() string {
	return "abc"
}

func Test_pg_insert(t *testing.T) {
	// load config
	opt := &lib.LoadOption{
		LoadMicroService: false,
		LoadEtcd:         false,

		//-----------attention here------------
		LoadPg: true,
		//-----------attention here------------

		LoadRedis: false,

		LoadMongo: false,

		LoadMq:  false,
		LoadJwt: false,
	}
	_, err := lib.InitModulesOfOptions(opt)
	if err != nil {
		log.Println(err)
		return
	}

	if ypg.XDB.HasTable(new(Abc)) {
		err := ypg.XDB.DropTable(new(Abc)).Error
		if err != nil {
			log.Println(err)
			return
		}
	}

	if !ypg.XDB.HasTable(new(Abc)) {
		er := ypg.XDB.CreateTable(new(Abc)).Error
		if er != nil {
			fmt.Println("---create table err---", err, "-----------")
			return
		}
	}

	//
	for i := 0; i < 10; i++ {
		bean := &Abc{
			ID: i,
			CName: "whr_test" +
				"",
		}

		err = ypg.XDB.Save(bean).Error
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("------", "ok", "-----------")
	}

	fmt.Println("------", "demo find", "-----------")
	l := make([]*Abc, 0)
	err = ypg.XDB.Find(&l).Error
	if err != nil {
		log.Println(err)
		return
	}

	spew.Dump(l)
}
