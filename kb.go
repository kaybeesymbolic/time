package main

import (
	"fmt"
	"time"
)

//TimeInterface is interface to be implemented
type TimeInterface interface {
	Clock() (hr, min, sec int)
	Date() (year int, month time.Month, day int)
	TimeElapsed() time.Duration
}

//KbTime struct is used to wrap time pointer
type KbTime struct {
	Time time.Time
}

//New returns pointer of KbTime
func New() *KbTime {
	time := NewTime().(*KbTime)
	return time
}

//NewTime is used to return object of new timer interface
func NewTime() TimeInterface {
	return &KbTime{
		Time: time.Now().UTC(),
	}
}

//Clock can be called on receiver, KbTime, to get the current clock.
//Current clock includes hour, minute and second elapsed.
//This method returns three values which can be used independently.
func (kb *KbTime) Clock() (hr, min, sec int) {
	hr, min, sec = kb.Time.Clock()
	return
}

//Date returns current date when calls on the receiver
//Calling to date returns three values which can be processed independently
func (kb *KbTime) Date() (year int, month time.Month, day int) {
	year, month, day = kb.Time.Date()
	return
}

//TimeElapsed returns time passed since particular event
func (kb *KbTime) TimeElapsed() time.Duration {
	duration := time.Since(kb.Time)
	return duration
}

func main() {
	kb := New()
	fmt.Println(kb.Clock())
	fmt.Println(kb.Date())
}
