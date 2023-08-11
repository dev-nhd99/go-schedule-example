package main

import (
	"runtime"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-schedule-example/cmd/pkg/utl/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)



func main() {
	
	s := gocron.NewScheduler(time.UTC)
	s.StartAsync()

	//api
	e := echo.New()
   
	e.GET("/create", func(c echo.Context) error {
		err := helper.CreateSchedule(s)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Create failed")
		}
        return c.String(http.StatusOK, "Create success")
    })

	e.GET("/system", func(c echo.Context) error {
		nums,_:= s.FindJobsByTag("Health")
		result := map[string]interface{}{
			"NumsThread": runtime.NumGoroutine(),
			"NumsJob": len(nums),
			"NumCpu": runtime.NumCPU(),
		}
		return c.JSON(http.StatusOK, result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}