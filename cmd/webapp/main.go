package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/arturskrzydlo/ports/internal/webapp"
	"github.com/caarlos0/env/v6"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type appConfig struct {
	LogLevel      string `env:"LOG_LEVEL" envDefault:"INFO"`
	ServerAddress string `env:"SERV_ADDRESS" envDefault:"0.0.0.0:8080"`
	ReadTimeout   int    `env:"READ_TIMEOUT_IN_SEC" envDefault:"5"`
	WriteTimeout  int    `env:"WRITE_TIMEOUT_IN_SEC" envDefault:"5"`
	IdleTimeout   int    `env:"IDLE_TIMEOUT_IN_SEC" envDefault:"5"`
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

	service := webapp.NewService(log)

	mux := http.NewServeMux()
	srv := &http.Server{
		Handler:      mux,
		Addr:         cfg.ServerAddress,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}

	handler := webapp.NewServiceHandler(service, srv)
	handler.Register(mux)
	handler.Run()
}