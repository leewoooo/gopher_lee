package resource

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

func TestFile_UPLOAD_DONWLOAD_SUCCESS(t *testing.T) {
	assert := assert.New(t)

	var uploadedFilePath string
	var uploadFileName string

	uploadFilePath := "../../../tmp/example.jpg"

	t.Run("FILE UPLOAD", func(t *testing.T) {
		//givne

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

		//	mkdir for test dir
		err = os.MkdirAll("./tmp", 0777)
		assert.NoError(err)

		// 	create handler
		handler := NewUpload(logrus.NewEntry(logrus.New()))
		handler.Single(resp, req)

		//then
		assert.Equal(http.StatusOK, resp.Code)
		response := &dto.UploadResponse{}

		//	chenck response body
		err = json.NewDecoder(resp.Body).Decode(response)
		assert.NoError(err)

		assert.Empty(response.Error)
		assert.Equal(1, len(response.Names))
		assert.Equal("example.jpg", response.Names[0])

		//	for download test
		list, err := os.ReadDir("./tmp")
		assert.NoError(err)

		uploadFileName = list[0].Name()
		uploadedFilePath = filepath.Join("./tmp", uploadFileName)

		// 	compare file
		uFile, err := os.Open(uploadedFilePath)
		assert.NoError(err)
		defer uFile.Close()

		oBuf := []byte{}
		_, err = oFile.Read(oBuf)
		assert.NoError(err)

		uBuf := []byte{}
		_, err = uFile.Read(uBuf)
		assert.NoError(err)

		assert.Equal(oBuf, uBuf)
	})

	t.Run("FILE DONWLOAD", func(t *testing.T) {
		//givne

		//when
		//	resp,req
		resp := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/download/"+uploadFileName, http.NoBody)

		//	create Handler
		handler := NewDownload(logrus.NewEntry(logrus.New()))

		// excute
		handler.Single(resp, req)

		// then
		assert.Equal(http.StatusOK, resp.Code)

		// compare
		uFile, err := os.Open(uploadFilePath)
		assert.NoError(err)
		defer uFile.Close()

		uBuf := []byte{}
		_, err = uFile.Read(uBuf)
		assert.NoError(err)

		dBuf := []byte{}
		_, err = resp.Body.Read(dBuf)
		assert.NoError(err)

		assert.Equal(uBuf, dBuf)
	})

	t.Run("FINALLY", func(t *testing.T) {
		// remove uploaded FIle
		err := os.Remove(uploadedFilePath)
		assert.NoError(err)

		// remove dir
		err = os.RemoveAll("./tmp")
		assert.NoError(err)
	})
}
