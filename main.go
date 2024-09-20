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
			"result": result,
		})
	})
	e.Start(":8080")
}
