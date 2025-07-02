package srv

import (
	"fmt"

	"github.com/watchman1989/steez/comm"

	"github.com/gin-gonic/gin"
)

func SrvStart() {
	r := gin.New()
	RegisterRouter(r)

	go func() {
		err := r.Run(fmt.Sprintf(":%s", "8080"))
		if err != nil {
			//
			fmt.Printf("http server run error: %s", err.Error())
		}
	}()

	<-comm.Quit
}
