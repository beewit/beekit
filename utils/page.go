package utils

import (
	"github.com/beewit/beekit/utils/convert"
)

type PageTable struct {
	Fields    string
	Table     string
	Where     string
	PageIndex int
	PageSize  int
}

type PageData struct {
	PageIndex int
	PageSize  int
	Count     int64
	Data      []map[string]interface{}
}

const PAGE_SIZE = 10

func GetPageIndex(pi string) int {
	pageIndex := 1
	if pi != "" && IsValidNumber(pi) {
		pageIndex = convert.MustInt(pi)
	}
	return pageIndex
}

func GetPageSize(ps string) int {
	pageSize := PAGE_SIZE
	if ps != "" && IsValidNumber(ps) {
		pageSize = convert.MustInt(ps)
	}
	return pageSize
}
