package test

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/vhaoran/vchat/common/ytime"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ypg"
	"log"
	"testing"
)

type GoodA struct {
	ID   int64
	Name string
	T    ytime.Date
	TM   ytime.Date
	B    ytime.DateM
}

func (GoodA) TableName() string {
	return "good_a"
}

func (r *GoodA) AfterSave(scope *gorm.Scope) (err error) {
	//bean := scope.Value.(*GoodA)
	log.Println("....after save...")
	spew.Dump(r)

	return nil
}

func init() {
	_, err := lib.InitModulesOfAll()
	if err != nil {
		panic(err.Error())
	}
	ypg.X.AutoMigrate(new(GoodA))
}

func Test_trigger(t *testing.T) {
	//
	bean := &GoodA{
		Name: "abc",
	}
	err := ypg.X.Save(bean).Error
	if err != nil {
		log.Println(err)
	}
}

func Test_trigger_2(t *testing.T) {
	err := ypg.X.Model(new(GoodA)).Where("id=?", 1).Update("name", "abcde").Error
	if err != nil {
		log.Println(err)
	}
}

func Test_insert(t *testing.T) {
	ypg.X.AutoMigrate(new(GoodA))
	ypg.X.LogMode(true)
	ytime.SetTimeZone()

	for i := 3; i < 10; i++ {
		//c := ytime.OfNowM()
		bean := &GoodA{
			Name: fmt.Sprint(i, " name_"),
			T:    ytime.OfNow(),
			TM:   ytime.OfNow(),
			//B:    ytime.OfNowM(),
		}
		err := ypg.X.Save(bean).Error
		if err != nil {
			fmt.Println("pg_test.go->", err)
			return
		}
		fmt.Println(i, " ok")
	}

}
