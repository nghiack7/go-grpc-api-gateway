package logFactory

import (
	"github.com/nghiack7/go-grpc-api-gateway/pkg/config"
)

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	config.ZAP: &ZapFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*config.LogConfig) error
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
