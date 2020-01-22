package ykit

import (
	"context"
	"fmt"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/yjwt"
	"github.com/vhaoran/vchat/lib/ylog"
	"log"
	"net/http"
	"strconv"
)

const (
	JWT_TOKEN = "jwt"
)

func Jwt2ctx() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		jwt := req.Header.Get(JWT_TOKEN)
		if len(jwt) > 0 {
			ylog.Debug("fetch jwt from request and set to context :", jwt)
			return context.WithValue(ctx, JWT_TOKEN, jwt)
		}
		return ctx
	}
}

func Jwt2Req() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		c := ctx
		a := ctx.Value(JWT_TOKEN)
		log.Println("----------", "jtw raw", "------------")
		ylog.DebugDump("jwt dump: ", a)

		l, ok := a.(string)
		if !ok {
			ylog.Debug("not get jwt token")
			return c
		}

		if len(l) == 0 {
			ylog.Debug("not get jwt or jwt is empty")
			return c
		}

		req.Header.Set(JWT_TOKEN, l)
		ylog.Debug("set jwt :", l)
		return c
	}
}

func GetUIDOfReq(req *http.Request) int64 {
	s := req.Header.Get(JWT_TOKEN)
	if len(s) > 0 {
		uid, err := yjwt.Parse(s)
		if err != nil {
			ylog.Error("jwt_utils.go->", err)
			return 0
		}
		return uid
	}
	return 0
}

func GetUIDOfContext(ctx context.Context) int64 {
	i := ctx.Value("Uid")

	uid, err := strconv.ParseInt(fmt.Sprint(i), 10, 64)
	if err != nil {
		ylog.Error("jwt_utils.go->", err)
		return 0
	}
	return uid
}
