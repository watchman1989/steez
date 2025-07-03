package srv

import (
	"github.com/gin-gonic/gin"
	"github.com/watchman1989/steez/srv/api"
)

func RegisterRouter(r *gin.Engine) {

	r.GET("/ping", Welcome)

	r.POST("/api/query_account", api.QueryAccount)

}
