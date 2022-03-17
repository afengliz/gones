package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/afengliz/gones/framework"
)

func footController(ctx *framework.Context) error {
	finishChan := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)
	tCtx, cancelFun := context.WithTimeout(ctx, time.Second)
	defer cancelFun()
	go func() {
		defer func() {
			if r := recover(); r != nil {
				panicChan <- r
			}
		}()
		time.Sleep(10 * time.Second)
		ctx.Json(200, "Hello liyanfneg")
		finishChan <- struct{}{}
	}()
	select {
	case <-finishChan:
		fmt.Println("finish")
	case p := <-panicChan:
		log.Println(p)
		ctx.Json(500, "panic")
	case <-tCtx.Done():
		ctx.Json(500, "time out")
		ctx.SetHasTimeOut()
	}
	return nil
}
