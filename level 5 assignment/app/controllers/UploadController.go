package controllers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	dtos "idstar.com/app/dtos"
	resp "idstar.com/app/dtos/response"
)

// type UploadController struct {
// 	service *services.UploadService
// }

// func NewUploadController(service *services.UploadService) *UploadController {
// 	return &UploadController{
// 		service: service,
// 	}
// }

type UploadController struct{}

func NewUploadController() *UploadController {
	return &UploadController{}
}

// UploadFile godoc
// @Summary Uploads a file
// @Description Uploads a file and saves it
// @Tags Files
// @Accept mpfd
// @Produce json
// @Param file formData file true "File to upload"
// @Router /upload [post]
func (trl *UploadController) UploadFile(c *gin.Context) {
	uploadFile := dtos.UploadFile{}
	if err := c.ShouldBind(&uploadFile); err != nil {
		response := resp.FailedResponse{
			Code:   http.StatusBadRequest,
			Status: err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file := uploadFile.File

	uploadDir := "upload"
	dst := filepath.Join(uploadDir, file.Filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {

		response := resp.FailedResponse{
			Code:   500,
			Status: err.Error(),
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := dtos.UploadResponse{
		FileName:        file.Filename,
		FileDownloadUrl: dst,
		FileType:        file.Header.Get("Content-Type"),
		Size:            strconv.FormatInt(file.Size, 10),
	}

	resSukses := resp.Response{
		Code:   200,
		Data:   data,
		Status: "sukses",
	}

	c.JSON(http.StatusOK, resSukses)
}

// ShowFile godoc
// @Summary Show a file
// @Description Show a file
// @Tags Files
//
//	@Accept			json
//	@Produce		json
//	@Param			filename	path	string	false	"filename"
//	@Router			/showFile/{filename} [get]
func (trl *UploadController) ShowFile(c *gin.Context) {

	filename := c.Param("filename")
	// Construct the full path to the file
	fileURL := "upload/" + filename
	c.File(fileURL)

}
