package main

import (
	"net/http"

	"github.com/afengliz/gones/framework"
)

func main() {
	server := http.Server{
		Addr:    ":8888",
		Handler: framework.NewCore(),
	}
	server.ListenAndServe()
}
