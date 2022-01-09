package param

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	keyInCtx = "param:ctx:key"
)

type Param struct {
	HTTPListenAddr string
	LogrusEntry    *logrus.Entry
	CorsHosts      []string
}

func Inject(e *gin.Engine, p *Param) {
	e.Use(func(ctx *gin.Context) {
		ctx.Set(keyInCtx, p)
		c := context.WithValue(ctx.Request.Context(), keyInCtx, ctx)
		ctx.Request = ctx.Request.WithContext(c)
		ctx.Next()
	})
}

func Eject(ctx *gin.Context) *Param {
	v, _ := ctx.Get(keyInCtx)
	return v.(*Param)
}
