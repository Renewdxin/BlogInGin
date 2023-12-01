package service

import (
	"BloginGin/global"
	"BloginGin/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, header *multipart.FileHeader) (*FileInfo, error) {
	fileName := upload.GetFileName(header.Filename)
	uploadFilePath := upload.GetSavePath()
	dst := uploadFilePath + "/" + fileName
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	if upload.CheckSavePath(uploadFilePath) {
		err := upload.CreateSavePath(uploadFilePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory")
		}
	}
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit")
	}
	if upload.CheckPermission(uploadFilePath) {
		return nil, errors.New("insufficient file permissions")
	}
	if err := upload.SaveFile(header, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
