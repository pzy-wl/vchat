package ykit

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/etcdv3"
	"github.com/go-kit/kit/sd/lb"
	tran "github.com/go-kit/kit/transport/http"

	"vchat/lib/yetcd"
)

/*--auth: whr  date:2019-12-05--------------------------
 这是微服务是基类，所有实现微服务的接口都
 需要"继承"此类，用于快速开发及部署
--------------------------------------- */
type (
	RootTran struct {
	}
)

func (r *RootTran) DecodeRequest(reqDataPtr interface{}, _ context.Context, req *http.Request) (interface{}, error) {
	spew.Dump("RootTran->DecodeRequest:reqDatePtr", reqDataPtr)
	//	spew.Dump("RootTran->DecodeRequest:req", req.Body)

	if err := json.NewDecoder(req.Body).Decode(reqDataPtr); err != nil {
		return nil, err
	}
	return reqDataPtr, nil
}

func (r *RootTran) EncodeRequestBuffer(_ context.Context, res *http.Request, requestData interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(requestData); err != nil {
		return err
	}
	res.Body = ioutil.NopCloser(&buf)
	return nil
}

func (r *RootTran) EncodeResponse(_ context.Context, wr http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(wr).Encode(res)
}

//manual proxy
func (r *RootTran) MakeProxyEndPoint(
	instance,
	method,
	path string,
	decodeResponse func(_ context.Context, r *http.Response) (interface{}, error),
	ctx context.Context) endpoint.Endpoint {
	//instance := "localhost:9999"

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	u, err := url.Parse(instance)
	if err != nil {
		panic(err)
	}
	u.Path = path
	return tran.NewClient(
		method,
		u,
		r.EncodeRequestBuffer,
		decodeResponse,
	).Endpoint()
}

//service auto discovery
func (r *RootTran) HandlerSD(ctx context.Context,
	serviceTag, method, path string,
	decodeRequestFunc func(ctx context.Context, req *http.Request) (interface{}, error),
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error)) *tran.Server {
	var err error
	var client etcdv3.Client
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	//etcdAddr   := flag.String("consul.addr", "", "Consul agent address")
	retryMax := 3
	retryTimeout := 500 * time.Millisecond

	//

	//etcdAddr := "127.0.0.1:2379"
	//options := etcdv3.ClientOptions{
	//	DialTimeout:   time.Second * 10,
	//	DialKeepAlive: time.Second * 50,
	//}

	if client, err = etcdv3.NewClient(ctx, yetcd.XETCDConfig.Hosts, yetcd.XETCDConfig.Options); err != nil {
		return nil
	}

	//
	instance, err := etcdv3.NewInstancer(client, serviceTag, logger)
	if err != nil {
		return nil
	}

	//
	factory := r.FactorySD(ctx, method, path, decodeResponseFunc)
	endPointer := sd.NewEndpointer(instance, factory, logger)
	balance := lb.NewRoundRobin(endPointer)
	retry := lb.Retry(retryMax, retryTimeout, balance)
	e := retry
	//
	return tran.NewServer(e, decodeRequestFunc, r.EncodeResponse)
}

// service discovery
func (r *RootTran) FactorySD(ctx context.Context, method, path string,
	decodeResponseFunc func(_ context.Context, res *http.Response) (interface{}, error)) sd.Factory {
	return func(instance string) (endpoint.Endpoint, io.Closer, error) {
		if !strings.HasPrefix(instance, "http") {
			instance = "http://" + instance
		}
		targetURL, err := url.Parse(instance)
		if err != nil {
			return nil, nil, err
		}
		targetURL.Path = path

		enc := r.EncodeRequestBuffer
		dec := decodeResponseFunc

		return tran.NewClient(method, targetURL, enc, dec).Endpoint(), nil, nil
	}
}
