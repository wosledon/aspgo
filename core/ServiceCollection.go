package core

import "reflect"

type ServiceDescription struct {
	ServiceType    reflect.Type
	Implementation interface{}
}

type ServiceCollection struct {
	Services map[reflect.Type]ServiceDescription
}

func NewServiceCollection() *ServiceCollection {
	return &ServiceCollection{
		Services: make(map[reflect.Type]ServiceDescription),
	}
}

func (sc *ServiceCollection) AddService(implementation interface{}) {
	serviceType := reflect.TypeOf(implementation)
	sc.Services[serviceType] = ServiceDescription{
		ServiceType:    serviceType,
		Implementation: implementation,
	}
}

func (sc *ServiceCollection) AddServiceImpl(serviceType reflect.Type, implementation interface{}) {
	sc.Services[serviceType] = ServiceDescription{
		ServiceType:    serviceType,
		Implementation: implementation,
	}
}

func (sc *ServiceCollection) BuildServiceProvider() *ServiceProvider {
	return &ServiceProvider{
		Services: sc,
	}
}
