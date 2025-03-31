package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

type AppContext struct {
	Context context.Context
	Router  *gin.Engine
	Logger  log.Logger
}
