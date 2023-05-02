package main

import (
	"context"
	"fmt"
	"io"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/arturskrzydlo/ports/internal/ports"

	"github.com/arturskrzydlo/ports/internal/pb"
)

type appConfig struct {
	LogLevel          string `env:"LOG_LEVEL" envDefault:"INFO"`
	GRPCServerAddress string `env:"GRPC_SERV_ADDRESS" envDefault:"0.0.0.0:8090"`
}

func main() {
	var cfg appConfig
	err := ParseConfig(&cfg)
	if err != nil {
		panic("failed to parse app config: " + err.Error())
	}

	zapCfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapCfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	log := zap.New(core)

	grpcServer, err := ports.NewServer(cfg.GRPCServerAddress, log)
	if err != nil {
		log.Error("error while creating a gRPC connection to ports service", zap.Error(err))
		return
	}

	pb.RegisterPortServiceServer(grpcServer, ports.NewPortsService(log))
	if err = grpcServer.Run(context.Background()); err != nil {
		log.Error("ports server experienced run error", zap.Error(err))
		return
	}
}

// ParseConfig parses a struct containing `env` tags and loads its values from
// environment variables.
// TODO: move to common code
func ParseConfig(cfg interface{}) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	return nil
}
