package services

import (
	"io"
	"log"
	"net/http"
	"os"
)

type UploadService struct{}

func NewUploadService() {}

func UploadFile(response http.ResponseWriter, request *http.Request) error {
	request.ParseMultipartForm(10 * 1024 * 1024)

	file, handler, err := request.FormFile("myfile")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer file.Close()

	filename := handler.Filename
	// dir := ""

	log.Println("filename: ", filename)
	log.Println("size: ", handler.Size)

	// fileLocation := filepath.Join(dir, "files", filename)
	// targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	// if err != nil {
	// 	log.Println(err)
	// 	return err
	// }
	// defer targetFile.Close()

	tempFile, err := os.CreateTemp("uploads", "upload-*.jpg")
	if err != nil {
		log.Println(err)
		return err
	}

	defer tempFile.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		log.Println(err)
		return err
	}

	tempFile.Write(fileBytes)

	return nil

}
