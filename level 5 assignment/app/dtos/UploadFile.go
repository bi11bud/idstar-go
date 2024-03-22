package dtos

import "mime/multipart"

type UploadFile struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type UploadResponse struct {
	FileName        string
	FileDownloadUrl string
	FileType        string
	Size            string
}
