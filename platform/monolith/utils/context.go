package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
)

type AppContext struct{
	Router *gin.Engine
	Logger log.Logger
}
