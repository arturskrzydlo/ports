package webapp

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/arturskrzydlo/ports/internal/pb"
)

const (
	portsEndpointName = "ports"
	maxPartSizeInMB   = 2
)

type WebAppService interface {
	CreatePort(ctx context.Context, port *Port) error
	FetchPorts(ctx context.Context) ([]*Port, error)
}

type ServiceHandler struct {
	svc        WebAppService
	httpServer *http.Server
	log        *zap.Logger
}

type Service struct {
	log         *zap.Logger
	portsClient pb.PortServiceClient
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
		err        error
		response   any
		statusCode int
	)

	switch request.Method {
	case http.MethodPost:
		response, err = sh.ingestPorts(request)
		statusCode = http.StatusCreated
	case http.MethodGet:
		response, err = sh.svc.FetchPorts(request.Context())
		statusCode = http.StatusOK
	default:
		http.Error(respWriter, "Method not allowed", http.StatusMethodNotAllowed)
	}

	// here could be various reasons which should be mapped to correct status code
	// custom error would be solution here - this way is just to speed up and keep it simple
	if err != nil {
		sh.renderErr(respWriter, err.Error(), http.StatusInternalServerError)
	}

	sh.renderResponse(respWriter, response, statusCode)
}

func (sh *ServiceHandler) ingestPorts(request *http.Request) (createdPortIDs []string, err error) {
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
	createdPortIDs = make([]string, 0)

	for decoder.More() {
		port, decodeErr := decodePort(decoder)
		if decodeErr != nil {
			return createdPortIDs, decodeErr
		}
		if port != nil {
			err = sh.svc.CreatePort(request.Context(), port)
			if err != nil {
				return createdPortIDs, fmt.Errorf("failed to create port %v, errMsg=%w", port, err)
			}
			createdPortIDs = append(createdPortIDs, port.ID)
		}
	}
	return createdPortIDs, nil
}

func (sh *ServiceHandler) fetchPorts() {
}

func NewService(logger *zap.Logger, portsClient pb.PortServiceClient) *Service {
	return &Service{log: logger, portsClient: portsClient}
}

func (s Service) CreatePort(ctx context.Context, port *Port) error {
	_, err := s.portsClient.CreatePort(ctx, &pb.CreatePortRequest{Port: portToPB(port)})
	if err != nil {
		return fmt.Errorf("failed to Create ports in Ports service:%w", err)
	}
	return nil
}

func (s Service) FetchPorts(ctx context.Context) ([]*Port, error) {
	getPortsResponse, err := s.portsClient.GetPorts(ctx, new(emptypb.Empty))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ports from Ports service:%w", err)
	}

	allPorts := make([]*Port, len(getPortsResponse.Ports))
	for i, portPb := range getPortsResponse.Ports {
		allPorts[i] = pbToPort(portPb)
	}
	return allPorts, nil
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
