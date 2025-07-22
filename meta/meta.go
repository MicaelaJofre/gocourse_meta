package meta

import (
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_count"`
	Page       int `json:"page,omitempty"`
	PerPage    int `json:"per_page,omitempty"`
	PageCount int `json:"page_count,omitempty"`
}

func New(page, perPage, totalCount int, pageLimitDef string) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pageLimitDef)
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0 
	if totalCount >= 0 {
		pageCount = (totalCount + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		TotalCount: totalCount,
		Page:    page,
		PerPage:    perPage,
		PageCount: pageCount,
	}, nil
}

func (p *Meta) Offset() int {
	return (p.Page - 1) * p.PerPage
}

func (p *Meta) Limit() int {
	return p.PerPage
}