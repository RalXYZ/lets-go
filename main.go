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

	e.GET("/create", func(c echo.Context) error {
		id := c.QueryParam("id")
		age, _ := strconv.Atoi(c.QueryParam("age"))
		content := dbCreate(&bar{id, age})
		return c.JSON(http.StatusOK, content)
	})

	e.GET("/retrieve", func(c echo.Context) error {
		uid, _ := strconv.ParseInt(c.QueryParam("uid"), 10, 64)
		content := dbRetrieve(uid)
		return c.JSON(http.StatusOK, content)
	})

	e.GET("/update", func(c echo.Context) error {
		uid, _ := strconv.ParseInt(c.QueryParam("uid"), 10, 64)
		id := c.QueryParam("id")
		age, _ := strconv.Atoi(c.QueryParam("age"))
		content := dbUpdate(&Content{uid, id, age})
		return c.JSON(http.StatusOK, content)
	})

	e.GET("/delete", func(c echo.Context) error {
		uid, _ := strconv.ParseInt(c.QueryParam("uid"), 10, 64)
		dbDelete(uid)
		return c.String(http.StatusOK, "Delete operation succeeded")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
