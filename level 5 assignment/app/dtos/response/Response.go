package dtos

type Response struct {
	Code   int    `json:"code"`
	Data   any    `json:"data"`
	Status string `json:"status"`
}

type FailedResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Content struct {
	Data             any      `json:"data"`
	Pageable         Pageable `json:"pageable"`
	Last             bool     `json:"last"`
	TotalElements    int      `json:"totalElements"`
	TotalPages       int      `json:"totalPages"`
	Size             int      `json:"size"`
	Number           int      `json:"number"`
	Sort             Sort     `json:"sort"`
	First            bool     `json:"first"`
	NumberOfElements int      `json:"numberOfElements"`
	Empty            bool     `json:"empty"`
}

type Pageable struct {
	Sort       Sort `json:"sort"`
	Offset     int  `json:"offset"`
	PageNumber int  `json:"pageNumber"`
	PageSize   int  `json:"pageSize"`
	Unpaged    bool `json:"unpaged"`
	Paged      bool `json:"paged"`
}

type Sort struct {
	Empty    bool `json:"empty"`
	Sorted   bool `json:"sorted"`
	Unsorted bool `json:"unsorted"`
}
