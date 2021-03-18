package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Methods(c echo.Context) error {
	respMap := map[string]interface{}{}
	respMap["args"] = c.QueryParams()
	respMap["headers"] = c.Request().Header
	respMap["origin"] = c.RealIP()
	respMap["url"] = c.Request().RequestURI
	if c.Request().Method != "GET" {
		respMap["data"] = ""
		respMap["files"] = []string{}
		respMap["form"] = map[string]string{}
		respMap["json"] = nil
	}
	return c.JSON(http.StatusOK, respMap)
}

func Image(c echo.Context) error {
	fileP := "static/img.jpeg"
	switch c.Request().URL.Path {
	case "/image":
		accept := c.Request().Header.Get("accept")
		if accept == "" {
			fileP = "static/img.jpeg"
		} else {
			for _, item := range strings.Split(accept, ",") {
				kv := strings.Split(item, "/")
				if len(kv) != 2 {
					fileP = "static/img.jpeg"
				}
				k, v := kv[0], kv[1]
				if k != "image" {
					continue
				}
				if strings.Contains(v, "png") || strings.Contains(v, "apng") {
					fileP = "static/img.png"
					break
				} else if strings.Contains(v, "webp") {
					fileP = "static/img.webp"
					break
				}
			}
		}

	case "/image/jpg", "/image/jpeg":
		fileP = "static/img.jpeg"
	case "/image/png":
		fileP = "static/img.png"
	case "/image/svg":
		// c.Response().Header().Set("Content-Type", "image/svg")
		fileP = "static/img.svg"
	case "/image/webp":
		// c.Response().Header().Set("Content-Type", "image/webp")
		fileP = "static/img.webp"
	}
	return c.File(fileP)
}

func HTMLResp(c echo.Context) error {
	return c.HTML(http.StatusOK, HTMLContent)
}

func JSONResp(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/json")
	return c.String(http.StatusOK, JSONContent)
}

func XMLResp(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "application/xml")
	return c.String(http.StatusOK, XMLContent)
}

func RobotsTXTResp(c echo.Context) error {
	return c.String(http.StatusOK, RobotsTxt)
}

func RequestIP(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"origin": c.RealIP()})
}

func RequestHeaders(c echo.Context) error {
	headersMap := map[string]string{}
	for k, v := range c.Request().Header {
		headersMap[k] = v[0]
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"headers": headersMap})
}

func RequestUserAgent(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"User-Agent": c.Request().UserAgent()})
}

func ResponseHeaders(c echo.Context) error {
	headersMap := map[string]string{}
	// 有问题
	for k, v := range c.Response().Header() {
		headersMap[k] = v[0]
	}
	return c.JSON(http.StatusOK, headersMap)
}

func DenyResp(c echo.Context) error {
	return c.String(http.StatusOK, DenyText)
}
