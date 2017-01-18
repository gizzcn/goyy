// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package templates

import (
	htmpl "html/template"
	ttmpl "text/template"

	"gopkg.in/goyy/goyy.v0/comm/profile"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"gopkg.in/goyy/goyy.v0/util/times"
)

// Text text funcMap
var Text = text{funcMapText}

// HTML html funcMap
var HTML = html{funcMapHTML}

type text struct {
	FuncMap ttmpl.FuncMap
}

type html struct {
	FuncMap htmpl.FuncMap
}

////////////////////////////////////////////////////////////
// template funcmap
////////////////////////////////////////////////////////////

var funcMapText = ttmpl.FuncMap{
	"yymd":     times.FormatYYMD,
	"yymdhms":  times.FormatYYMDHMS,
	"yymdhm":   times.FormatYYMDHM,
	"uyymd":    times.FormatUnixYYMD,
	"uyymdhms": times.FormatUnixYYMDHMS,
	"uyymdhm":  times.FormatUnixYYMDHM,

	"blank":     strings.IsBlank,
	"notblank":  strings.IsNotBlank,
	"abbr":      strings.Abbr,
	"anon":      strings.Anon,
	"anonymity": strings.Anonymity,

	"divisible": divisible,
	"exist":     exist,
	"zhstate":   zhstate,
	"eqindex":   eqindex,
	"eqshow":    eqshow,
	"eqadd":     eqadd,
	"eqedit":    eqedit,
	"neindex":   neindex,
	"neshow":    neshow,
	"neadd":     neadd,
	"needit":    needit,

	"profile":    getProfile,    // profile[production|development|test]
	"apis":       getAPIs,       // The URL of the apis.
	"assets":     getAssets,     // The URL of the static file.
	"statics":    getStatics,    // static file[web|wap|adm]
	"developers": getDevelopers, // The URL of the development of relevant documents
	"operations": getOperations, // The URL of the operation related documents.
	"uploads":    getUploads,    // The URL of the uploaded file.
}

var funcMapHTML = htmpl.FuncMap{
	"yymd":     times.FormatYYMD,
	"yymdhms":  times.FormatYYMDHMS,
	"yymdhm":   times.FormatYYMDHM,
	"uyymd":    times.FormatUnixYYMD,
	"uyymdhms": times.FormatUnixYYMDHMS,
	"uyymdhm":  times.FormatUnixYYMDHM,

	"blank":     strings.IsBlank,
	"notblank":  strings.IsNotBlank,
	"abbr":      strings.Abbr,
	"anon":      strings.Anon,
	"anonymity": strings.Anonymity,

	"set":       set,
	"divisible": divisible,
	"exist":     exist,
	"zhstate":   zhstate,
	"eqindex":   eqindex,
	"eqshow":    eqshow,
	"eqadd":     eqadd,
	"eqedit":    eqedit,
	"neindex":   neindex,
	"neshow":    neshow,
	"neadd":     neadd,
	"needit":    needit,

	"profile":    getProfile,    // profile[production|development|test]
	"apis":       getAPIs,       // The URL of the apis.
	"assets":     getAssets,     // The URL of the static file.
	"statics":    getStatics,    // static file[web|wap|adm]
	"developers": getDevelopers, // The URL of the development of relevant documents
	"operations": getOperations, // The URL of the operation related documents.
	"uploads":    getUploads,    // The URL of the uploaded file.
}

////////////////////////////////////////////////////////////
// common
////////////////////////////////////////////////////////////

var set = func(args map[string]interface{}, key string, value interface{}) htmpl.HTML {
	args[key] = value
	return htmpl.HTML("")
}

var divisible = func(x, y int) bool {
	if x%y == 0 {
		return true
	}
	return false
}

var exist = func(m map[string]interface{}, key string) (ok bool) {
	_, ok = m[key]
	return
}

////////////////////////////////////////////////////////////
// template state
////////////////////////////////////////////////////////////

var zhstate = func(t string) string {
	switch t {
	case enIndex:
		return zhAdd
	case enShow:
		return zhShow
	case enAdd:
		return zhAdd
	case enEdit:
		return zhEdit
	}
	return ""
}

var eqindex = func(t string) bool {
	if enIndex == t {
		return true
	}
	return false
}

var eqshow = func(t string) bool {
	if enShow == t {
		return true
	}
	return false
}

var eqadd = func(t string) bool {
	if enAdd == t {
		return true
	}
	return false
}

var eqedit = func(t string) bool {
	if enEdit == t {
		return true
	}
	return false
}

var neindex = func(t string) bool {
	return !eqindex(t)
}

var neshow = func(t string) bool {
	return !eqshow(t)
}

var neadd = func(t string) bool {
	return !eqadd(t)
}

var needit = func(t string) bool {
	return !eqedit(t)
}

////////////////////////////////////////////////////////////
// static state
////////////////////////////////////////////////////////////

func getProfile() string {
	return profile.Default()
}

var (
	// GetAPIs return to apis in static resources
	GetAPIs func() string
	// GetAssets return to assets in static resources
	GetAssets func() string
	// GetStatics return to statics in static resources
	GetStatics func() string
	// GetDevelopers return to developers in static resources
	GetDevelopers func() string
	// GetOperations return to operations in static resources
	GetOperations func() string
	// GetUploads return to uploads in static resources
	GetUploads func() string
)

func getAPIs() string {
	if GetAPIs == nil {
		return ""
	}
	return GetAPIs()
}

func getAssets() string {
	if GetAssets == nil {
		return ""
	}
	return GetAssets()
}

func getStatics() string {
	if GetStatics == nil {
		return ""
	}
	return GetStatics()
}

func getDevelopers() string {
	if GetDevelopers == nil {
		return ""
	}
	return GetDevelopers()
}

func getOperations() string {
	if GetOperations == nil {
		return ""
	}
	return GetOperations()
}

func getUploads() string {
	if GetUploads == nil {
		return ""
	}
	return GetUploads()
}
