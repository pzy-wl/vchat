package test

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/jinzhu/gorm"
	"github.com/weihaoranW/vchat/lib"
	"github.com/weihaoranW/vchat/lib/ypg"
	"log"
	"testing"
)

type GoodA struct {
	ID   int64
	Name string
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
