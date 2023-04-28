package webapp

import (
	"context"
	"log"
	"net/http"

	"go.uber.org/zap"
)

type WebAppService interface {
	CreatePorts(ctx context.Context) error
}

type Service struct {
	log *zap.Logger
}

func NewService(log *zap.Logger) *Service {
	return &Service{log: log}
}

func (sh *ServiceHandler) ports(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
}

func (s Service) CreatePorts(ctx context.Context) error {
	return nil
}

type ServiceHandler struct {
	svc        WebAppService
	httpServer *http.Server
}

func NewServiceHandler(svc WebAppService, httpServer *http.Server) *ServiceHandler {
	return &ServiceHandler{
		svc:        svc,
		httpServer: httpServer,
	}
}

// Register connects the handlers to the router.
func (sh *ServiceHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/ports", sh.ports)
}

func (sh *ServiceHandler) Run() {
	if err := sh.httpServer.ListenAndServe(); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
