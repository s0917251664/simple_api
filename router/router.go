package router

import (
	"net/http"

	"simple_api/controller"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

//CustomValidator - 驗證struct
type CustomValidator struct {
	validator *validator.Validate
}

//Validate - 驗證
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Router(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	// 新增
	e.POST("/comment", controller.CreateData)
	// 取得
	e.GET("/comment/:uuid", controller.GetData)
	// 修改
	e.PUT("/comment/:uuid", controller.UpdateData)
	// 刪除
	e.DELETE("/comment/:uuid", controller.DelData)
}
