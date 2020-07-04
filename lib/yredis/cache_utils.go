package yredis

import (
	"encoding/json"
	"fmt"

	"github.com/vhaoran/vchat/lib/ylog"

	"github.com/vhaoran/vchat/common/reflectUtils"
	"github.com/vhaoran/vchat/common/ymodel"
)

func CacheAutoGetH(ptrTableBean interface{}, field interface{},
	callback func() (interface{}, error)) (interface{}, error) {
	tbName := ymodel.TableName(ptrTableBean)
	key := CacheKeyTableH(tbName)
	fd := fmt.Sprint(field)
	ylog.Debug("redis cache key is:", key)
	//相当于定义了一个函数的别名 并且执行了一遍
	doCallbackAndSet := func() (interface{}, error) {
		v, err := callback()
		if err != nil {
			return nil, err
		}
		//println("123")
		var s []byte
		//正常执行,不出错误的话
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
		ylog.Debug("redis cache hit,key:", key, "  field: ", fd, " v: ", s)
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
