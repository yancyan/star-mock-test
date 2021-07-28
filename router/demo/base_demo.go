package demo

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseDemo struct {
}

func (cr *BaseDemo) InitRouter(router *gin.RouterGroup) gin.IRoutes {
	testGroup := router.Group("test")
	testGroup.GET("/test01", Test01)
	return testGroup
}

func Test01(cxt *gin.Context) {
	cxt.String(http.StatusOK, "test - 01 ok...")
}
