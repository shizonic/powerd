package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/gen2brain/beeep"
)

type Daemon struct {
	batteryInfo *BatteryInfo
	config      *Config
}

func NewDaemon() *Daemon {
	return &Daemon{
		batteryInfo: NewBatteryInfo(),
		config:      NewConfig(),
	}
}

func (d *Daemon) WarnLevelReached() bool {
	if d.batteryInfo.level <= d.config.warnOnLevel {
		return true
	}
	return false
}

func (d *Daemon) ExecLevelReached() bool {
	if d.batteryInfo.level <= d.config.execOnLevel {
		return true
	}
	return false
}

func (d *Daemon) Start() {
	for {
		if d.WarnLevelReached() {
			beeep.Notify("powerd - Warn Level", d.config.warnMessage, "")
		}
		if d.ExecLevelReached() {
			beeep.Alert("powerd - Command Execution", fmt.Sprintf("Executing '%s' in 3 seconds", d.config.execCommand), "")
			time.Sleep(3 * time.Second)
			output, err := exec.Command("sh", "-c", d.config.execCommand).CombinedOutput()
			if err != nil {
				panic(err)
			}
			beeep.Notify("powerd - Command Result", string(output), "")
		}
		time.Sleep(d.config.checkInterval * time.Second)
	}
}
