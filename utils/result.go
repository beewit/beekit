package utils

import (
	"github.com/labstack/echo"
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/beewit/beekit/log"
)

type ResultParam struct {
	Ret  int64       `json:"ret"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	SUCCESS_CODE       = 200
	ERROR_CODE         = 400
	LOGIN_INVALID_CODE = 402
	NULL_DATA          = 404
)

func ToResultParam(b []byte) ResultParam {
	var rp ResultParam
	err := json.Unmarshal(b[:], &rp)
	if err != nil {
		log.Logger.Error(err.Error())
		return ResultParam{}
	}
	return rp
}

func Success(c echo.Context, msg string, data interface{}) error {
	return Result(c, SUCCESS_CODE, msg, data)
}

func Error(c echo.Context, msg string, data interface{}) error {
	return Result(c, ERROR_CODE, msg, data)
}

func NullData(c echo.Context) error {
	return Result(c, NULL_DATA, "暂无数据", nil)
}

func ResultApi(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, data)
}

func Result(c echo.Context, ret int64, msg string, data interface{}) error {
	resultMap := map[string]interface{}{
		"ret":  ret,
		"msg":  msg,
		"data": data,
	}
	return c.JSON(http.StatusOK, resultMap)
}

func Redirect(c echo.Context, url string) error {
	return RedirectAndAlert(c, "", url)
}

func Alert(c echo.Context, tip string) error {
	return RedirectAndAlert(c, tip, "")
}

func RedirectAndAlert(c echo.Context, tip, url string) error {
	var js string
	if tip != "" {
		js += fmt.Sprintf("alert('%v');", tip)
	}
	js += fmt.Sprintf("location.href = '%v';", url)
	return ResultHtml(c, fmt.Sprintf("<script>%v</script>", js))
}

func ResultHtml(c echo.Context, html string) error {
	return c.HTML(http.StatusOK, html)
}

func ResultString(c echo.Context, str string) error {
	return c.String(http.StatusOK, str)
}
