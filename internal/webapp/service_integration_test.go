package webapp

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestPortsStoring(t *testing.T) {
	testFilePath := "testdata/ports.json"

	t.Run("should store ports from valid json file", func(t *testing.T) {
		// given
		// setup a request - create it from file
		requestBody, writer := createRequestBodyFromTestFile(t, testFilePath)

		// setup a server service
		handler := setupServer(t)

		// send a request
		req := httptest.NewRequest(http.MethodPost, "/"+portsEndpointName, requestBody)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		// when
		handler.ports(recorder, req)

		// then
		assert.Equal(t, http.StatusCreated, recorder.Code)
		// all ports from ports.json file
		expectedResponse := `["52000","52001"]`
		assert.Equal(t, expectedResponse, recorder.Body.String())
		// assert that json has stored all values by requesting next call
		req = httptest.NewRequest(http.MethodGet, "/ports", nil)
		req.Header.Set("Content-Type", writer.FormDataContentType())
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

func setupServer(t *testing.T) *ServiceHandler {
	t.Helper()
	log, err := zap.NewDevelopment()
	require.NoError(t, err)
	service := NewService(log)
	mux := http.NewServeMux()
	handler := NewServiceHandler(service, &http.Server{Handler: mux}, log)
	handler.Register(mux)
	return handler
}
