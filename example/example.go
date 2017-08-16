// example
package main

import (
	"fmt"

	"github.com/beewit/beekit/mysql"
)

var (
	DB = mysql.DB
)

func main() {
	results, err := DB.Query("show tables")
	if err != nil {
		fmt.Errorf("error msg:", err)
	}
	for _, result := range results {
		for k, v := range result {
			fmt.Printf("%s -> %s\n", k, v)
		}
	}
}
