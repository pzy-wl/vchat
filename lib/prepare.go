package lib

import (
	"github.com/weihaoranW/vchat/common/ytime"
	_ "github.com/weihaoranW/vchat/common/ytime"
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

func InitModulesOfAll() (*yconfig.YmlConfig, error) {
	cfg := LoadOption{
		LoadMicroService: true,
		LoadEtcd:         true,
		LoadPg:           true,
		LoadRedis:        true,
		LoadMongo:        true,
		LoadMq:           true,
		LoadJwt:          true,
	}

	return InitModulesOfOptions(&cfg)
}

func InitModulesOfOptions(opt *LoadOption) (*yconfig.YmlConfig, error) {
	var (
		cfg *yconfig.YmlConfig
		err error
	)
	//initOthers()

	if cfg, err = yconfig.GetYmlConfig(); err != nil {
		return nil, err
	}
	ylog.Debug("----------", "config-file", "------------")
	//spew.Dump(cfg)
	//log.Println("----------", "----", "------------")

	if err = ylog.InitLog(cfg.Log); err != nil {
		ylog.DebugDump(cfg.Log)
		return nil, err
	}
	//--------etcd -----------------------------
	// 微服务注册地址设置set XEtcdConfig
	if opt.LoadEtcd {
		ylog.Debug("etcd connecting...", cfg.Etcd.Hosts)
		if err := yetcd.InitETCD(cfg.Etcd); err != nil {
			return nil, err
		}
		ylog.Debug("etcd connected ok")
	}

	//-------- postgres sql -----------------------------
	//postgres 数据库配置参数设置 X
	if opt.LoadPg {
		ylog.Debug("postgres connecting...", cfg.Postgres.URL)
		if err := ypg.InitPG(cfg.Postgres); err != nil {
			return nil, err
		}
		ylog.Debug("postgres connected ok")
	}

	//--------load redis -----------------------------
	//redis cluster连接设置 xred
	if opt.LoadRedis {
		//set X
		ylog.Debug("redis connecting...", cfg.Redis.Addrs)
		if err := yredis.InitRedis(cfg.Redis); err != nil {
			return nil, err
		}
		ylog.Debug("redis connected ok")
	}

	//--------load emq -----------------------------

	//--------load mongo -----------------------------
	if opt.LoadMongo {
		ylog.Debug("mongo connecting...", cfg.Mongo.URL)
		if err := ymongo.InitMongo(cfg.Mongo); err != nil {
			return nil, err
		}
		ylog.Debug("mongo connected ok")
	}

	if opt.LoadJwt {
		if err := yjwt.InitJwt(cfg.Jwt); err != nil {
			return nil, err
		}
	}

	if opt.LoadMq {
		ylog.Debug("emqx connecting...", cfg.Emq.Url, "  ;  ", cfg.Emq.Host)
		if err := ymq.InitMq(cfg.Emq); err != nil {
			return nil, err
		}
		ylog.Debug("emqx connected ok")
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

//初始化其它数据
func initOthers() {
	ytime.SetTimeZone()
}
