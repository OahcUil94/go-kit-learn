package service

import (
	"github.com/king19800105/go-kit-learn/demo1/pkg/service"
	"github.com/king19800105/go-kit-learn/demo1/pkg/endpoint"
	"github.com/king19800105/go-kit-learn/demo1/pkg/http"
	nethttp "net/http"
	kithttp "github.com/go-kit/kit/transport/http"
)

// 创建服务
func createService() nethttp.Handler {
	// 创建业务对象
	svc := service.New(nil)
	// 创建端点对象
	eps := endpoint.New(svc, nil)
	// 设置http服务服务中间件
	options := defaultHttpOptions()
	// 端点绑定到http服务上
	return http.NewHTTPHandler(eps, options)
}

// HTTP服务中间件（服务的aop）
func defaultHttpOptions() map[string][]kithttp.ServerOption {
	options := map[string][]kithttp.ServerOption{
		"Create": {
			kithttp.ServerErrorEncoder(http.ErrorEncoder),
		},
	}
	return options
}
