// generated by xgen -- DO NOT EDIT
package blacklist

import (
	"gopkg.in/goyy/goyy.v0/web/controller"
)

var ctl = &Controller{
	JSONController: controller.JSONController{
		Settings: controller.Settings{
			Project: "sys",
			Module:  "blacklist",
			Title:   "BLACKLIST",
		},
		Mgr: Mgr,
	},
}

type Controller struct {
	controller.JSONController
}
