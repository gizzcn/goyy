// generated by xgen -- DO NOT EDIT
package blacklist

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/service"
)

var Mgr = &Manager{
	Manager: service.Manager{
		Entity: func() entity.Interface {
			return NewEntity()
		},
		Entities: func() entity.Interfaces {
			return NewEntities(10)
		},
	},
}

type Manager struct {
	service.Manager
}
