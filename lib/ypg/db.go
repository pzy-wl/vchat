package ypg

import (
	"log"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"vchat/lib/yconfig"
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

		cnt.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
		cnt.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
		cnt.Callback().Delete().Replace("gorm:delete", deleteCallback)

		XDB = cnt
		return nil
	}
}

func NewPGCnt(cfg *yconfig.PGConfig) (*gorm.DB, error) {
	//connStr := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
	//	"127.0.0.1", //viper.GetString("DB_HOST"),
	//	"root",      //viper.GetString("DB_USER"),
	//	"test",      //viper.GetString("DB_NAME"),
	//	"password",  ///viper.GetString("DB_PASS"),
	//)
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
