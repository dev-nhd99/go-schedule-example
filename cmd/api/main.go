package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-schedule-example/cmd/pkg/utl/helper"

	"net/http"

	redislock "github.com/go-co-op/gocron-redis-lock"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)



func main() {
	redisOptions := &redis.Options{
		Addr: "192.168.0.254:31000",
	}
	redisClient := redis.NewClient(redisOptions)
	locker, err := redislock.NewRedisLocker(redisClient, redislock.WithTries(1))
	if err != nil {
		fmt.Println("Err >>>", err)
	}
	s := gocron.NewScheduler(time.UTC)
	s.WithDistributedLocker(locker)
	s.StartAsync()

	//api
	e := echo.New()
   
	e.GET("/create/:num", func(c echo.Context) error {
		num := c.Param("num")
		jobs, _ := strconv.Atoi(num)
		err := helper.CreateSchedule(s, jobs)
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