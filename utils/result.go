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
	//成功
	SUCCESS_CODE = 200
	//错误
	ERROR_CODE = 400
	//授权失败
	AUTH_FAIL = 401
	//登陆无效
	LOGIN_INVALID_CODE = 402
	//失败
	FAIL_DATA = 403
	//数据为空
	NULL_DATA = 404
	//不是会员
	NOT_MEMBER = 500
	//会员续费通知
	MEMBER_RENEW = 501
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

func SuccessNullMsg(c echo.Context, data interface{}) error {
	return Result(c, SUCCESS_CODE, "", data)
}

func Error(c echo.Context, msg string, data interface{}) error {
	return Result(c, ERROR_CODE, msg, data)
}

func ErrorNull(c echo.Context, msg string) error {
	return Result(c, ERROR_CODE, msg, nil)
}

func NullData(c echo.Context) error {
	return Result(c, NULL_DATA, "暂无数据", nil)
}

func AuthFail(c echo.Context, msg string) error {
	return Result(c, AUTH_FAIL, msg, nil)
}

func AuthFailNull(c echo.Context) error {
	return Result(c, AUTH_FAIL, "未登录或登陆已失效", nil)
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
