// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xhttp

import (
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/webs"
	"net/http"
)

type router struct {
	methods  map[httpMethod]map[string]action
	handlers Handlers
}

func (me *router) GET(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodGet, pattern, handle, permissions...)
}

func (me *router) POST(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodPost, pattern, handle, permissions...)
}

func (me *router) PUT(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodPut, pattern, handle, permissions...)
}

func (me *router) DELETE(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodDelete, pattern, handle, permissions...)
}

func (me *router) PATCH(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodPatch, pattern, handle, permissions...)
}

func (me *router) HEAD(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodHead, pattern, handle, permissions...)
}

func (me *router) OPTIONS(pattern string, handle Handle, permissions ...string) {
	me.addRouter(httpMethodOptions, pattern, handle, permissions...)
}

// Attachs a global middleware to the router. ie. the middlewares attached though Use() will be
// included in the handlers chain for every single request. Even 404, 405, static files...
func (me *router) Use(middlewares ...Handler) Router {
	me.handlers = append(me.handlers, middlewares...)
	return me
}

func (me *router) isPermission(c Context, permission string) bool {
	if strings.IsBlank(permission) {
		return false
	}
	if c == nil || !c.Session().IsLogin() {
		return false
	}
	if p, err := c.Session().Principal(); err == nil {
		if strings.Contains(p.Permissions, permission) {
			return true
		}
	} else {
		logger.Error(err.Error())
		return false
	}
	return false
}

// ServeHTTP makes the router implement the http.Handler interface.
func (me *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if v, ok := me.methods[httpMethod(r.Method)]; ok {
		requestURI := r.RequestURI
		if strings.Contains(requestURI, "?") {
			requestURI = strings.Before(r.RequestURI, "?")
		}
		if a, aok := v[requestURI]; aok {
			c := me.newContext(w, r)
			if a.Permissions != nil && len(a.Permissions) > 0 {
				isPermission := false
				for _, permission := range a.Permissions {
					if me.isPermission(c, permission) {
						isPermission = true
					}
				}
				if !isPermission {
					if strings.IsNotBlank(Conf.Err.Err403) {
						http.Redirect(w, r, Conf.Err.Err403, http.StatusFound)
						return
					} else {
						w.WriteHeader(403)
						w.Write([]byte(default403Body))
						return
					}
				}
			}
			c.Next()
			a.Handle(c)
		} else {
			logger.Errorf("No match for router:%s:%s", r.Method, r.RequestURI)
			c := me.newContext(w, r)
			c.Next()
			if strings.IsNotBlank(Conf.Err.Err404) {
				http.Redirect(w, r, Conf.Err.Err404, http.StatusFound)
				return
			} else {
				serveError(c, 404, []byte(default404Body))
			}
		}
	} else {
		logger.Error("No match for router:", r.Method)
		c := me.newContext(w, r)
		c.Next()
		if strings.IsNotBlank(Conf.Err.Err404) {
			http.Redirect(w, r, Conf.Err.Err404, http.StatusFound)
			return
		} else {
			serveError(c, 404, []byte(default404Body))
		}
	}
}

func (me *router) addRouter(method httpMethod, pattern string, handle Handle, permissions ...string) {
	if me.methods == nil {
		me.methods = make(map[httpMethod]map[string]action)
	}
	if _, ok := me.methods[method]; !ok {
		me.methods[method] = make(map[string]action)
	}
	me.addAction(method, pattern, handle, permissions...)
}

func (me *router) addAction(method httpMethod, pattern string, handle Handle, permissions ...string) {
	actions := me.methods[method]
	if _, ok := actions[pattern]; !ok {
		a := action{
			Handle:      handle,
			Permissions: permissions,
		}
		actions[pattern] = a
	}
}

func (me *router) newContext(w http.ResponseWriter, r *http.Request) Context {
	values, err := webs.Values(r)
	if err != nil {
		logger.Error(err.Error())
	}
	s := newSession4Redis(w, r)
	return newContext(w, s, r, values, me.handlers)
}
