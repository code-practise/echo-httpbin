package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo httpbin!")
	})
	e.GET("/favicon.ico", Favicon)
	e.GET("/favicon", Favicon)
	e.GET("/get", Methods)
	e.POST("/post", Methods)
	e.PATCH("/patch", Methods)
	e.PUT("/put", Methods)
	e.DELETE("/delete", Methods)

	e.GET("/image", Image)
	e.GET("/image/jpeg", Image)
	e.GET("/image/jpg", Image)
	e.GET("/image/png", Image)
	e.GET("/image/svg", Image)
	e.GET("/image/webp", Image)

	e.GET("/xml", XMLResp)
	e.GET("/json", JSONResp)
	e.GET("/html", HTMLResp)
	e.GET("/robots.txt", RobotsTXTResp)

	// request inspect
	e.GET("/headers", RequestHeaders)
	e.GET("/ip", RequestIP)
	e.GET("/user-agent", RequestUserAgent)

	// response inspect
	e.GET("/response-headers", RequestHeaders)
	e.POST("/response-headers", RequestHeaders)

	// resposne format
	e.GET("/deny", DenyResp)
	e.Logger.Fatal(e.Start(":3000"))
}
