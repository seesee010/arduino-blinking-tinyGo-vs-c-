// blink an LED, golang version
package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for true {
		led.High()
		time.Sleep(2)

		led.Low()
		time.Sleep(2)
	}
}
