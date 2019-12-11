package yconfig

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_getyml_config(t *testing.T) {
	bean, err := GetYmlConfig()
	fmt.Println("-----------------", bean, err)
	fmt.Println("-----------------")
	spew.Dump(bean)
}
