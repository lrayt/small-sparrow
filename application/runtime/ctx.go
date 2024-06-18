package runtime

import "context"

type Context struct {
	c context.Context
}

func NewCTX() *Context {
	return &Context{c: context.Background()}
}

func (ctx Context) GetLogger() {

}
