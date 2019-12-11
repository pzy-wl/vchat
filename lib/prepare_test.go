package lib

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
	"time"
	"vchat/common/g"
	"vchat/lib/ylog"
	"vchat/lib/ymongo"

	"github.com/davecgh/go-spew/spew"

	"vchat/lib/yetcd"
	"vchat/lib/ypg"
	"vchat/lib/yredis"
)

//load config from vchat/.

func Test_config_load_etcd(t *testing.T) {
	PrepareLibs(true, false, false, false, false)
	spew.Dump(yetcd.XETCDConfig)
	err := yetcd.RegisterService("api", "www.sina.com.cn", "3333")
	log.Println(err)
}

func Test_config_load_pg(t *testing.T) {
	err := PrepareLibs(false, true, false, false, false)
	if err != nil {
		log.Println(err)
		return
	}
	spew.Dump(yetcd.XETCDConfig)
	if ypg.XDB == nil {
		log.Println("xdb is null")
	}

	b := ypg.XDB.HasTable("t")
	fmt.Println("------exists-----------")
	log.Println("b", b)
	fmt.Println("------mapResult-----------")
	l := make([]interface{}, 0)
	if err = ypg.XDB.Raw("select * from t").First(&l).Error; err != nil {
		log.Println("select error:", err)
	}

	spew.Dump(l)
}

func Test_load_config_redis(t *testing.T) {
	err := PrepareLibs(
		false,
		false,
		true, //
		false,
		false)
	if err != nil {
		log.Println(err)
		return
	}
	//spew.Dump(yetcd.XETCDConfig)
	//
	ret, er := yredis.XRed.Set("key", "aaa_value", time.Hour).Result()
	log.Println(ret, er)
	fmt.Println("------after set-----------")
	ret, er = yredis.XRed.Get("key").Result()
	log.Println(ret)
	fmt.Println("-----------------")
	log.Println(err)
	//
	lua := `local a = redis.call('get','key') 
			return "this is a good result"  `
	//ret, err = yredis.XRed.ScriptLoad(lua).Result()
	result, err1 := yredis.XRed.Eval(lua, nil).Result()
	// it should a wrong test
	// not lua script

	fmt.Println("------lua result-----------")
	log.Println(result)
	log.Println(err1)
}

func Test_load_config_mongo(t *testing.T) {
	err := PrepareLibs(
		false,
		false,
		false, //
		false,
		true)
	if err != nil {
		log.Println(err)
		return
	}
	//

	var ctx = context.Background()
	var docs []interface{}

	client := ymongo.XMongo
	defer client.Disconnect(ctx)

	log.Println("cnt ok")

	//d
	dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)

	h := 10000
	t0 := time.Now()
	for i := 0; i < h; i++ {
		docs = append(docs, bson.M{"a": i, "b": i * 10})
		if err != nil {
			log.Println("insert One err:", err)
		}
	}

	if _, err := tb.InsertMany(ctx, docs); err != nil {
		log.Println("insert err:", err)
	}
	fmt.Println("-------------time----", time.Since(t0))

	fmt.Println("------count-----------")
	c, err := tb.CountDocuments(ctx, bson.M{})
	if err != nil {
		fmt.Println("count err:", err)
		return
	}
	fmt.Println("------count:----", c)
	//--------result -----------------------------
}

func Test_load_config_log(t *testing.T) {
	err := PrepareLibs(
		false,
		false,
		false, //
		false,
		true)
	if err != nil {
		log.Println(err)
		return
	}
	//-------- -----------------------------
	h := 100000
	wg := g.NewWaitGroupN(400)
	for i := 0; i < h; i++ {
		wg.Call(func() error {
			ylog.Info(i, "---- ok")
			ylog.InfoF("---- ok %d", i)
			ylog.InfoDump(i)

			ylog.Debug(i, "---- ok")
			ylog.DebugF("---- ok %d", i)
			ylog.DebugDump(i)

			ylog.Warn(i, "---- ok")
			ylog.WarnF("---- ok %d", i)
			ylog.WarnDump("---- ok ")

			ylog.Error(i, "---- ok")
			ylog.ErrorF("---- ok %d", i)
			ylog.ErrorDump(i)
			return err
		})
	}
	fmt.Println("----time--- ", wg.Wait(), "  of", h)

}
