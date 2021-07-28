package partner

import (
	"github.com/gin-gonic/gin"
	"kingdee-test/requests"
	"net/http"
)

type Contract struct {
}

func (pc *Contract) InitRouter(router *gin.RouterGroup) gin.IRoutes {
	router.GET("/byId", getContractById)

	return router
}

func getContractById(cxt *gin.Context) {
	id, _ := cxt.GetQuery("id")

	requests.Logger.Infof("request params is id = " + id)
	cxt.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": "name " + id,
		"code": "code " + id,
	})
}
