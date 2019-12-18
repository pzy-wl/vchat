package testcase

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"

	"vchat/lib"
	"vchat/lib/ymongo"
)

type ABC struct {
	ID int `json:"id omitempty"`

	Name    string `json:"name omitempty"`
	Age     int    `json:"age omitempty"`
	AgeIsOk int    `json:"test_b omitempty"`
}

var (
	db *ymongo.MongoClientWrapper
)

func init() {
	err := lib.LoadModules(
		false,
		false,
		false,
		false,
		true,
	)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	db = &ymongo.MongoClientWrapper{
		Base: ymongo.XMongo,
	}
	if db == nil {
		log.Println("----------", "init", "------------")
		panic("db not prepare")
	} else {
		log.Println("----------", "db cnt passed", "------------")
	}
}

func Test_wrapper_insert_one(t *testing.T) {

	bean := &ABC{
		ID:   5,
		Name: "whr",
		Age:  3,
	}

	ret, err := db.DoInsertOne("test", "abc", bean)
	log.Println("--------", "", "--------------")
	log.Println(ret)
	log.Println("----------------------")
	log.Println(err)
}

func Test_wrapper_insert_many(t *testing.T) {
	l := make([]*ABC, 0)
	for i := 0; i < 1234; i++ {
		bean := &ABC{
			ID:   i,
			Name: fmt.Sprint("whr_aaaa_", i),
			Age:  i * 10,
		}
		l = append(l, bean)
	}

	ret, err := db.DoInsertMany("test", "abc", l)
	log.Println("--------", "", "--------------")
	log.Println(ret)
	log.Println("----------------------")
	log.Println(err)
	//

}

func Test_mongo_client_wrapper_find(t *testing.T) {
	client := db.Base

	ctx := context.Background()

	dbName, tbName := "test", "abc"
	tb := client.Database(dbName).Collection(tbName)
	c, err := tb.Find(ctx, bson.M{"id": 1})
	log.Println("----------", err, "------------")
	for c.Next(ctx) {
		log.Println(c.Current)
		bean := &ABC{}
		if er1 := bson.Unmarshal(c.Current, &bean); er1 != nil {
			log.Println(err)
		} else {
			log.Println("----------", "unmarshal ok", "------------")
			spew.Dump(bean)
		}
	}

}

func Test_test_wrapper_find(t *testing.T) {
	l := make([]*ABC, 0)
	err := db.DoFind(&l, "test", "abc", bson.M{"id": 1})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	spew.Dump(l)
}
