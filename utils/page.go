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
	Order     string
	Groupby   string
}

type PageData struct {
	PageIndex  int
	PageSize   int
	PageNumber int
	Count      int
	Data       []map[string]interface{}
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
	if pageSize > 100 {
		//防止大量数据查询
		pageSize = 100
	}
	return pageSize
}
