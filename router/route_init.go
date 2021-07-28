package router

import (
	"github.com/gin-gonic/gin"
	"kingdee-test/requests"
)

func InitRouter() *gin.Engine {

	var rt = gin.Default()

	//customerR := router.RouteGroupApp.Customer
	demoR := RouteGroupApp.Demo
	partnerR := RouteGroupApp.Partner

	PublicGroup := rt.Group("")
	{
		demoR.BaseDemo.InitRouter(PublicGroup)
		demoR.RequestValidateDemo.InitRouter(PublicGroup)
	}
	PrivateGroup := rt.Group("v1")
	PrivateGroup.Use(auth)
	{
		partnerR.Contract.InitRouter(PrivateGroup)
		partnerR.ContractResource.InitRouter(PrivateGroup)
	}

	return rt
}

type Login struct {
	Username string `header:"Username"`
	Password string `header:"Password"`
}

func auth(cxt *gin.Context) {
	var login Login

	if err := cxt.ShouldBindHeader(&login); err != nil {
		cxt.JSON(500, err)
	}

	requests.Logger.Infof("========> author success !!! username: %s, password: %s", login.Username, login.Password)
}
