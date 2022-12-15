package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/auth"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/config"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/container"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/container/logFactory"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/container/logger"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/order"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/product"
	"github.com/pkg/errors"
)

var logConfig config.LogConfig

func init() {
	logConfig = config.LogConfig{
		Code:         "zap",
		Level:        "info",
		EnableCaller: true,
	}
}

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		logger.Log.Errorf("Can''t load Config")
	}
	err = loadLogger(logConfig)
	if err != nil {
		logger.Log.Errorf("Failed loadLogger")
	}
	logger.Log.Infof("Continue")
	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	container.RegisterRoutes(r, logger.Log)

	r.Run(c.Port)
}

func loadLogger(lc config.LogConfig) error {
	loggerType := lc.Code
	err := logFactory.GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil

}
