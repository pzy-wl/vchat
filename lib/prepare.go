package lib

import (
	"log"

	"github.com/davecgh/go-spew/spew"

	"github.com/weihaoranW/vchat/lib/yconfig"
	"github.com/weihaoranW/vchat/lib/yetcd"
	"github.com/weihaoranW/vchat/lib/yjwt"
	"github.com/weihaoranW/vchat/lib/ylog"
	"github.com/weihaoranW/vchat/lib/ymongo"
	"github.com/weihaoranW/vchat/lib/ymq"
	"github.com/weihaoranW/vchat/lib/ypg"
	"github.com/weihaoranW/vchat/lib/yredis"
)

type LoadOption struct {
	LoadMicroService bool //0
	LoadEtcd         bool //1
	LoadPg           bool //2
	LoadRedis        bool //3
	LoadMongo        bool //4
	LoadMq           bool //5
	LoadJwt          bool //6
}

func InitModulesOfOptions(opt *LoadOption) (*yconfig.YmlConfig, error) {
	var (
		cfg *yconfig.YmlConfig
		err error
	)
	if cfg, err = yconfig.GetYmlConfig(); err != nil {
		return nil, err
	}
	log.Println("----------", "config-file", "------------")
	//spew.Dump(cfg)
	//log.Println("----------", "----", "------------")

	if err = ylog.InitLog(cfg.Log); err != nil {
		spew.Dump(cfg.Log)
		return nil, err
	}
	//--------etcd -----------------------------
	// 微服务注册地址设置set XEtcdConfig
	if opt.LoadEtcd {
		if err := yetcd.InitETCD(cfg.Etcd); err != nil {
			return nil, err
		}
	}

	//-------- postgres sql -----------------------------
	//postgres 数据库配置参数设置 XDB
	if opt.LoadPg {
		if err := ypg.InitPG(cfg.Postgres); err != nil {
			return nil, err
		}
	}

	//--------load redis -----------------------------
	//redis cluster连接设置 xred
	if opt.LoadRedis {
		//set XRed
		if err := yredis.InitRedis(cfg.Redis); err != nil {
			return nil, err
		}
	}

	//--------load emq -----------------------------

	//--------load mongo -----------------------------
	if opt.LoadMongo {
		if err := ymongo.InitMongo(cfg.Mongo); err != nil {
			return nil, err
		}
	}

	if opt.LoadJwt {
		if err := yjwt.InitJwt(cfg.Jwt); err != nil {
			return nil, err
		}
	}

	if opt.LoadMq {
		if err := ymq.InitMq(cfg.Emq); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}

func InitModules(loadEtcd, loadPostgres, loadRedis, loadEMQ, loadMongo bool) error {
	opt := &LoadOption{
		LoadEtcd:  loadEtcd,
		LoadPg:    loadPostgres,
		LoadRedis: loadRedis,
		LoadMongo: loadMongo,
		LoadMq:    loadEMQ,
		LoadJwt:   true,
	}

	_, err := InitModulesOfOptions(opt)
	return err
}
