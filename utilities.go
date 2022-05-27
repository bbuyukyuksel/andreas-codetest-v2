package main

import "time"

func TimeWithOffset(offset int, ctime time.Time) time.Time {
	return ctime.Add(time.Duration(offset) * time.Second)
}
