package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget/material"
)

func run(window *app.Window) error {
	theme := material.NewTheme()
	var ops op.Ops

	for {
		switch e := window.Event().(type) {

		case app.DestroyEvent:
			return e.Err

		case app.FrameEvent:
			// This graphics context is used for managing the rendering state.
			gtx := app.NewContext(&ops, e)

			// Define an large label with an appropriate text:
			title := material.H1(theme, "Hello, Gio")

			// Change the color of the label.
			textColor := color.NRGBA{R: 0, G: 0, B: 0, A: 255}
			title.Color = textColor

			// Change the position of the label.
			title.Alignment = text.Middle

			// Draw the label to the graphics context.
			title.Layout(gtx)

			drawRect(&ops)

			// Pass the drawing operations to the GPU.
			e.Frame(gtx.Ops)
		}

	}
}

func drawRect(ops *op.Ops) {
	defer clip.Rect{Max: image.Pt(100, 100)}.Push(ops).Pop()
	paint.ColorOp{Color: color.NRGBA{R: 0x80, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
