package utils

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)

func Success(c echo.Context, msg string, data interface{}) error {
	return Result(c, 200, msg, data)
}

func Error(c echo.Context, msg string, data interface{}) error {
	return Result(c, 400, msg, data)
}

func Result(c echo.Context, ret int64, msg string, data interface{}) error {
	resultMap := map[string]interface{}{
		"ret":  ret,
		"msg":  msg,
		"data": data,
	}
	result, _ := json.Marshal(resultMap)
	return c.String(http.StatusOK, string(result))
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
		js += fmt.Sprintf("alert('%s');", tip)
	} else {
		js += fmt.Sprintf("location.href = '%s';", url)
	}
	return ResultHtml(c, fmt.Sprintf("<script>%s</script>", js))
}

func ResultHtml(c echo.Context, html string) error {
	return c.HTML(http.StatusOK, html)
}
