package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Service interface {
	DaysLeft() int64
	HoursLeft() int64
	MinutesLeft() int64
}

type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s,
	}
}

func (e *Endpoint) Status(c echo.Context) error {

	d := e.s.DaysLeft()
	h := e.s.HoursLeft()
	m := e.s.MinutesLeft()

	s := fmt.Sprintf("До Нового Года осталось %d дней %d часов %d минут!", d, h, m)

	err := c.String(http.StatusOK, s)
	if err != nil {
		return err
	}
	return nil
}
