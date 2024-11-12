package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! (Version 2.1!)")
	})

	// The path "/test" should make an HTTP request to a url passed in by an environment variable, then output the contents.
	// The environment variable should be named "TEST_URL" and should be a valid URL.
	e.GET("/test", func(c echo.Context) error {
		url := os.Getenv("TEST_URL")
		if url == "" {
			return c.String(http.StatusInternalServerError, "TEST_URL environment variable not set")
		}
		resp, err := http.Get(url)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		defer resp.Body.Close()
		return c.String(http.StatusOK, resp.Status)
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/unhealthy", func(c echo.Context) error {
		return c.JSON(http.StatusServiceUnavailable, struct{ Status string }{Status: "Not OK"})
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
