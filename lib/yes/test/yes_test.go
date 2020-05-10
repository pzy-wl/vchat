package test

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"testing"

	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/yes"
	"github.com/vhaoran/vchat/lib/ylog"
)

type Product struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	CateID   int64  `json:"cate_id,omitempty"`
	CateName string `json:"cate_name,omitempty"`
	Tag      string `json:"tag,omitempty"`

	Price  float32 `json:"price,omitempty"`
	Remark string  `json:"remark,omitempty"`
}

func init() {
	_, err := lib.InitModulesOfOptions(
		&lib.LoadOption{
			LoadMicroService: false,
			LoadEtcd:         false,
			LoadPg:           false,
			LoadRedis:        false,
			LoadMongo:        false,
			LoadMq:           false,
			LoadRabbitMq:     false,
			LoadJwt:          false,
			LoadES:           true,
		})
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(" ---------- init ok---------")
}

func Test_add(t *testing.T) {
	for i := int64(0); i < 100; i++ {
		bean := Product{
			//ID:       i,
			Name:     fmt.Sprint("name_", i),
			CateID:   0,
			CateName: fmt.Sprint("cate_", 1),
			Tag:      "汽车 飞机 大炮",
			Price:    float32(i * 10.0),
			Remark:   "this is a good test",
		}
		r, err := yes.X.Index().Index("index").BodyJson(bean).Do(context.Background())
		ylog.Debug("--------yes_test.go------", err)
		ylog.DebugDump("--------yes_test.go------", r)
	}
}

func Test_match(t *testing.T) {
	q := elastic.NewTermQuery("tag", "飞机")

	r, err := yes.X.Search("index").
		Query(q).
		Do(context.Background())
	ylog.Debug("--------yes_test.go------", err)
	ylog.DebugDump("--------yes_test.go------", r)
}
//
func Test_match2(t *testing.T) {
	//q := elastic.NewTermQuery("tag", "飞机")
	//q := elastic.NewQueryStringQuery()
	//q := elastic.NewMatchAllQuery()
	q := elastic.NewMatchQuery()


	//q := elastic.NewScriptQuery("")
	//q := elastic.NewWrapperQuery()

	r, err := yes.X.Search("index").Query().
		Do(context.Background())
	ylog.Debug("--------yes_test.go------", err)
	ylog.DebugDump("--------yes_test.go------", r)
}
