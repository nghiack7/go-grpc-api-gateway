package logger

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type level string

const (
	Info  level = "info"
	Debug       = "debug"
	Error       = "error"
	Warn        = "warn"
)

type LoggerRequest struct {
	Message string `json:"message"`
	Caller  string `json:"caller"`
	Level   level  `json:"level"`
}

func WriteLogger(c *gin.Context, logger Logger) {
	body := LoggerRequest{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	switch body.Level {
	case Info:
		logger.Infof(body.Message, body.Caller)
	case Error:
		logger.Errorf(body.Message, body.Caller)
	case Warn:
		logger.Warnf(body.Message, body.Caller)
	case Debug:
		logger.Debugf(body.Message, body.Caller)
	default:
		logger.Infof(body.Message, body.Caller)
	}
}
