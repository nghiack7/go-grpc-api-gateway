package container

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/container/logger"
)

type loggerSvc struct {
	logger logger.Logger
}

func RegisterRoutes(r *gin.Engine, logger logger.Logger) {
	svc := &loggerSvc{logger: logger}
	routes := r.Group("/logger")
	routes.POST("", svc.WriteLogger)
}

func (l *loggerSvc) WriteLogger(c *gin.Context) {
	logger.WriteLogger(c, l.logger)
}
