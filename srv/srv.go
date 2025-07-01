package srv

import (
	"fmt"
	"steez/comm"

	"github.com/gin-gonic/gin"
)

func SrvStart() {
	r := gin.New()
	RegisterRouter(r)

	go func() {
		err := r.Run(fmt.Sprintf(":%s", "8808"))
		if err != nil {
			//
			fmt.Printf("http server run error: %s", err.Error())
		}
	}()

	<-comm.Quit
}
