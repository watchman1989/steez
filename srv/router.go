package srv

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {

	r.GET("/ping", Welcome)

}
