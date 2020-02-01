package ypage

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"strconv"
)

func GetSort(l bson.D) string {
	s := ""
	for _, v := range l {
		str := fmt.Sprint(v.Value)
		//
		i, _ := strconv.Atoi(str)
		sign := " asc "
		if i < 0 {
			sign = " desc "
		}
		if len(s) == 0 {
			s = v.Key + sign
		} else {
			s += "," + v.Key + sign
		}
	}
	if len(s) == 0 {
		s = " id asc "
	}

	return s
}
