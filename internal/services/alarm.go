package services

import (
	"fmt"
	"time"
)

const (
	StateRinging = iota + 1
	StateSnoozing
	StateDisabled
	StateWaiting
)

type AlarmTime struct {
	NextRing time.Time
	Ring     time.Time
}

type AlarmService struct {
	Alarms []*AlarmTime
}

var Alarm = NewAlarmService()

func NewAlarmService() *AlarmService {
	return &AlarmService{
		Alarms: make([]*AlarmTime, 5),
	}

}

func (as *AlarmService) NewAlarm(ringtime time.Time) *AlarmTime {
	a := &AlarmTime{}
	as.Alarms = append(as.Alarms, a)
	now := time.Now()

	rt := time.Date(now.Year(), now.Month(), now.Day(), ringtime.Hour(), ringtime.Minute(), 0, 0, time.Local)
	if rt.Before(now) {
		rt = rt.Add(time.Hour * 24)
	}

	a.Ring = rt
	a.NextRing = rt

	return a
}

func (as *AlarmService) NextAlarm() (*AlarmTime, error) {
	now := time.Now()
	var next *AlarmTime
	var least time.Duration

	for _, a := range as.Alarms {
		if a == nil {
			continue
		}

		ring := a.Ring
		for now.After(ring) {
			ring = ring.Add(time.Hour * 24)
		}

		if d := ring.Sub(now); d < least || least.Seconds() == 0 {
			least = d
			next = a
		}
	}

	if next == nil {
		return nil, fmt.Errorf("no alarm in the future found")
	}

	return next, nil
}
