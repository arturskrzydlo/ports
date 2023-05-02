//go:build integration

package webapp

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	grpc2 "github.com/arturskrzydlo/ports/internal/common/grpc"

	portsgrpc "github.com/arturskrzydlo/ports/internal/common/pb"
)

func TestPortsStoring(t *testing.T) {
	testFilePath := "testdata/ports.json"

	t.Run("should store ports from valid json file", func(t *testing.T) {
		// given
		// setup a request - create it from file
		requestBody, writer := createRequestBodyFromTestFile(t, testFilePath)

		// setup a server service
		handler, conn := setupServer(t)
		defer conn.Close()

		// send a request
		req := httptest.NewRequest(http.MethodPost, "/"+portsEndpointName, requestBody)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		// when
		handler.ports(recorder, req)

		// then
		assert.Equal(t, http.StatusCreated, recorder.Code)
		// all ports from ports.json file
		expectedResponse := `["AEAJM","AEAUH"]`
		assert.Equal(t, expectedResponse, recorder.Body.String())

		// assert that json has stored all values by requesting next call
		recorder = httptest.NewRecorder()
		req = httptest.NewRequest(http.MethodGet, "/ports", nil)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		handler.ports(recorder, req)

		var ports []*Port
		err := json.NewDecoder(recorder.Body).Decode(&ports)
		require.NoError(t, err)

		// TODO: this can be validated in much better way, to compare all the fields
		// and not rely on hardcoded strings but rather values from json file
		expectedPortIDs := []string{"AEAJM", "AEAUH"}
		counter := 0
		for _, expPortID := range expectedPortIDs {
			for _, port := range ports {
				if port.ID == expPortID {
					counter++
				}
			}
		}
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Equal(t, len(expectedPortIDs), counter)
	})
}

func createRequestBodyFromTestFile(t *testing.T, testFilePath string) (*bytes.Buffer, *multipart.Writer) {
	t.Helper()
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	part, err := writer.CreateFormFile("ports", testFilePath)
	require.NoError(t, err, "failed to read test data")

	testFile, err := os.ReadFile(testFilePath)
	require.NoError(t, err, "failed to read test data")
	_, err = part.Write(testFile)
	require.NoError(t, err, "failed to write part to file")
	writer.Close()
	return requestBody, writer
}

func setupServer(t *testing.T) (sh *ServiceHandler, conn *grpc.ClientConn) {
	t.Helper()
	// TODO: get proper ports address
	conn, err := grpc2.NewClientConnectionContext(context.Background(), ":8090", 60)
	require.NoError(t, err)
	log, err := zap.NewDevelopment()
	require.NoError(t, err)
	service := NewService(log, portsgrpc.NewPortServiceClient(conn))
	mux := http.NewServeMux()
	handler := NewServiceHandler(service, &http.Server{Handler: mux}, log)
	handler.Register(mux)
	return handler, conn
}
