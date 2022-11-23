package routes

import "go.uber.org/fx"

var Module = fx.Options()

type Routes []Route

type Route interface {
	Setup()
}

func NewRoutes() Routes {
	return Routes{}
}

func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
