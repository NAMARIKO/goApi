package main

import (
	"github.com/labstack/echo"
	"router"
)

func main()  {
	e := echo.New()
	router.GetRouter(e)
}