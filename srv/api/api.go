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
	args := QueryAccountArgs{}
	if err := c.ShouldBindJSON(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), QueryAccountTimeout)
	defer cancel()
	records, err := do.RecursiveQuery(ctx, args.AccountNo, args.Level)
	if err != nil {
		comm.GContext.Logger.Errorf("recursive query failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"records": records})
}