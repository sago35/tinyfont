package main

import (
	"fmt"
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/drivers/ili9341"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/examples/initdisplay"
)

func main() {
	displayer := initdisplay.InitDisplay()
	display, ok := displayer.(*ili9341.Device)
	if !ok {
		for {
			tinyfont.WriteLine(display, &MplusConst10pt, 3, 15, "ili9341 device only", color.RGBA{0, 0, 0, 255})
			time.Sleep(time.Hour)
		}
	}

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	i := 0
	tinyfont.Wrap = true

	display.FillScreen(color.RGBA{255, 255, 255, 255})
	tinyfont.WriteLine(display, &MplusConst10pt, 3, 15, fmt.Sprintf("%s", s[i].UserName), color.RGBA{0, 0, 0, 255})
	tinyfont.WriteLine(display, &MplusConst10pt, tinyfont.Cx, tinyfont.Cy, fmt.Sprintf(" @%s\n", s[i].ScreenName), color.RGBA{128, 128, 128, 255})

	tinyfont.WriteLine(display, &MplusConst10pt, 3, tinyfont.Cy, fmt.Sprintf("%s fav:%d rt:%d\n", s[i].CreatedAt, s[i].FavoriteCount, s[i].RetweetCount), color.RGBA{128, 128, 128, 255})
	tinyfont.WriteLine(display, &MplusConst10pt, 3, tinyfont.Cy+5, s[i].FullText, color.RGBA{0, 0, 0, 255})

	for {

		led.Toggle()
		tinyfont.WriteLine(display, &MplusConst10pt, 3, 315, fmt.Sprintf("%d / %d", i+1, len(s)), color.RGBA{64, 64, 64, 255})

	}
}
