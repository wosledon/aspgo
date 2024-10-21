package core

import "reflect"

type ServiceProvider struct {
	Services    *ServiceCollection
	Controllers *ServiceCollection
}

func NewServiceProvider(services *ServiceCollection, controllers *ServiceCollection) *ServiceProvider {
	return &ServiceProvider{
		Services:    services,
		Controllers: controllers,
	}
}

func (sp *ServiceProvider) GetService(serviceType interface{}) interface{} {
	return sp.Services.Services[reflect.TypeOf(serviceType)].Implementation
}

func (sp *ServiceProvider) GetServiceImpl(serviceType reflect.Type) interface{} {
	return sp.Services.Services[serviceType].Implementation
}
