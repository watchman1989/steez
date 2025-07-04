package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/watchman1989/steez/comm"
	"github.com/watchman1989/steez/srv/do"
	"time"
)

const (
	QueryAccountTimeout = 15 * time.Second
)

type QueryAccountArgs struct {
	AccountNo string `json:"account_no"`
	Level int `json:"level"`
}

func QueryAccount(c *gin.Context) {
	var accountNo string
	var level int
	switch c.Request.Method {
	case "GET":
		accountNo = c.Query("account")
		level = comm.StrToInt(c.Query("level"))
	case "POST":
		args := QueryAccountArgs{}
		if err := c.ShouldBindJSON(&args); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		accountNo = args.AccountNo
		level = args.Level
	default:
		c.JSON(http.StatusMethodNotAllowed, gin.H{"error": "Method not allowed"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), QueryAccountTimeout)
	defer cancel()
	records, err := do.RecursiveQuery(ctx, accountNo, level)
	if err != nil {
		comm.GContext.Logger.Errorf("recursive query failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": records})
}