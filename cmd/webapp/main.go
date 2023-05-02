package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"

	"github.com/arturskrzydlo/ports/internal/common/grpc"

	"github.com/arturskrzydlo/ports/internal/common/pb"

	"github.com/arturskrzydlo/ports/internal/webapp"
)

type appConfig struct {
	LogLevel          string `env:"LOG_LEVEL" envDefault:"INFO"`
	ServerAddress     string `env:"SERV_ADDRESS" envDefault:"0.0.0.0:8080"`
	ReadTimeout       int    `env:"READ_TIMEOUT_IN_SEC" envDefault:"5"`
	ReadHeaderTimeout int    `env:"READ_HEADER_TIMEOUT_IN_SEC" envDefault:"5"`
	WriteTimeout      int    `env:"WRITE_TIMEOUT_IN_SEC" envDefault:"5"`
	IdleTimeout       int    `env:"IDLE_TIMEOUT_IN_SEC" envDefault:"5"`

	PortsGRPServerAddress  string `env:"PORTS_GRPC_ADDRESS" envDefault:"0.0.0.0:8090"`
	GRPCKeepAliveInSeconds int    `eng:"GRPC_KEEP_ALIVE_IN_SECONDS" envDefault:"60"`
}

// ParseConfig parses a struct containing `env` tags and loads its values from
// environment variables.
func ParseConfig(cfg interface{}) error {
	if err := env.Parse(cfg); err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}
	return nil
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

	conn, err := grpc.NewClientConnectionContext(ctx, cfg.PortsGRPServerAddress,
		cfg.GRPCKeepAliveInSeconds)
	if err != nil {
		log.Error("error while creating a gRPC connection to ports service", zap.Error(err))
		return
	}
	defer conn.Close()

	service := webapp.NewService(log, pb.NewPortServiceClient(conn))

	mux := http.NewServeMux()
	srv := &http.Server{
		Handler:           mux,
		Addr:              cfg.ServerAddress,
		ReadTimeout:       time.Duration(cfg.ReadTimeout) * time.Second,
		ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeout) * time.Second,
		WriteTimeout:      time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:       time.Duration(cfg.IdleTimeout) * time.Second,
	}

	go cancelOnSignal(cancel, signalCh, log, srv)

	handler := webapp.NewServiceHandler(service, srv, log)
	handler.Register(mux)

	log.Info("Successfully started webapp service")
	handler.Run()
}

func cancelOnSignal(cancel context.CancelFunc, ch chan os.Signal, log *zap.Logger, server *http.Server) {
	sig := <-ch
	log.Info("Shutting down application on signal", zap.String("signal", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Error("failed to shutdown a server", zap.Error(err))
	}
	cancel()
}
