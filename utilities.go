package main

import "time"

func TimeWithOffset(offset int, ctime time.Time) time.Time {
	// Add offset to time.Time object
	return ctime.Add(time.Duration(offset) * time.Second)
}
