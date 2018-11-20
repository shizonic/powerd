package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	BatteryStatusPath   = "/sys/class/power_supply/BAT0/status"
	BatteryCapacityPath = "/sys/class/power_supply/BAT0/capacity"
	AcOnlinePath        = "/sys/class/power_supply/AC/online"
)

type BatteryInfo struct {
	onAC   bool
	status string
	level  int
}

func NewBatteryInfo() *BatteryInfo {
	s := &BatteryInfo{
		onAC:   false,
		status: "Unknown",
		level:  0,
	}
	s.Update()
	return s
}

func (b *BatteryInfo) Update() {
	bytesArray, err := ioutil.ReadFile(BatteryStatusPath)
	if err != nil {
		panic(err)
	}
	b.status = string(bytesArray)
	if _, err := os.Stat(AcOnlinePath); !os.IsNotExist(err) {
		b.onAC = true
	}
	bytesArray, err = ioutil.ReadFile(BatteryCapacityPath)
	if err != nil {
		panic(err)
	}
	level, err := strconv.Atoi(strings.TrimSuffix(string(bytesArray), "\n"))
	if err != nil {
		panic(err)
	}
	b.level = level
}
