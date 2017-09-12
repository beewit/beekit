package utils

import (
	"github.com/beewit/beekit/utils/convert"
)

type PageTable struct {
	Table     string
	Where     string
	Fields    string
	PageIndex int
	PageSize  int
}

type PageData struct {
	PageIndex int
	PageSize  int
	Count     int64
	Data      []map[string]interface{}
}

func GetPageIndex(pi string) int {
	pageIndex := 1
	if pi != "" && IsValidNumber(pi) {
		pageIndex = convert.MustInt(pi)
	}
	return pageIndex
}
