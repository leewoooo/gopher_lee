package api

import (
	"bytes"
	"encoding/json"
	"gopher_lee/file/internal/dto"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSingleFileUpload_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	//givne
	uploadFilePath := "../../../tmp/example.jpg"

	//when
	oFile, err := os.Open(uploadFilePath)
	assert.NoError(err)
	defer oFile.Close()

	//	create MultiForm
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(uploadFilePath))
	assert.NoError(err)

	_, err = io.Copy(multi, oFile)
	assert.NoError(err)
	writer.Close()

	// 	create req, resp
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/upload", buf)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	//	mkdir
	err = os.MkdirAll("./tmp", 0777)
	assert.NoError(err)
	defer os.RemoveAll("./tmp")

	handler := NewUpload(logrus.NewEntry(logrus.New()))
	handler.Single(resp, req)

	//then
	assert.Equal(http.StatusOK, resp.Code)
	response := &dto.UploadResponse{}

	err = json.NewDecoder(resp.Body).Decode(response)
	assert.NoError(err)

	assert.Empty(response.Error)
	assert.Equal(1, len(response.Names))
	assert.Equal("example.jpg", response.Names[0])

	// 	compare file
	list, err := os.ReadDir("./tmp")
	assert.NoError(err)

	uFile, err := os.Open(filepath.Join("./tmp", list[0].Name()))
	assert.NoError(err)

	oBuf := []byte{}
	_, err = oFile.Read(oBuf)
	assert.NoError(err)

	uBuf := []byte{}
	_, err = uFile.Read(uBuf)
	assert.NoError(err)

	assert.Equal(oBuf, uBuf)
}
