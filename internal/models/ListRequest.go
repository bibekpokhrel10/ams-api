package models

import (
	"strings"
	"time"
)

type ListRequest struct {
	Page          uint64 `form:"page"`
	Size          uint64 `form:"size"`
	SortColumn    string `form:"sort_column"`
	SortDirection string `form:"sort_direction"`
	Query         string `form:"query"`
	StartDate     string `form:"start_date"`
	EndDate       string `form:"end_date"`
}

func (req *ListRequest) Prepare() {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 20
	}
	req.SortColumn = strings.TrimSpace(req.SortColumn)
	req.SortDirection = strings.TrimSpace(req.SortDirection)
	if req.SortColumn == "" {
		req.SortColumn = "created_at"
	}
	if req.SortDirection == "" {
		req.SortDirection = "desc"
	}
	endDate := time.Now()
	startDate := endDate.AddDate(0, -1, 0)
	if req.StartDate != "" {
		startDate, _ = time.Parse("2006-01-02", req.StartDate)
	}
	if req.EndDate != "" {
		endDate, _ = time.Parse("2006-01-02", req.EndDate)
	}
	req.StartDate = startDate.Format("2006-01-02")
	req.EndDate = endDate.Format("2006-01-02")
}
