package lib

import (
	"github.com/davecgh/go-spew/spew"
	"log"
	"vchat/lib/yconfig"
	"vchat/lib/yetcd"
	"vchat/lib/ylog"
	"vchat/lib/ymongo"
	"vchat/lib/ypg"
	"vchat/lib/yredis"
)

func PrepareLibs(loadEtcd, loadPostgres, loadRedis, loadEMQ, loadMongo bool) error {
	var (
		cfg *yconfig.YmlConfig
		err error
	)
	if cfg, err = yconfig.GetYmlConfig(); err != nil {
		return err
	}

	log.Println("----------", "config-file", "------------")
	spew.Dump(cfg)
	log.Println("----------", "----", "------------")

	if err = ylog.InitLog(cfg.Log); err != nil {
		spew.Dump(cfg.Log)
		return err
	}

	//--------etcd -----------------------------
	// 微服务注册地址设置set XEtcdConfig
	if loadEtcd {
		if err := yetcd.InitETCD(cfg.Etcd); err != nil {
			return err
		}
	}

	//-------- postgres sql -----------------------------
	//postgres 数据库配置参数设置 XDB
	if loadPostgres {
		if err := ypg.InitPG(cfg.Postgres); err != nil {
			return err
		}
	}

	//--------load redis -----------------------------
	//redis cluster连接设置 xred
	if loadRedis {
		//set XRed
		if err := yredis.InitRedis(cfg.Redis); err != nil {
			return err
		}
	}

	//--------load emq -----------------------------

	//--------load mongo -----------------------------
	if loadMongo {
		if err := ymongo.InitMongo(cfg.Mongo); err != nil {
			return err
		}
	}

	return nil
}
