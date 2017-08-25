package log

import (
	"fmt"

	"github.com/beewit/beekit/conf"
	"github.com/astaxie/beego/logs"
)

var Logger = logs.GetBeeLogger()

var (
	cfg = conf.New("config.json")
)

func init() {
	conf := fmt.Sprintf(
		`{
		  "filename": "%s",
	      "maxdays": %s,
	      "daily": %s,
		  "rotate": %s,
	      "level": %s,
	      "separate": %s
	 	}`,
		cfg.Get("logger.path"),
		cfg.Get("logger.maxdays"),
		"true",
		"true",
		cfg.Get("logger.level"),
		cfg.Get("logger.separate"),
	)

	logs.SetLogger(logs.AdapterMultiFile, conf)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
}
