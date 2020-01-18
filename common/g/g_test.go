package g

import (
	"fmt"
	"log"
	"testing"
)

type (
	Good struct {
		ID   int64
		Name string
	}
)

func Test_get_buffer(t *testing.T) {
	s, _ := GetBufferForMq("abc")
	log.Println(string(s))

	s, _ = GetBufferForMq([]byte("abc"))
	log.Println(string(s))

	s, _ = GetBufferForMq([]int{1, 2, 3})
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq(3)
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq(Good{ID: 1, Name: "33"})
	fmt.Println("------", string(s), "-----------")

	s, _ = GetBufferForMq([]Good{Good{ID: 1, Name: "33"}, Good{ID: 2, Name: "3344"}})
	fmt.Println("------", string(s), "-----------")
}
