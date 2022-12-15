package logFactory

import (
	"github.com/nghiack7/go-grpc-api-gateway/pkg/config"
	"github.com/nghiack7/go-grpc-api-gateway/pkg/container/zap"
	"github.com/pkg/errors"
)

// receiver for zap factory
type ZapFactory struct{}

// build zap logger
func (mf *ZapFactory) Build(lc *config.LogConfig) error {
	err := zap.RegisterLog(*lc)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}
