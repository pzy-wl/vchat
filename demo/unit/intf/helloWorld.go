package intf

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	tran "github.com/go-kit/kit/transport/http"

	"vchat/lib/ykit"
)

const (
	HelloWorld_HANDLER_PATH = "/HelloWorld"
)

type (
	HelloWorldService interface {
		Hello(in *HelloWorldRequest) (string, error)
	}
	//input data
	HelloWorldRequest struct {
		S string `json:"s"`
	}

	//output data
	//Result struct {
	//	Code int         `json:"code"`
	//	Msg  string      `json:"msg"`
	//	Data interface{} `json:"data"`
	//}

	// handler implements
	HelloWorldHandler struct {
		base ykit.RootTran
	}
)

func (r *HelloWorldHandler) MakeLocalEndpoint(svc HelloWorldService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//modify
		in := request.(*HelloWorldRequest)
		ret, err := svc.Hello(in)
		if err != nil {
			return &ykit.Result{
				Code: 500,
				Msg:  "服务器内部错误",
				Data: nil,
			}, err
		}
		return &ykit.Result{
			Code: 200,
			Msg:  "ok",
			Data: ret,
		}, nil
	}
}

//个人实现,参数不能修改
func (r *HelloWorldHandler) DecodeRequest(ctx context.Context, req *http.Request) (interface{}, error) {
	return r.base.DecodeRequest(new(HelloWorldRequest), ctx, req)
}

//个人实现,参数不能修改
func (r *HelloWorldHandler) DecodeResponse(_ context.Context, res *http.Response) (interface{}, error) {
	var response ykit.Result
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

//handler for router，微服务本地接口，
func (r *HelloWorldHandler) HandlerLocal(service HelloWorldService) *tran.Server {
	endpoint := r.MakeLocalEndpoint(service)
	handler := tran.NewServer(
		endpoint,
		r.DecodeRequest,
		r.base.EncodeResponse,
	)
	//handler = loggingMiddleware()
	return handler
}

//sd,proxy实现,用于etcd自动服务发现时的handler
func (r *HelloWorldHandler) HandlerSD() *tran.Server {
	return r.base.HandlerSD(
		context.Background(),
		"api",
		"POST",
		HelloWorld_HANDLER_PATH,
		r.DecodeRequest,
		r.DecodeResponse)
}

// for test
//测试proxy方式的实现,用於測試某一微服務的運行情況
func (r *HelloWorldHandler) HandlerProxyForTest() *tran.Server {
	e := r.MakeProxyEndPointForTest(context.Background())
	handler := tran.NewServer(
		e,
		r.DecodeRequest,
		r.base.EncodeResponse,
	)

	return handler
}

// for test
//sd,proxy实现,调用 指定位置的endPoint
func (r *HelloWorldHandler) MakeProxyEndPointForTest(
	ctx context.Context) endpoint.Endpoint {
	//modify
	return r.base.MakeProxyEndPoint(
		//此为被调用的微服务的(host:port),
		"localhost:9001",
		"POST",
		HelloWorld_HANDLER_PATH,
		r.DecodeResponse,
		ctx)
}
