// example
package main

import (
	"fmt"
	"time"

	"github.com/beewit/beekit/mysql"
)

func main() {
	fmt.Println("Hello World!")
	DB := mysql.DB
	fmt.Println("i have sleep first")
	time.Sleep(1 * time.Second)
	fmt.Println("i wake up")
	results, err := DB.Query("select * from user")
	if err != nil {
		fmt.Errorf("error msg:", err)
	}
	for _, result := range results {
		for k, v := range result {
			fmt.Printf("%s -> %s\n", k, v)
		}
		fmt.Println(result)
	}
}
