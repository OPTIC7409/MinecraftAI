package ui

import (
	"github.com/gioui/gio"
	"github.com/gioui/gio/app"
	"github.com/gioui/gio/layout"
	"github.com/gioui/gio/widget"
	"github.com/gioui/gio/unit"
	"image/color"
)

func DisplaySuggestion(suggestion string) {
	go func() {
		w := app.NewWindow()
		var ops layout.Ops
		msg := widget.Label{
			Text: suggestion,
			TextSize: unit.Dp(16),
			Color: color.White,
		}

		for event := range w.Events() {
			switch e := event.(type) {
			case app.DrawEvent:
				gtx := layout.NewContext(&ops, e)
				layout.Center.Layout(gtx, func(gtx C) D {
					return msg.Layout(gtx)
				})
				w.Invalidate()
			}
		}
	}()
	app.Main()
}
