package core

import (
	"net/http"
	"reflect"

	"go.uber.org/fx"
)

type WebApplicationBuilder struct {
	services    *ServiceCollection
	controllers *ServiceCollection
	mux         *http.ServeMux
}

func (builder *WebApplicationBuilder) AddService(implementation interface{}) {
	builder.services.AddService(implementation)
}

func (builder *WebApplicationBuilder) AddController(constructor interface{}) {
	// 检测是否是构造函数
	if !isConstructor(constructor) {
		panic("The controller must be a constructor")
	}
	builder.controllers.AddService(constructor)
}

func isConstructor(constructor interface{}) bool {
	return reflect.TypeOf(constructor).Kind() == reflect.Func
}

type WebApplication struct {
	ServiceProvider *ServiceProvider
	Mux             *http.ServeMux
	middlewares     []func(http.Handler) http.Handler
	DI              *fx.App
}

func CreateBuilder() *WebApplicationBuilder {
	return &WebApplicationBuilder{
		services:    NewServiceCollection(),
		controllers: NewServiceCollection(),
		mux:         http.NewServeMux(),
	}
}

func (builder *WebApplicationBuilder) Build() *WebApplication {
	return NewWebApplication(NewServiceProvider(
		builder.services,
		builder.controllers,
	), builder.mux)
}

func (app *WebApplication) Run(address string) {
	app.DI = fx.New(
		app.registerServices(),
		app.registerControllers(),

		fx.Invoke(func() {
			finalHandler := app.registerMiddlewares(app.Mux)

			http.ListenAndServe(address, finalHandler)
		}),
	)
	app.DI.Run()
}

func NewWebApplication(serviceProvider *ServiceProvider, mux *http.ServeMux) *WebApplication {
	return &WebApplication{
		ServiceProvider: serviceProvider,
		Mux:             mux,
	}
}

func (app *WebApplication) Use(middleware func(http.Handler) http.Handler) {
	app.middlewares = append(app.middlewares, middleware)
}

func (app *WebApplication) UseMiddleware(middleware Middleware) {
	app.middlewares = append(app.middlewares, middleware.Invoke)
}

func (app *WebApplication) UseStaticFiles(path string) {
	app.Mux.Handle("/wwwroot/", http.StripPrefix("/wwwroot/", http.FileServer(http.Dir(path))))
}

func (app *WebApplication) registerControllers() fx.Option {
	var impls []interface{}
	for _, service := range app.ServiceProvider.Controllers.Services {
		constructor := reflect.ValueOf(service.Implementation)
		instance := constructor.Call(nil)[0].Interface()
		controller := instance.(Controller)
		controller.RegisterRoutes(app.Mux, controller)

		impls = append(impls, service.Implementation)
	}

	return fx.Provide(impls...)
}

func (app *WebApplication) registerServices() fx.Option {
	var impls []interface{}
	for _, value := range app.ServiceProvider.Services.Services {
		impls = append(impls, value.Implementation)
	}

	return fx.Provide(impls...)
}

func (app *WebApplication) registerMiddlewares(handler http.Handler) http.Handler {
	for i := len(app.middlewares) - 1; i >= 0; i-- {
		handler = app.middlewares[i](handler)
	}
	return handler
}
