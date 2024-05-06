package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	go StartUpdateTask()

	e := echo.New()

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/interfaces/:name", func(c echo.Context) error {
		interfaceName := c.Param("name")
		lock.RLock()
		defer lock.RUnlock()

		data, ok := interfacesMap[interfaceName]
		if !ok {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "interface not found"})
		}

		return c.JSON(http.StatusOK, data)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
