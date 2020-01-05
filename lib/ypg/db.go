package ypg

import (
	"log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/weihaoranW/vchat/lib/yconfig"
)

var (
	XDB *gorm.DB
)

/*--auth: whr  date:2019/12/511:44--------------------------
 ####请勿擅改此功能代码####
 用途：连接到数据库功能
 --->：yconfig
--------------------------------------- */

func InitPG(cfg yconfig.PGConfig) (err error) {
	if cnt, err := NewPGCnt(&cfg); err != nil {
		return err
	} else {
		cnt.Callback().Create().Replace("gorm:update_time_stamp", createCallback)
		cnt.Callback().Update().Replace("gorm:update_time_stamp", updateCallback)
		//cnt.Callback().Delete().Replace("gorm:delete", deleteCallback)
		XDB = cnt
		return nil
	}
}

func NewPGCnt(cfg *yconfig.PGConfig) (*gorm.DB, error) {
	connStr := cfg.URL

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	db.DB().SetMaxOpenConns(cfg.PoolMax)
	db.DB().SetMaxIdleConns(cfg.PoolMax)

	return db, nil
}
