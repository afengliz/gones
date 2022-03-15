package main

import "github.com/afengliz/gones/framework"

func footController(ctx *framework.Context) error {
	return ctx.Json(200, "Hello liyanfneg")
}
