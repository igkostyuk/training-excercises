//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
// Hint: time.Ticker can be used
// Hint 2: to calculate timediff for Advanced lvl use:
//
//  start := time.Now()
//	// your work
//	t := time.Now()
//	elapsed := t.Sub(start) // 1s or whatever time has passed

package main

import (
	"sync"
	"time"
)

var (
	freemiumTime  = 10 * time.Second
	checkInterval = 1 * time.Second
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mu        sync.Mutex
}

func (u *User) isTimeLimited() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.TimeUsed > int64(freemiumTime/time.Second) && !u.IsPremium
}
func (u *User) increaseTimeUsed(t time.Duration) {
	u.mu.Lock()
	u.TimeUsed += int64(t / time.Second)
	u.mu.Unlock()

}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.isTimeLimited() {
		return false
	}
	start := time.Now()
	t := time.NewTicker(checkInterval)
	defer t.Stop()
	done := make(chan bool)
	go func() {
		process()
		done <- true
	}()
	for {
		select {
		case <-done:
			u.increaseTimeUsed(time.Since(start) / time.Second)
			return true
		case start = <-t.C:
			u.increaseTimeUsed(checkInterval)
			if u.isTimeLimited() {
				return false
			}
		}
	}
}

func main() {
	RunMockServer()
}
