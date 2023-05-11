package engine

import (
	"encoding/json"
	"github.com/tiyee/palmon/lib"
	"github.com/tiyee/palmon/schema"
	"github.com/tiyee/palmon/vo"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{w: w, r: r}
}
func (c *Context) Request() *http.Request {
	return c.r
}
func (c *Context) Response() http.ResponseWriter {
	return c.w
}
func (c *Context) JSONArgs(v schema.ISchema) error {
	if err := lib.JSONArgs(c.Request().Body, v); err != nil {
		return err
	}
	v.Hook()
	return nil

}
func (c *Context) AjaxJson(code int, message string, data any) {
	ret := vo.Base{
		Error:   code,
		Message: message,
		Data:    data,
	}
	c.Response().WriteHeader(200)
	if bs, err := json.Marshal(ret); err == nil {
		c.Response().Write(bs)
	}

}
func (c *Context) AjaxError(code int, message string, data any) {
	c.AjaxJson(code, message, data)
}
func (c *Context) AjaxSuccess(message string, data any) {
	c.AjaxJson(0, message, data)
}
