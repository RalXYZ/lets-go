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

type contentRaw struct {
	uid string
	id  string
	age string
}

func getParameters(c echo.Context) *contentRaw {
	return &contentRaw{c.QueryParam("uid"), c.QueryParam("id"), c.QueryParam("age")}
}

func main() {
	e := echo.New()

	e.POST("/create", func(c echo.Context) error {
		parameters := getParameters(c)
		if parameters.id == "" || parameters.age == "" {
			return c.JSON(http.StatusBadRequest, nil)
		} else {
			age, _ := strconv.Atoi(parameters.age)
			content := dbCreate(&bar{parameters.id, age})
			return c.JSON(http.StatusOK, content)
		}
	})

	e.GET("/retrieve", func(c echo.Context) error {
		parameters := getParameters(c)
		switch parameters.uid {
		case "":
			return c.JSON(http.StatusBadRequest, nil)
		default:
			uid, _ := strconv.ParseInt(parameters.uid, 10, 64)
			return c.JSON(dbRetrieve(uid))
		}
	})

	e.PUT("/update", func(c echo.Context) error {
		parameters := getParameters(c)
		if parameters.uid == "" || parameters.id == "" || parameters.age == "" {
			return c.JSON(http.StatusBadRequest, nil)
		} else {
			uid, _ := strconv.ParseInt(parameters.uid, 10, 64)
			age, _ := strconv.Atoi(parameters.age)
			return c.JSON(dbUpdate(&Content{uid, parameters.id, age}))
		}
	})

	e.DELETE("/delete", func(c echo.Context) error {
		parameters := getParameters(c)
		switch parameters.uid {
		case "":
			return c.JSON(http.StatusBadRequest, nil)
		default:
			uid, _ := strconv.ParseInt(parameters.uid, 10, 64)
			return c.JSON(dbDelete(uid), nil)
		}
	})

	e.Logger.Fatal(e.Start(":8080"))
}
