package resource

import (
	"encoding/json"
	"gopher_lee/file/internal/dto"
	"gopher_lee/log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const downloadFilePath = "./tmp"

// Download download file handler interface
type Download interface {
	Single(w http.ResponseWriter, r *http.Request)
}

// DownloadImpl implement Download interface
type DownloadImpl struct {
	logger log.Logger
}

// NewDownload create download interface
func NewDownload(logger log.Logger) Download {
	return &DownloadImpl{logger: logger}
}

// Single download single file handler
func (d *DownloadImpl) Single(w http.ResponseWriter, r *http.Request) {
	// TODO: get file path or file name where request

	// get file Name where reqeust URI
	fileName := strings.TrimPrefix(r.URL.RequestURI(), "/download/")
	if fileName == "" {
		fileName = "example.jpg"
	}
	filePath := filepath.Join(downloadFilePath, fileName)

	// make response body
	resp := &dto.DownloadResponse{}

	// exist check with stat
	_, err := os.Stat(filePath)
	if err != nil {
		d.logger.Warnf("could not exist file path:%s", filePath)

		w.WriteHeader(http.StatusNotFound)
		resp.Error, resp.Names = "could not exist file", nil
		json.NewEncoder(w).Encode(resp)

		return
	}

	// file open
	oFile, err := os.Open(filePath)
	if err != nil {
		d.logger.Errorf("could not open file path:%s", filePath)

		w.WriteHeader(http.StatusInternalServerError)
		resp.Error, resp.Names = "could not open file", nil
		json.NewEncoder(w).Encode(resp)

		return
	}
	defer oFile.Close()

	// get file MIME type
	buf := make([]byte, 512)
	_, err = oFile.Read(buf)
	if err != nil {
		d.logger.Errorf("could not read file path:%s", filePath)

		w.WriteHeader(http.StatusInternalServerError)
		resp.Error, resp.Names = "could not read file", nil
		json.NewEncoder(w).Encode(resp)

		return
	}
	mime := http.DetectContentType(buf)

	// set http response header
	w.Header().Add("Content-Type", mime)
	w.Header().Add("Content-Disposition", "attachment; filename="+filepath.Base(filePath))
	w.WriteHeader(http.StatusCreated)

	// download 1
	// set http response body
	http.ServeFile(w, r, filePath)

	// download 2
	// _, err = io.Copy(w, oFile)
	// if err != nil {
	// 	d.logger.Errorf("could not copy file to response writer error:%v", err)

	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	resp.Error, resp.Names = "could not read file", nil
	// 	json.NewEncoder(w).Encode(resp)

	// 	return
	// }

	// w.WriteHeader(http.StatusOK)
}
