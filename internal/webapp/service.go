package webapp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
)

const (
	portsEndpointName = "ports"
	maxPartSizeInMB   = 2
)

type WebAppService interface {
	CreatePort(ctx context.Context, port *Port) error
}

type ServiceHandler struct {
	svc        WebAppService
	httpServer *http.Server
	log        *zap.Logger
}

type Service struct {
	log *zap.Logger
}

type errorResp struct {
	Error string `json:"error_message"`
}

func NewServiceHandler(svc WebAppService, httpServer *http.Server, logger *zap.Logger) *ServiceHandler {
	return &ServiceHandler{
		svc:        svc,
		httpServer: httpServer,
		log:        logger,
	}
}

// Register connects the handlers to the router.
func (sh *ServiceHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("/"+portsEndpointName, sh.ports)
}

func (sh *ServiceHandler) Run() {
	if err := sh.httpServer.ListenAndServe(); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}

func (sh *ServiceHandler) ports(respWriter http.ResponseWriter, request *http.Request) {
	var (
		err              error
		createdPortNames []string
	)

	switch request.Method {
	case http.MethodPost:
		createdPortNames, err = sh.ingestPorts(request)
	default:
		http.Error(respWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// here could be various reasons which should be mapped to correct status code
	// custom error would be solution here - this way is just to speed up and keep it simple
	if err != nil {
		sh.renderErr(respWriter, err.Error(), http.StatusInternalServerError)
	}

	sh.renderResponse(respWriter, createdPortNames, http.StatusCreated)
}

func (sh *ServiceHandler) ingestPorts(request *http.Request) (createdPortCodes []string, err error) {
	// Get the JSON file from the request body
	err = request.ParseMultipartForm(maxPartSizeInMB << 20)
	if err != nil {
		return nil, fmt.Errorf("failed to parse multipart form: %w", err)
	}
	file, _, fileErr := request.FormFile("ports")
	if fileErr != nil {
		return nil, fmt.Errorf("failed to read part from multipart form: %w", err)
	}
	defer file.Close()

	fileReader := bufio.NewReader(file)
	decoder := json.NewDecoder(fileReader)
	createdPortCodes = make([]string, 0)

	for decoder.More() {
		port, decodeErr := decodePort(decoder)
		if decodeErr != nil {
			return createdPortCodes, decodeErr
		}
		if port != nil {
			err = sh.svc.CreatePort(request.Context(), port)
			if err != nil {
				return createdPortCodes, fmt.Errorf("failed to create port %v, errMsg=%w", port, err)
			}
			createdPortCodes = append(createdPortCodes, port.Code)
		}
	}
	return createdPortCodes, nil
}

func NewService(logger *zap.Logger) *Service {
	return &Service{log: logger}
}

func (s Service) CreatePort(ctx context.Context, port *Port) error {
	// send call via grpc to ports service to create a new port
	return nil
}

func (sh *ServiceHandler) renderErr(w http.ResponseWriter, errMsg string, status int) {
	sh.renderResponse(w, errorResp{Error: errMsg}, status)
}

func (sh *ServiceHandler) renderResponse(w http.ResponseWriter, res interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	content, err := json.Marshal(res)
	if err != nil {
		sh.log.Warn("failed to marshal response", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	if _, err = w.Write(content); err != nil {
		sh.log.Warn("failed to send response", zap.Error(err))
	}
}
