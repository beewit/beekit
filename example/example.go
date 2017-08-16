// example
package main

import (
	"fmt"

	"github.com/beewit/beekit/mysql"
	"github.com/beewit/beekit/redis"
)

var (
	db = mysql.DB
	rd = redis.Cache
)

func main() {
	results, err := db.Query("show tables")
	if err != nil {
		fmt.Errorf("error msg:", err)
	}
	for _, result := range results {
		for k, v := range result {
			fmt.Printf("%s -> %s\n", k, v)
		}
	}

	setStrResult, setStrErr := rd.SetString("testKey", "testValue")

	fmt.Println(setStrResult)
	if setStrErr != nil {
		fmt.Println(setStrErr)
	}

	value, getStrErr := rd.GetString("testKey")

	if getStrErr != nil {
		fmt.Println(getStrErr)
	}
	fmt.Println(value)
}

func check(err error) {
	if err != nil {
		fmt.Errorf("err msg:\n", err)
	}
}
