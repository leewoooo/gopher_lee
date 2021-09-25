package api

import (
	"encoding/json"
	"fmt"
	"gopher_lee/file/internal/dto"
	"gopher_lee/log"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const uploadFilePath = "./tmp"

// Upload upload file handler interface
type Upload interface {
	Single(w http.ResponseWriter, r *http.Request)
}

// UploadImpl implement Upload interface
type UploadImpl struct {
	logger log.Logger
}

// NewUpload create Upload instance
func NewUpload(logger log.Logger) Upload {
	return &UploadImpl{logger: logger}
}

// Single handling to upload multipart file
func (u *UploadImpl) Single(w http.ResponseWriter, r *http.Request) {
	// set Max Size 2Mib
	r.Body = http.MaxBytesReader(w, r.Body, 2<<20)
	defer r.Body.Close()

	// create response
	resp := &dto.UploadResponse{}
	// get from file
	uFile, header, err := r.FormFile("upload_file")

	// loage file check
	if err != nil && err == multipart.ErrMessageTooLarge {
		u.logger.Errorf("upload file could not over 2Mib error:%v", err)

		w.WriteHeader(http.StatusBadRequest)
		resp.Error, resp.Names = "upload file could not over 2Mib", nil
		json.NewEncoder(w).Encode(resp)
		return
	}

	// get formfile error
	if err != nil {
		u.logger.Errorf("could not get formFile error:%v", err)

		w.WriteHeader(http.StatusBadRequest)
		resp.Error, resp.Names = "could not get FormFile", nil
		json.NewEncoder(w).Encode(resp)

		return
	}
	defer uFile.Close()

	// extarct file info
	oName := filepath.Base(header.Filename)
	size := header.Size

	// formfile을 할 때 얻어온 Header안에 Conternt type으로 해당 File의 MIME를 알아낼 수 있다.
	// MIME := header.Header.get("Content-Type")

	// prevent coverd file
	// save originName and newName to repository and use it
	nName := uuid.NewString() + filepath.Ext(oName)
	u.logger.Infof("file upload request with fields fileName:%s, size:%d byte", oName, size)

	// save file
	oFile, err := os.Create(fmt.Sprintf("%s/%s", uploadFilePath, nName))
	if err != nil {
		u.logger.Errorf("could not open File Name:%s, error:%v", oName, err)

		w.WriteHeader(http.StatusInternalServerError)
		resp.Error, resp.Names = "could not open File Name:"+oName, nil
		json.NewEncoder(w).Encode(resp)

		return
	}
	defer oFile.Close()

	if _, err = io.Copy(oFile, uFile); err != nil {
		u.logger.Errorf("could not create File Name:%s, error:%v", oName, err)

		w.WriteHeader(http.StatusInternalServerError)
		resp.Error, resp.Names = "could not create File Name:"+oName, nil
		json.NewEncoder(w).Encode(resp)

		return
	}

	// success
	w.WriteHeader(http.StatusOK)
	resp.Error, resp.Names = "", []string{oName}
	json.NewEncoder(w).Encode(resp)
}
