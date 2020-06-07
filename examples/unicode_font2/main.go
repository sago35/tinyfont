package main

import (
	"image/color"
	"machine"
	"time"

	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/examples/initdisplay"
	"tinygo.org/x/tinyfont/notosans"
)

func main() {
	display := initdisplay.InitDisplay()

	str := "https://tour.golang.org/welcome/2\n" +
		"\n" +
		"Brazilian Portuguese\n    — Português do Brasil\n" +
		"Catalan — Català\n" +
		"Simplified Chinese — 中文（简体）\n" +
		"Traditional Chinese — 中文（繁體）\n" +
		"Czech — Česky\n" +
		"French — Français\n" +
		"German — Deutsch\n" +
		"Hebrew — עִבְרִית\n" +
		"Indonesian — Bahasa Indonesia\n" +
		"Italian — Italiano\n" +
		"Japanese — 日本語\n" +
		"Korean — 한국어\n" +
		"Romanian — Română\n" +
		"Russian - Русский\n" +
		"Spanish — Español\n" +
		"Thai - ภาษาไทย\n" +
		"Turkish - Türkçe\n" +
		"Ukrainian — Українська\n" +
		"Uzbek — Ўзбекча\n"

	tinyfont.WriteLine(display, &notosans.Notosans12pt, 3, 0x16, str, color.RGBA{0, 0, 0, 255})

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 500)

		led.High()
		time.Sleep(time.Millisecond * 500)
	}
}
