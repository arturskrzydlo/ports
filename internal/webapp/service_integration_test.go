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
		// setup a request - create it from file
		requestBody := &bytes.Buffer{}
		writer := multipart.NewWriter(requestBody)
		defer writer.Close()
		part, err := writer.CreateFormFile("ports", testFilePath)
		require.NoError(t, err, "failed to read test data")

		testFile, err := os.ReadFile(testFilePath)
		require.NoError(t, err, "failed to read test data")
		_, err = part.Write(testFile)
		require.NoError(t, err, "failed to write part to file")

		// setup a server service
		log, err := zap.NewDevelopment()
		require.NoError(t, err)
		service := NewService(log)
		mux := http.NewServeMux()
		handler := NewServiceHandler(service, &http.Server{Handler: mux})
		handler.Register(mux)

		// send a request
		req := httptest.NewRequest(http.MethodPost, "/ports", requestBody)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		// webapp service call
		handler.ports(recorder, req)

		// assert that json has stored all values by requesting next call
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}
