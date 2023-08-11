package main

import (
	"fmt"
	_ "log"
	"runtime"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/go-schedule-example/cmd/pkg/utl/helper"

	"net/http"

	"github.com/labstack/echo/v4"
)



func main() {
	// task := func(){
	// 	fmt.Println("This is shedule UTC at: ",time.Now().UTC())
	// 	}

	// killJob := func(s *gocron.Scheduler, medId string){
	// 		fmt.Println("Kill >>>")
	// 		s.RemoveByTags(medId)
	// 	}
	fmt.Println("Now >>>", time.Now().UTC())
	// startAt := time.Unix(1691570940,0).UTC()
	medID := "269a4e18-e1cc-4bd2-bc0e-ca7e1b079909"
	startAt := time.Date(2023, 8, 31, 8, 42, 0, 0,time.UTC)
	// startAt1 := time.Date(2023, 8, 10, 10, 10, 0, 0,time.UTC)
	fmt.Println("StartAt:", startAt)

	s := gocron.NewScheduler(time.UTC)
	s.StartAsync()


	// _, err := s.Every(1).Hour().StartAt(startAt).Do(func(){
	// 	fmt.Println("This is shedule UTC !!!!!")
	// 	})

	// job, err := s.Every(1).Days().StartAt(startAt).Tag(medID).Do(task)
	// job1, err := s.Every(1).Days().StartAt(startAt1).Tag(medID).Do(task)
	// fmt.Println("Next>>>",job.NextRun())
	// fmt.Println("Next>>>",job1.NextRun())
	// j,_ := s.FindJobsByTag(medID)
	// fmt.Println("Jobs >>>>",j[0].NextRun())

	// job, err := s.Every(7).Day().StartAt(startAt).Tag(medID).Do(task)
	// fmt.Println("Next time >>>", job.NextRun())
	// js,_ := s.FindJobsByTag(medID)
	// fmt.Println("Jobs>>>>>",js)
	// s.RemoveByTags(medID)

	// job, err := s.Every(1).Month(12).StartAt(startAt).Tag(medID).Do(task)
	// fmt.Println("Next time >>>", job.NextRun())

	// job, err := s.Every(1).MonthLastDay().At("8:02;8:01").Tag(medID).Do(task)
	// fmt.Println("Next time >>>", job.NextRun())
	// s.RemoveByTags(medID+"MN")

	// job, err := s.Cron("0 0 1 * *").Tag(medID).Do(task)
	// fmt.Println("Next time >>>", job.NextRun())
	
	// job, err := s.Every(1).Minute().StartAt(startAt).Tag(medID).Do(killJob, s, medID)

	// if err != nil {
	// 	fmt.Println("err >>>>>>>>", err)
	// }
	
	//api
	e := echo.New()
    // e.GET("/next", func(c echo.Context) error {
    //     return c.String(http.StatusOK, job.NextRun().String())
    // })
	e.GET("/jobs", func(c echo.Context) error {
		jobs, _ := s.FindJobsByTag(medID)
		fmt.Println("Jobs >>>>> ",jobs)
        return c.String(http.StatusOK, "Oke")
    })
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