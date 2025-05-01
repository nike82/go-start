package intermediate

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	specificTime := time.Date(2024, time.December, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("specific time:", specificTime)

	parsTime, _ := time.Parse("2006-01-02", "2020-05-01")
	parsTime1, _ := time.Parse("06-01-02", "20-05-01")
	parsTime2, _ := time.Parse("06-1-2", "20-5-1")
	parsTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 22-04")
	fmt.Println(parsTime)
	fmt.Println(parsTime1)
	fmt.Println(parsTime2)
	fmt.Println(parsTime3)

	t := time.Now()
	fmt.Println("formated time:", t.Format("Monday 06-02-03 15-04-05"))

	oneDayLater := t.Add(time.Hour * 24)
	fmt.Println(oneDayLater)
	fmt.Println(oneDayLater.Weekday())

	fmt.Println("rounded time:", t.Round(time.Hour))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	t = time.Date(2024, time.April, 8, 14, 16, 40, 00, time.UTC)
	tLocal := t.In(loc)
	roundedTime := t.Round(time.Hour)
	roundedTimeLocal := roundedTime.In(loc)
	fmt.Println("original time utc:", t)
	fmt.Println("original time local:", tLocal)
	fmt.Println("rounded time utc:", roundedTime)
	fmt.Println("rounded time local:", roundedTimeLocal)

	t = time.Now()
	fmt.Println("trunkated time:", t.Truncate(time.Hour))

	location, _ := time.LoadLocation("America/New_York")
	tInNY := time.Now().In(location)
	fmt.Println("New-York time:", tInNY)

	t1 := time.Date(2025, time.August, 4, 12, 0, 0, 0, time.UTC)
	t2 := time.Date(2025, time.August, 4, 18, 0, 0, 0, time.UTC)
	duration:=t2.Sub(t1)
	fmt.Println(duration)

	fmt.Println("t2 after t1:", t2.After(t1))
}
