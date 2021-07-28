package demo

import "github.com/gin-gonic/gin"

type RequestValidateDemo struct {
}

func (cr *RequestValidateDemo) InitRouter(router *gin.RouterGroup) gin.IRoutes {
	testGroup := router.Group("request")
	testGroup.GET("/test01", RvTest01)
	return testGroup
}

func RvTest01(cxt *gin.Context) {

}
