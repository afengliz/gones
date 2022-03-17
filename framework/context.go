package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var _ context.Context = (*Context)(nil)

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
func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}
func (c *Context) Done() <-chan struct{} {
	return c.ctx.Done()
}

func (c *Context) Err() error {
	return c.ctx.Err()
}
func (c *Context) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}
func (c *Context) Json(status int, obj interface{}) error {
	if c.hasTimeOut {
		return nil
	}
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
func (c *Context) GetRequest() *http.Request {
	return c.request
}

func (c *Context) GetResponse() http.ResponseWriter {
	return c.response
}
func (c *Context) SetHasTimeOut() {
	c.hasTimeOut = true
}
func (c *Context) HasTimeOut() bool {
	return c.hasTimeOut
}

func (c *Context) BaseContext() context.Context {
	return c.ctx
}

func (c *Context) WriterMux() *sync.Mutex {
	return c.writerMux
}

func (c *Context) QueryAll() map[string][]string {
	if c.request != nil {
		return map[string][]string(c.request.URL.Query())
	}
	return map[string][]string{}
}

func (c *Context) QueryInt(key string, def int) int {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			intVal, err := strconv.Atoi(vals[len(vals)-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (c *Context) QueryString(key string, def string) string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[len(vals)-1]
		}
	}
	return def
}

func (c *Context) QueryArray(key string, def []string) []string {
	params := c.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals
		}
		return def
	}
	return def
}
func (ctx *Context) FormAll() map[string][]string {
	if ctx.request != nil && ctx.request.PostForm != nil {
		return map[string][]string(ctx.request.PostForm)
	}
	return map[string][]string{}
}

func (c *Context) FormInt(key string, def int) int {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			intVal, err := strconv.Atoi(vals[len(vals)-1])
			if err != nil {
				return def
			}
			return intVal
		}
	}
	return def
}

func (c *Context) FormString(key string, def string) string {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[len(vals)-1]
		}
	}
	return def
}

func (c *Context) FormArray(key string, def []string) []string {
	params := c.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals
		}
		return def
	}
	return def
}
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}

func (c *Context) HTML(status int, obj interface{}, template string) error {
	return nil
}

func (c *Context) Text(status int, obj string) error {
	return nil
}
