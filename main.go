package main

import (
	"github.com/labstack/echo/v4"
	"math/rand"
	"net/http"
	"time"
)

type ListRandomizer struct {
	List []string `json:"list"`
}

func main() {
	e := echo.New()
	e.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"date":    time.Now().Format(time.RFC3339),
			"message": "Hi, this is simple app list randomizer.",
			"version": "v0.0.1",
		})
	})
	e.POST("/list", func(c echo.Context) error {
		reqBody := new(ListRandomizer)

		if err := c.Bind(reqBody); err != nil {
			return err
		}

		var result string
		rand.NewSource(time.Now().UnixMicro())
		idxRandom := rand.Intn(len(reqBody.List))
		result = reqBody.List[idxRandom]
		return c.JSON(http.StatusOK, map[string]string{
			"version": "v1.0.0",
			"date":    time.Now().Format(time.RFC3339),
			"result":  result,
		})
	})
	e.Start(":8080")
}
