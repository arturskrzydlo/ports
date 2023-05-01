package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/arturskrzydlo/ports/internal/pb"

	"github.com/arturskrzydlo/ports/internal/webapp"
)

type appConfig struct {
	LogLevel          string `env:"LOG_LEVEL" envDefault:"INFO"`
	ServerAddress     string `env:"SERV_ADDRESS" envDefault:"0.0.0.0:8080"`
	ReadTimeout       int    `env:"READ_TIMEOUT_IN_SEC" envDefault:"5"`
	ReadHeaderTimeout int    `env:"READ_HEADER_TIMEOUT_IN_SEC" envDefault:"5"`
	WriteTimeout      int    `env:"WRITE_TIMEOUT_IN_SEC" envDefault:"5"`
	IdleTimeout       int    `env:"IDLE_TIMEOUT_IN_SEC" envDefault:"5"`

	PortsGRPServerAddress string `env:"PORTS_GRPC_ADDRESS" envDefault:"0.0.0.0:8090"`
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

	zapCfg := zap.NewProductionConfig()
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapCfg.EncoderConfig),
		zapcore.AddSync(io.Discard),
		zapcore.InfoLevel,
	)
	log := zap.New(core)

	conn, err := webapp.NewClientConnectionContext(context.Background(), cfg.PortsGRPServerAddress)
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

	handler := webapp.NewServiceHandler(service, srv, log)
	handler.Register(mux)
	handler.Run()
}
