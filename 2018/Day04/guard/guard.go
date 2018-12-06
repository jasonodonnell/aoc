package guard

import (
	"fmt"
	"strings"
	"time"
)

var events []event
var sleepTime map[int]int

type event struct {
	Date    time.Time
	GuardID int
	Action  string
}

type SleepyGuard struct {
	GuardID int
	Slept   int
	Minute  int
}

func AnalyzeShifts(records []string) {
	var lastGuard int
	for _, record := range records {
		event := newEvent(record)
		if event.GuardID == 0 {
			event.GuardID = lastGuard
		} else {
			lastGuard = event.GuardID
		}
		events = append(events, event)
	}
	s := SleepyGuard{}

	s.sleepGuard()
	s.sleepMinute()
	guard, minute := s.frequentSleepyGuardMinute()
	fmt.Printf("Part 1: %d\n", s.GuardID*s.Minute)
	fmt.Printf("Part 2: %d\n", guard*minute)
}

func newEvent(record string) event {
	var event event
	var year, month, day, hour, minute int
	_, _ = fmt.Sscanf(record, "[%d-%d-%d %d:%d]", &year, &month, &day, &hour, &minute)

	event.Date = time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC)

	action := strings.Split(record, "] ")
	_, _ = fmt.Sscanf(action[1], "Guard #%d begins shift", &event.GuardID)

	switch {
	case event.GuardID != 0:
		event.Action = "Begin"
	case action[1] == "falls asleep":
		event.Action = "Asleep"
	case action[1] == "wakes up":
		event.Action = "Awake"
	}

	return event
}

func (s *SleepyGuard) sleepGuard() {
	sleepTime = make(map[int]int)
	var asleepTime int
	for _, event := range events {
		switch {
		case event.Action == "Asleep":
			asleepTime = event.Date.Minute()
		case event.Action == "Awake":
			asleepFor := event.Date.Minute() - asleepTime
			sleepTime[event.GuardID] += asleepFor
		}
	}

	s.Slept = -1
	for guardID, timeSlept := range sleepTime {
		if timeSlept > s.Slept {
			s.GuardID = guardID
			s.Slept = timeSlept
		}
	}
}

func (s *SleepyGuard) sleepMinute() {
	minutes := [60]int{}
	var asleepTime, awakeTime int

	for _, event := range events {
		if event.GuardID != s.GuardID {
			continue
		}
		switch {
		case event.Action == "Asleep":
			asleepTime = event.Date.Minute()
		case event.Action == "Awake":
			awakeTime = event.Date.Minute()
			for i := asleepTime; i < awakeTime; i++ {
				minutes[i]++
				if minutes[i] > minutes[s.Minute] {
					s.Minute = i
				}
			}
		}
	}
}
func (s *SleepyGuard) frequentSleepyGuardMinute() (int, int) {
	minutes := map[int]*[60]int{}
	var asleepTime, awakeTime, sleepyGuard, sleepMinute int

	for _, event := range events {
		switch {
		case event.Action == "Begin":
			if minutes[event.GuardID] == nil {
				minutes[event.GuardID] = &[60]int{}
			}
			if minutes[sleepyGuard] == nil {
				sleepyGuard = event.GuardID
			}
		case event.Action == "Asleep":
			asleepTime = event.Date.Minute()
		case event.Action == "Awake":
			awakeTime = event.Date.Minute()
			for i := asleepTime; i < awakeTime; i++ {
				minutes[event.GuardID][i]++
				if minutes[event.GuardID][i] > minutes[sleepyGuard][sleepMinute] {
					sleepyGuard = event.GuardID
					sleepMinute = i
				}
			}
		}
	}
	return sleepyGuard, sleepMinute
}

// var sleepyguard, sleepyminute int
// minutes := map[int]*[60]int{}
// var guard, from int
// for _, e := range entries {
// 	switch e.action {
// 	case beginShift:
// 		guard = e.guard
// 		if minutes[guard] == nil {
// 			minutes[guard] = &[60]int{}
// 		}
// 		if minutes[sleepyguard] == nil {
// 			sleepyguard = guard
// 		}

// 	case fallAsleep:
// 		from = e.date.Minute()
// 	case wakeUp:
// 		to := e.date.Minute()
// 		for i := from; i < to; i++ {
// 			minutes[guard][i]++
// 			if minutes[guard][i] > minutes[sleepyguard][sleepyminute] {
// 				sleepyguard = guard
// 				sleepyminute = i
// 			}
// 		}
// 	}
// }
