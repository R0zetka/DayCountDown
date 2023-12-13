package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (e *Service) DaysLeft() int64 {
	d := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)

	dur := time.Until(d)

	return int64(dur.Hours()) / 24
}

func (e *Service) HoursLeft() int64 {
	d := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)

	dur := time.Until(d)

	return int64(dur.Hours()) % 24
}

func (e *Service) MinutesLeft() int64 {
	d := time.Date(time.Now().Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)

	dur := time.Until(d)

	return int64(dur.Minutes()) % 60
}
