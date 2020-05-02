package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/examples/initdisplay"
)

func main() {
	display := initdisplay.InitDisplay()

	str := "0"
	tinyfont.WriteLine(display, &ConstFont, 3, 0x16, str, color.RGBA{0, 0, 0, 255})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
