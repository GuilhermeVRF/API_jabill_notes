package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

type FileUploader struct{
	UploadDir string
}

func NewFilesUploader(uploadDir string) FileUploader{
	return FileUploader{
		UploadDir: uploadDir,
	}
}

func (fileUploader *FileUploader) SaveFile(file multipart.File, header *multipart.FileHeader) (string, error){
	err := os.MkdirAll(fileUploader.UploadDir, os.ModePerm)

	if err != nil{
		return "", err
	}
	
	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), header.Filename)
	filePath := filepath.Join(fileUploader.UploadDir, fileName)

	destFile, err := os.Create(filePath)
	if err != nil{
		return "", err
	}

	defer destFile.Close()
	_, err = io.Copy(destFile, file)
	if err != nil{
		return "", err
	}

	return filePath, nil
}