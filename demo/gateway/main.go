package main

import (
	golog "log"
	"net/http"

	"vchat/demo/unit/intf"
	"vchat/lib"
)

func init() {
	//------------ prepare modules----------
	//本步骤主要是装入系统必备的模块
	_, err := lib.LoadModulesOfOptions(&lib.LoadOption{
		LoadMicroService: true, //这不同必需要的
		LoadEtcd:         true, //etcd必須開啟，否則無法自動發現服務
		LoadPg:           false,
		LoadRedis:        false,
		LoadMongo:        false,
		LoadMq:           false,
		LoadJwt:          false,
	})
	if err != nil {
		panic(err.Error())
	}
}

//gateway功能不需要每一个模块来实现，但用这个模块可以测试微服务是否能补成功调用
func main() {
	addr := "localhost:9999"
	//ctx := context.Background()

	http.Handle("/api/HelloWorld", new(intf.HelloWorldHandler).HandlerSD())
	//http.Handle("/api/SayGoodBye", new(intf.SayGoodByeHandler).HandlerSD())

	golog.Println(
		`start at :9999,url is curl:localhost/hello`,
		`test command:`,
		`curl -X POST  -H 'Content-Type:application/json'  -d '{"S":"hello,world pass in data"}' localhost:9999/HelloWorld -v`)

	golog.Fatal(http.ListenAndServe(addr, nil))
}
