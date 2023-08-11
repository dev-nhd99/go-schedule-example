package helper

import (
	// "fmt"
	// "time"

	"github.com/go-co-op/gocron"
)

func NewScheduler() {
	
}
func CreateSchedule(s *gocron.Scheduler , num int) (error) {
	for i:=0 ; i < num ; i ++ {
		_, err := s.Every(5).Hour().Tag("Health").Do(func(){
			// fmt.Println("This is shedule !!!!!", time.Now())
			// GetHealth()
			})
		
			if err != nil {
				return err
			}
		
	}
	return nil
}

