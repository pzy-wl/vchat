package ykit

import (
	"context"
	"net/http"

	tran "github.com/go-kit/kit/transport/http"
)

//不必yyoc原来的head,无条件转发
func OriginHead() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		c := ctx
		for k, v := range req.Header {
			c = context.WithValue(c, k, v)
		}
		
		return ctx
	}
}
