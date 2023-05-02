package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"

	"github.com/arturskrzydlo/ports/internal/common/grpc"

	"github.com/arturskrzydlo/ports/internal/common/pb"

	"github.com/arturskrzydlo/ports/internal/ports/adapters"

	"github.com/arturskrzydlo/ports/internal/ports"
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

	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer log.Sync()

	ctx, cancel := context.WithCancel(context.Background())
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	grpcServer, err := grpc.NewServer(cfg.GRPCServerAddress, log)
	if err != nil {
		log.Error("error while creating a gRPC connection to ports service", zap.Error(err))
		return
	}

	pb.RegisterPortServiceServer(grpcServer, ports.NewPortsService(log, adapters.NewInMemoryRepo(log)))

	go func() {
		if err = grpcServer.Run(ctx); err != nil {
			log.Error("ports server experienced run error", zap.Error(err))
			return
		}
	}()

	// can be adde
	log.Info("Successfully started ports service")

	cancelOnSignal(cancel, signalCh, log)
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

func cancelOnSignal(cancel context.CancelFunc, ch chan os.Signal, log *zap.Logger) {
	sig := <-ch
	log.Info("Shutting down application on signal", zap.String("signal", sig.String()))
	cancel()
}
