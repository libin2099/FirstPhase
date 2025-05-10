package main

import (
	"fmt"
	"time"
)

type MyCalendar struct {
	start int
	end   int
}

func (c *MyCalendar) book(calendars *[]MyCalendar) {
	if len(*calendars) == 0 {
		*calendars = append(*calendars, *c)
	} else {
		canBook := true
		for _, v := range *calendars {
			if c.start < v.end && c.end >= v.start {
				canBook = false
				break
			}
		}

		if canBook {
			*calendars = append(*calendars, *c)
		}
	}
}

func main() {

	var calendars []MyCalendar = make([]MyCalendar, 0)
	const layout = "2006-01-02 15:04:05"
	startStr := "2025-04-10 18:03:31"
	endStr := "2025-04-18 20:03:31"

	tstart, errstart := time.Parse(layout, startStr)
	if errstart != nil {
		panic("error")
	}
	tend, errend := time.Parse(layout, endStr)
	if errend != nil {
		panic("error")
	}

	//fmt.Printf("%v - %v", reflect.TypeOf(tstart.Unix()), reflect.TypeOf(tend.Unix()))

	myCalendar := &MyCalendar{int(tstart.Unix()), int(tend.Unix())}
	myCalendar.book(&calendars)
	fmt.Println(*myCalendar)
	fmt.Println(calendars)

	startStr = "2025-04-11 18:03:31"
	endStr = "2025-04-16 20:03:31"

	tstart, errstart = time.Parse(layout, startStr)
	if errstart != nil {
		panic("error")
	}
	tend, errend = time.Parse(layout, endStr)
	if errend != nil {
		panic("error")
	}

	myCalendar = &MyCalendar{int(tstart.Unix()), int(tend.Unix())}
	myCalendar.book(&calendars)
	fmt.Println(*myCalendar)
	fmt.Println(calendars)

	startStr = "2025-04-01 18:03:31"
	endStr = "2025-04-09 20:03:31"

	tstart, errstart = time.Parse(layout, startStr)
	if errstart != nil {
		panic("error")
	}
	tend, errend = time.Parse(layout, endStr)
	if errend != nil {
		panic("error")
	}

	myCalendar = &MyCalendar{int(tstart.Unix()), int(tend.Unix())}
	myCalendar.book(&calendars)
	fmt.Println(*myCalendar)
	fmt.Println(calendars)
}
