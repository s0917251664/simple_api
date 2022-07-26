package controller

import (
	"fmt"
	"net/http"
	"simple_api/module/db"
	rd "simple_api/module/rowdata"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func CreateData(c echo.Context) error {
	body := new(rd.RowData)
	// 驗證必要資料
	if err := c.Bind(body); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("%+v\n", body)
	if e := c.Validate(body); e != nil {
		fmt.Println(e)
		return c.JSON(http.StatusBadRequest, e.Error())
	}
	// 加入 uuid
	id := uuid.New()
	body.Uuid = id.String()
	// 新增資料
	if err := db.CreateRowData(body); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, body)
}
func GetData(c echo.Context) error {
	// 驗證必要資料
	uuid := strings.TrimSpace(c.Param("uuid"))
	if uuid == "" {
		return c.JSON(http.StatusBadRequest, "uuid is required")
	}
	// Find Data
	err, result := db.FindOneRowData(uuid)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Println(result)
	return c.JSON(http.StatusOK, result)
}
func UpdateData(c echo.Context) error {
	// 驗證必要資料
	uuid := strings.TrimSpace(c.Param("uuid"))
	if uuid == "" {
		return c.JSON(http.StatusBadRequest, "uuid is required")
	}
	body := new(rd.RowData)
	if err := c.Bind(body); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	fmt.Printf("%+v\n", body)
	if e := c.Validate(body); e != nil {
		fmt.Println(e)
		return c.JSON(http.StatusBadRequest, e.Error())
	}
	// Update Data
	err, result := db.UpdateData(uuid, body)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if result.MatchedCount >= 1 {
		return c.JSON(http.StatusOK, body)
	} else {
		return c.JSON(http.StatusBadRequest, "do not find this uuis .")
	}

}
func DelData(c echo.Context) error {
	// 驗證必要資料
	uuid := strings.TrimSpace(c.Param("uuid"))
	if uuid == "" {
		return c.JSON(http.StatusBadRequest, "uuid is required")
	}
	// Del Data
	err, result := db.DeleteRowData(uuid)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if result.DeletedCount >= 1 {
		return c.JSON(http.StatusOK, "success .")
	} else {
		return c.JSON(http.StatusBadRequest, "delete fail .")
	}

}
