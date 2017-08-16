package conf

import (
	"fmt"
	"testing"
)

func Test_Config(t *testing.T) {
	config := New("config.json")
	fmt.Println(config)

	fmt.Println(config.Get("mysql.host"))
	fmt.Println(config.Get("mysql.database"))
	/*d := config.Get("domain")
	fmt.Println(d)


	host := config.Get("mysql.host")
	if d == nil {
		t.Error("domain is null")
	}

	if host == nil {
		t.Error("host is null")
	}*/
}
