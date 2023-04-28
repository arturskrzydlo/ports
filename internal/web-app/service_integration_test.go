package web_app

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

		// send a request
		req := httptest.NewRequest("POST", "/ports", requestBody)
		req.Header.Set("Content-Type", writer.FormDataContentType())
		recorder := httptest.NewRecorder()

		// web-app service call

		// assert that json has stored all values by requesting next call
		assert.Equal(t, http.StatusCreated, recorder.Code)
	})
}
