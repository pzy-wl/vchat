package ymongo

import (
	"context"
	"errors"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"vchat/common/reflectUtils"
)

type (
	MongoClientWrapper struct {
		Base *mongo.Client
	}
)

func (r *MongoClientWrapper) DoInsertOne(dbName, tbName string,
	doc interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	var ctx = context.Background()

	client := r.Base
	log.Println("cnt ok")

	//var doc = bson.M{"a": 100, "b": 30}
	//d
	//dbName, tbName := "test", "t"
	tb := client.Database(dbName).Collection(tbName)
	return tb.InsertOne(ctx, doc, opts...)
}

func (r *MongoClientWrapper) tran2Slice(a interface{}) ([]interface{}, error) {
	//only suupoer struct pointer or slice,array
	l := make([]interface{}, 0)

	v := reflect.Indirect(reflect.ValueOf(a))
	if !v.IsValid() {
		return nil, errors.New("wrong type,only support array/slice/struct/pointer of struct")
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		{
			for i := 0; i < v.Len(); i++ {
				l = append(l, v.Index(i).Interface())
			}
			return l, nil
		}
	case reflect.Struct:
		{
			l = append(l, v.Interface())
			return l, nil
		}
	}

	return nil, errors.New("data is empty")
}

func (r *MongoClientWrapper) DoInsertMany(dbName, tbName string,
	doc interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	var ctx = context.Background()
	db := r.Base

	l, err := r.tran2Slice(doc)
	if err != nil {
		return nil, err
	}

	//var doc = bson.M{"a": 100, "b": 30}
	//d
	//dbName, tbName := "test", "t"
	tb := db.Database(dbName).Collection(tbName)
	return tb.InsertMany(ctx, l, opts...)
}

func (r *MongoClientWrapper) DoFind(retSlicePtr interface{},
	dbName, tbName string,
	filter interface{},
	opts ...*options.FindOptions) error {
	if !reflectUtils.IsSlice(retSlicePtr) || !reflectUtils.IsPointer(retSlicePtr) {
		return errors.New("not supported type,muse be pointer of slice,may be you need &slice")
	}

	var ctx = context.Background()
	db := r.Base
	tb := db.Database(dbName).Collection(tbName)

	c, err := tb.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}

	l := reflect.Indirect(reflect.ValueOf(retSlicePtr))
	for c.Next(ctx) {
		if bean, err := reflectUtils.MakeSliceElemPtr(retSlicePtr); err == nil {
			//log.Println("--BBB-beanType---", reflect.TypeOf(bean), "-----------")

			if err = bson.Unmarshal(c.Current, bean); err != nil {
				return err
			}
			//log.Println("--lType----", l.Type(), "-----------")
			//log.Println("---beanType---", reflect.TypeOf(bean), "-----------")
			//log.Println("------", reflect.ValueOf(bean), "-----------")
			l = reflect.Append(l, reflect.ValueOf(bean))
		} else {
			return err
		}
	}

	//rewrite
	v := reflect.ValueOf(retSlicePtr)
	v.Elem().Set(l)
	return nil
}
