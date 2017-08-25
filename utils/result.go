package utils

import (
	"encoding/json"
	"github.com/labstack/echo"
	"net/http"
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

