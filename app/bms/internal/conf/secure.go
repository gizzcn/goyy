package conf

import (
	"gopkg.in/goyy/goyy.v0/comm/xtype"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func init() {
	xhttp.Conf.Secure.ForbidUrl = "/err/403"
	xhttp.Conf.Secure.SuccessUrl = "/"
	xhttp.Conf.Secure.Filters = []xtype.Map{
		{"/**.(css|js|html|map)", "anon"},
		{"/**.(jpg|gif|png|bmp|ico)", "anon"},
		{"/**.html", "anon"},
		{"/login", "anon"},
		{"/signin", "anon"},
		{"/logout", "anon"},
		{"/err/**", "anon"},
		{"/**", "authc"},
	}
}