package yredis

import (
	"fmt"
)

//支持表名+主键的hashSet
func CacheKeyTableH(tbName string) string {
	return fmt.Sprintf("/tableCache/%s", tbName)
}
