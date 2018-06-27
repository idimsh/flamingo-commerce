package templatefunctions

import (
	"flamingo.me/flamingo-commerce/w3cDatalayer/application"
	"flamingo.me/flamingo/framework/web"
)

type (
	W3cDatalayerService struct {
		ApplicationServiceProvider application.ServiceProvider `inject:""`
	}
)

// Name alias for use in template
func (w3cdl W3cDatalayerService) Name() string {
	return "w3cDatalayerService"
}

// Func template function factory
func (w3cdl W3cDatalayerService) Func(ctx web.Context) interface{} {
	// Usage
	// w3cDatalayerService().get()
	return func() *application.Service {
		service := w3cdl.ApplicationServiceProvider()
		service.Init(ctx)
		return service
	}
}