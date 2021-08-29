package main

import (
	"time"
	"fmt"
	"gauth"
)

var startSecounds = 30
var currentKey = ""

func StartLooping(key string) {
	setNewAccessKey()
	for t = range time.Tick(3*time.Second) {
		repeat(key)
	}
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}

func displayAccessKey() {
	clearConsole()
	fmt.Println("%d -> %s", startSecounds, key)
}

func setNewAccessKey(key string) {
	currentKey = gauth.GenerateKey()
}

func repeatlyGetCode(key string) bool {
	epoch, countDown := gauth.GetCurrentEpoch()

	if(epoch % 30 === 0) {
		setNewAccessKey(key)
		displayAccessKey()
	}else {
		displayAccessKey()
	}
}