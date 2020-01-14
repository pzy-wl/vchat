package testcase

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/vhaoran/vchat/common/ypage"
	"github.com/vhaoran/vchat/lib"
	"github.com/vhaoran/vchat/lib/ylog"
	"github.com/vhaoran/vchat/lib/ymongo"
)

type ABC struct {
	ID int `json:"id omitempty"`
	//
	Name    string `json:"name omitempty"`
	Age     int    `json:"age omitempty"`
	AgeIsOk int    `json:"test_b omitempty"`
}

var (
	db *ymongo.MongoClientWrapper
)

func init() {
	err := lib.InitModules(
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

	db = ymongo.X
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

	ret, err := db.DoInsertMany("test", "t", l)
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
	err := db.DoFindMany(&l, "test", "abc", bson.M{"_id": 1, "b": 10})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("------", "", "-----------")
	spew.Dump(l)
}

func Test_mongo_update_one(t *testing.T) {
	//var id primitive.ObjectID

	opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{bson.E{Key: "a", Value: 1}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "b", Value: 333}}}}
	var updatedDocument bson.M

	client := db.Base
	tb := client.Database("test").Collection("t")
	err := tb.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)

	if err != nil {
		log.Println("########## err:", err)
		return
	}

	fmt.Printf("updated document %v", updatedDocument)
}

func Test_mongo_update_many(t *testing.T) {
	//var id primitive.ObjectID

	//opts := options.FindOneAndUpdate().SetUpsert(true)
	filter := bson.D{{"a", 1}}
	update := bson.D{{"$set", bson.D{{"b", 100}}}}
	var updatedDocument bson.M

	client := db.Base
	tb := client.Database("test").Collection("abc")
	//err := tb.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedDocument)
	_, err := tb.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if err == mongo.ErrNoDocuments {
			log.Println("########## err:", err)
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("updated document %v", updatedDocument)
}

func Test_wrap_update_one(t *testing.T) {
	err := db.DoUpdateOne("test", "t",
		bson.D{{"a", 1}},
		bson.D{{"b", 32}},
	)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	fmt.Println("------", "ok", "-----------")
}

func Test_wrap_update_many(t *testing.T) {
	err := db.DoUpdateMany("test", "t",
		bson.D{{"a", 1}},
		bson.D{{"b", 42}})

	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	fmt.Println("------", "ok", "-----------")
}

func Test_del_many(t *testing.T) {
	i, err := db.DoDelMany("test", "t",
		bson.D{{"a", 1}})
	log.Println("----------", "aaa", "------------")
	ylog.Debug("i", i, "  err", err)
}

func Test_del_one(t *testing.T) {
	i, err := db.DoDelOne("test", "t",
		bson.D{{"a", ""}})
	log.Println("----------", "aaa", "------------")
	ylog.Debug("i", i, "  err", err)
}

type T struct {
	Key   string
	Value string
}

func Test_a_test(t *testing.T) {
	c := []T{{"a", "b"}, {"c", "d"}}
	for _, v := range c {
		fmt.Println(v)
	}
}

func Test_a_insert_One(t *testing.T) {
	bean := &ABC{
		ID:      1,
		Name:    "abc",
		Age:     0,
		AgeIsOk: 0,
	}
	_, err := db.DoInsertOne("test", "abc", bean)
	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
}

func Test_a_b_findOne(t *testing.T) {
	bean := &ABC{

	}

	err := db.DoFindOne(bean, "test", "abc",
		bson.M{"name": "abc"})
	if err != nil {
		ylog.Error("### mongoClientWrapper_test.go->", err)
		return
	}
	//
	log.Println("----------", "aaa", "------------")
	ylog.DebugDump(bean)
}

func Test_count(t *testing.T) {
	var opts *options.CountOptions

	i, err := db.Table("test", "abc").CountDocuments(
		context.Background(),
		bson.D{{"name", "abc"}}, opts)
	if err != nil {
		ylog.Error("####mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("---------count: ", i, "------------")
}

func Test_count_1(t *testing.T) {
	i, err := db.DoCount("test", "abc",
		bson.D{{"name", "abc"}})
	if err != nil {
		ylog.Error("####mongoClientWrapper_test.go->", err)
		return
	}
	log.Println("---------count: ", i, "------------")
}

func Test_a_b_find_many(t *testing.T) {
	l := make([]*ABC, 0)

	i := int64(5)
	opts := &options.FindOptions{
		Limit: &i,
		Sort: bson.D{{
			"Name",
			1,
		}},
	}

	err := db.DoFindMany(&l, "test", "abc",
		bson.M{"name": bson.M{"$ne": "abc"}}, opts)
	if err != nil {
		ylog.Error("### mongoClientWrapper_test.go->", err)
		return
	}

	//
	log.Println("----------", "aaa", "------------")
	ylog.DebugDump(l)
}

func Test_page_bean(t *testing.T) {
	bean := &ypage.PageBean{
		PageNo:      2,
		RowsPerPage: 10,
		PagesCount:  0,
		RowsCount:   0,
		Where:       bson.D{{}},
		Sort:        nil, //bson.D{{"name", 1}},
		Data:        nil,
	}

	l := make([]*ABC, 0)

	err := db.DoPage(&l, "test", "abc", bean)

	if err != nil {
		ylog.Error("mongoClientWrapper_test.go->", err)
		return
	}
	//
	fmt.Println("------", "", "-----------")
	ylog.DebugDump("", bean)
}
