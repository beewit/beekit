package log

import (
	"fmt"

	"github.com/astaxie/beego/logs"
	"github.com/beewit/beekit/conf"
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
			"separate": "[%s]"
		}`,
		cfg.Get("logger.path"),
		cfg.Get("logger.maxdays"),
		"true",
		"true",
		cfg.Get("logger.level"),
		cfg.Get("logger.separate"),
	)
	fmt.Println(conf)

	logs.SetLogger(logs.AdapterMultiFile, conf)
	logs.SetLogger("console")
	logs.EnableFuncCallDepth(true)
}
