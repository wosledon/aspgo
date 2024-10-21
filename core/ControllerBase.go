package core

import (
	"log"
	"net/http"
	"reflect"
	"strings"
)

type Controller interface {
	RegisterRoutes(mux *http.ServeMux, instance interface{})
}

type ControllerBase struct {
}

func (c *ControllerBase) RegisterRoutes(mux *http.ServeMux, instance interface{}) {
	t := reflect.TypeOf(instance)
	v := reflect.ValueOf(instance)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		if strings.HasPrefix(method.Name, "Handle") {
			log.Printf("Registering route: /%s", strings.ToLower(method.Name[6:]))
			mux.HandleFunc("/"+strings.ToLower(method.Name[6:]), func(w http.ResponseWriter, r *http.Request) {
				v.Method(i).Call([]reflect.Value{reflect.ValueOf(w), reflect.ValueOf(r)})
			})
		}
	}
}
