package main

import "github.com/afengliz/gones/framework"

func registerRouter(core *framework.Core) {
	core.Get("/foo", footController)
}
