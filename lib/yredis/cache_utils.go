package yredis

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/weihaoranW/vchat/common/reflectUtils"
	"github.com/weihaoranW/vchat/common/ymodel"
)

func CacheAutoGetH(ptrTableBean interface{}, field interface{},
	callback func(field interface{}) (interface{}, error)) (interface{}, error) {
	tbName := ymodel.TableName(ptrTableBean)
	key := TableCacheH(tbName)
	log.Println("cache key is:", key)
	fd := fmt.Sprint(field)

	doCallbackAndSet := func() (interface{}, error) {
		v, err := callback(field)
		if err != nil {
			return nil, err
		}
		//
		if s, err := json.Marshal(v); err == nil {
			_, err := X.HSet(key, fd, string(s)).Result()
			err != nil{

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
