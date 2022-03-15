package framework

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
)

type Context struct {
	request    *http.Request
	response   http.ResponseWriter
	ctx        context.Context
	hasTimeOut bool
	writerMux  *sync.Mutex
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context {
	return &Context{
		request:   r,
		response:  w,
		ctx:       r.Context(),
		writerMux: &sync.Mutex{},
	}
}

func (c *Context) Json(status int, obj interface{}) error {
	c.response.Header().Set("Content-Type", "application/json")
	c.response.WriteHeader(status)
	bytes, err := json.Marshal(obj)
	if err != nil {
		c.response.WriteHeader(500)
		return err
	}
	_, err = c.response.Write(bytes)
	if err != nil {
		c.response.WriteHeader(500)
		return err
	}
	return nil
}
