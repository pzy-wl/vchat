package ykit

import (
	"context"
	"fmt"
	tran "github.com/go-kit/kit/transport/http"
	"github.com/vhaoran/vchat/lib/yjwt"
	"github.com/vhaoran/vchat/lib/ylog"
	"net/http"
	"strconv"
	"strings"
)

const (
	JWT_TOKEN = "Jwt"
	UID_KEY   = "Uid"
)

func Jwt2ctx() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		jwt := req.Header.Get(JWT_TOKEN)
		if len(jwt) > 0 {
			return context.WithValue(ctx, JWT_TOKEN, jwt)
		}
		return ctx
	}
}

func Jwt2Req() tran.RequestFunc {
	return func(ctx context.Context, req *http.Request) context.Context {
		c := ctx
		a := ctx.Value(JWT_TOKEN)
		ylog.Debug("----------", "jwt raw", "------------")

		l, ok := a.(string)
		if !ok {
			return c
		}

		if len(l) == 0 {
			return c
		}

		req.Header.Set(JWT_TOKEN, l)
		return c
	}
}

func UIDOfTest(s string) int64 {
	if strings.Contains(s, "/") {
		l := strings.Split(s, "/")
		if len(l) > 1 {
			uid, err := strconv.ParseInt(l[1], 10, 64)
			if err != nil {
				ylog.Error("jwt_utils.go->", err)
				return 0
			}

			return uid
		}
		return 0
	}
	return 0
}

func GetUIDOfReq(req *http.Request) int64 {
	s := req.Header.Get(JWT_TOKEN)
	if len(s) > 0 {
		if i := UIDOfTest(s); i > 0 {
			return i
		}

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

	if i != nil {
		uid, err := strconv.ParseInt(fmt.Sprint(i), 10, 64)
		if err != nil {
			ylog.Error("jwt_utils.go->", err)
			return 0
		}
		return uid
	}

	//--------not found uid then parse jwt -----------------------------
	i = ctx.Value(JWT_TOKEN)
	s, ok := i.(string)
	if !ok {
		return 0
	}

	if len(s) == 0 {
		return 0
	}
	if i := UIDOfTest(s); i > 0 {
		return i
	}

	ii, err := yjwt.Parse(s)
	if err != nil {
		ylog.Error("jwt_utils.go->", err)
		return 0
	}
	return ii
}
