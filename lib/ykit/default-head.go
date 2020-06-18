package ykit

import (
	"context"
	"net/http"

	tran "github.com/go-kit/kit/transport/http"

	"github.com/vhaoran/vchat/lib/ylog"
)

func CommonHead() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		req.Header.Set("Content-Type", "application/json;charset:utf-8")
		//req.Header.Set("Content-Type", "charset:utf-8")
		req.Header.Set("Accept", "*/*")

		return ctx
	}
}

func DebugHead() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		for k, v := range req.Header {
			ylog.Debug("default-head.go->header: ", k, ":", v)
		}
		ylog.Debug("--------visit:", req.URL)

		return ctx
	}
}
