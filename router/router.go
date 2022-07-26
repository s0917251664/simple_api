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
	quiz := e.Group("/quiz/v1")

	// 新增
	quiz.POST("/comment", controller.CreateData)
	// 取得
	quiz.GET("/comment/:uuid", controller.GetData)
	// 修改
	quiz.PUT("/comment/:uuid", controller.UpdateData)
	// 刪除
	quiz.DELETE("/comment/:uuid", controller.DelData)
}
