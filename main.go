package main

import (
	"net/http"

	"github.com/afengliz/gones/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := http.Server{
		Addr:    ":8250",
		Handler: core,
	}
	server.ListenAndServe()
}
