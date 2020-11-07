package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Content struct {
	Uid int64  `json:"uid"`
	Id  string `json:"id"`
	Age int    `json:"age"`
}

func main() {
	e := echo.New()

	e.POST("/create", func(context echo.Context) error {
		id := context.QueryParam("id")
		age, _ := strconv.Atoi(context.QueryParam("age"))
		content := dbCreate(&bar{id, age})
		return context.JSON(http.StatusOK, content)
	})

	e.GET("/retrieve", func(context echo.Context) error {
		uid, _ := strconv.ParseInt(context.QueryParam("uid"), 10, 64)
		content := dbRetrieve(uid)
		return context.JSON(http.StatusOK, content)
	})

	e.PUT("/update", func(context echo.Context) error {
		uid, _ := strconv.ParseInt(context.QueryParam("uid"), 10, 64)
		id := context.QueryParam("id")
		age, _ := strconv.Atoi(context.QueryParam("age"))
		content := dbUpdate(&Content{uid, id, age})
		return context.JSON(http.StatusOK, content)
	})

	e.DELETE("/delete", func(context echo.Context) error {
		uid, _ := strconv.ParseInt(context.QueryParam("uid"), 10, 64)
		dbDelete(uid)
		return context.String(http.StatusOK, "Delete operation succeeded")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
