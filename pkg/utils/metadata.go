package utils

import "math"

type MetaData struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_pages,omitempty"`
}

func CalculateMetaData(totalRecords, page, PageSize int) MetaData {
	if totalRecords == 0 {
		return MetaData{}
	}
	return MetaData{
		CurrentPage:  page,
		PageSize:     PageSize,
		FirstPage:    1,
		LastPage:     int(math.Ceil(float64(totalRecords) / float64(PageSize))),
		TotalRecords: totalRecords,
	}
}
