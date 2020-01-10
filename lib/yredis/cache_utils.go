package yredis

import (
	"encoding/json"
	"fmt"
	"github.com/weihaoranW/vchat/lib/ylog"
	"log"

	"github.com/weihaoranW/vchat/common/reflectUtils"
	"github.com/weihaoranW/vchat/common/ymodel"
)

func CacheAutoGetH(ptrTableBean interface{}, field interface{},
	callback func(field interface{}) (interface{}, error)) (interface{}, error) {
	tbName := ymodel.TableName(ptrTableBean)
	key := CacheKeyTableH(tbName)
	log.Println("cache key is:", key)
	fd := fmt.Sprint(field)

	doCallbackAndSet := func() (interface{}, error) {
		v, err := callback(field)
		if err != nil {
			return nil, err
		}
		//
		var s []byte
		if s, err = json.Marshal(v); err == nil {
			_, err = X.HSet(key, fd, string(s)).Result()
			if err != nil {
				ylog.Error("cache_utils.go->", err)
				return v, nil
			}
		}
		return v, nil
	}

	//
	s, err := X.HGet(key, fd).Result()
	// if find
	if err == nil {
		obj, err := reflectUtils.MakeStructObj(ptrTableBean)
		if err != nil {
			if obj, err = doCallbackAndSet(); err != nil {
				return obj, nil
			}
			return nil, err
		}
		err = json.Unmarshal([]byte(s), obj)
		return obj, nil
	}

	//if not find
	obj1, err1 := doCallbackAndSet()
	if err1 != nil {
		return nil, err1
	}

	return obj1, nil
}

func CacheClearH(ptrTableBean interface{}, fields ...interface{}) {
	tbName := ymodel.TableName(ptrTableBean)
	key := CacheKeyTableH(tbName)

	l := make([]string, 0)
	for _, v := range fields {
		l = append(l, fmt.Sprint(v))
	}
	_ = X.HDel(key, l...)
}
