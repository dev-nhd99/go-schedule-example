package helper

// import (
// 	"context"
// 	"fmt"
// 	"testing"
// 	"time"

// 	"github.com/go-co-op/gocron"
// 	"github.com/stretchr/testify/require"
// 	"go.chromium.org/luci/common/clock"
// )
  
//   type ClockTime struct {
// 	  clock.Clock
//   }
  
//   var _ gocron.TimeWrapper = ClockTime{}
  
//   func (ct ClockTime) Now(loc *time.Location) time.Time {
// 	  return ct.Clock.Now().In(loc)
//   }
  
//   func (ClockTime) Unix(sec, nsec int64) time.Time {
// 	  return time.Unix(sec, nsec)
//   }
  
//   func (ct ClockTime) Sleep(d time.Duration) {
// 	  ct.Clock.Sleep(context.Background(), d)
//   }
  
//   func TestSchedulerMock(t *testing.T) {
// 	  // A clock that runs 60x faster than real clock.
// 	  clock := testclock.NewFastClock(time.Now(), 60)
// 	  defer clock.Close()
  
// 	  fmt.Printf("clock set at %s\n", clock.Now().UTC().Format(time.RFC3339))
  
// 	  s := gocron.NewScheduler(time.UTC)
// 	  s.CustomTime(ClockTime{clock})
  
// 	  var count int
// 	  _, err := s.Every(2).Second().Do(func() {
// 		  fmt.Printf("Triggered: %v (real time=%v)\n", clock.Now().UTC().Format(time.RFC3339), time.Now().UTC().Format(time.RFC3339))
// 		  count++
// 	  })
// 	  require.NoError(t, err)
// 	  s.StartAsync()
  
// 	  start := time.Now()
// 	  for {
// 		  if count == 10 {
// 			  break
// 		  }
// 	  }
// 	  fmt.Printf("done in %v\n", time.Since(start))
//   }