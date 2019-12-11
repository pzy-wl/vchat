package ymongo

import (
	"context"
	"log"
	"vchat/lib/yconfig"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	XMongo *mongo.Client
)

func InitMongo(cfg yconfig.MongoConfig) error {
	cnt, err := NewMongoClient(cfg)
	if err != nil {
		return err
	}
	XMongo = cnt
	return nil
}

func NewMongoClient(cfg yconfig.MongoConfig) (*mongo.Client, error) {
	var err error
	var client *mongo.Client
	// for cluster
	//uri := "mongodb://localhost/argos?replicaSet=replset&authSource=admin"
	// for single host
	//mongodb://mongodb0.example.com:27017/admin

	//uri := "mongodb://root:password@192.168.0.99:27017/test?&authSource=admin"
	uri := cfg.URL
	//uri := "mongodb://root:password@192.168.0.99:27017/test"

	opts := options.Client()
	opts.ApplyURI(uri)

	//set other properties
	//todo whr
	opts.SetMaxPoolSize(cfg.Options.MaxPoolSize)
	opts.SetMinPoolSize(cfg.Options.MinPoolSize)
	//opts.SetMaxConnIdleTime(cfg.Options.MaxConnIdleTime * time.Second)

	if client, err = mongo.Connect(context.Background(), opts); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return client, nil
}
