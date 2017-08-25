package template_functions

import (
	"flamingo/core/pug_template/pugjs"
	"flamingo/framework/router"
	"flamingo/framework/web"
)

type (
	// GetFunc allows templates to access the router's `get` method
	GetFunc struct {
		Router *router.Router `inject:""`
	}
)

// Name alias for use in template
func (g GetFunc) Name() string {
	return "get"
}

// Func as implementation of get method
func (g *GetFunc) Func(ctx web.Context) interface{} {
	return func(what string, params ...*pugjs.Map) interface{} {
		var p = make(map[interface{}]interface{})
		if len(params) == 1 {
			for k, v := range params[0].Items {
				p[k.String()] = v.String()
			}
		}
		return g.Router.Get(what, ctx, p)
	}
}