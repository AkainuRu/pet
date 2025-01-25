package main

import (
	h "gomodule/comfortableFuncs"
	"gomodule/databases"
	"github.com/labstack/echo/v4"
)


func main() {
	databases.InitDB()

	e := echo.New()

	e.GET("/messages", h.GetHandler)
	e.POST("/messages", h.PostHandler)
	e.PATCH("/messages/:ID", h.PatchHandler)
	e.DELETE("/messages/:ID", h.DeleteHandler)
	e.Start(":5252")
}

