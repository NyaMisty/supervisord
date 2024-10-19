// +build !windows

package main

import (
	"github.com/ramr/go-reaper"
)

// ReapZombie reap the zombie child process
func ReapZombie() {
	go reaper.Start(reaper.Config{
		Pid:              -1,
		Options:          0,
		DisablePid1Check: true,
		Debug:            false,
	})
}
