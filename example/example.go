// example
package main

import (
	"fmt"

	_ "github.com/beewit/beekit/conf"
	"github.com/beewit/beekit/mysql"
)

func main() {
	fmt.Println("Hello World!")
	DB.query("select 1+1")
}
