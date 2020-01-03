package ymid

import (
	"time"

	"github.com/go-kit/kit/endpoint"

	"vchat/lib/ylog"
)

//内置中间件
func MidCommon(ep endpoint.Endpoint) endpoint.Endpoint {
	t0 := time.Now()
	defer func() {
		ylog.Debug("#####time.Since:", time.Since(t0))
	}()
	
	return ep
}
