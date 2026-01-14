package util

import "net/http"

type Pagination struct {
	CurrentPage int64 `json:"currentPage"`
	Limit       int64 `json:"limit"`
	TotalItems  int64 `json:"totalItems"`
	TotalPages  int64 `json:"totalPages"`
}

type PaginatedData struct {
	Data       any        `json:"data"`
	Pagination Pagination `json:"pagination"`
}

func SendPage(w http.ResponseWriter, data any, page int64, limit int64, cnt int64) {
	totalPages := cnt / limit
	if cnt%limit != 0 {
		totalPages++
	}

	paginatedData := PaginatedData{
		Data: data,
		Pagination: Pagination{
			CurrentPage: page,
			Limit:       limit,
			TotalItems:  cnt,
			TotalPages:  totalPages,
		},
	}

	SendData(w, http.StatusOK, paginatedData)
}
