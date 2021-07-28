package router

import (
	"kingdee-test/router/demo"
	"kingdee-test/router/partner"
)

type Router interface {
	InitRouter()
}

type RouteGroup struct {
	Partner partner.RouteGroup
	Demo    demo.RouteGroup
}

var RouteGroupApp = new(RouteGroup)
