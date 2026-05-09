package gui

// Mostly stolen from https://github.com/raysan5/raygui/
// Currently we dont support styling...
// We assume all no borders, backgrounds are raywhite and text is left 16 sized

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GuiState int64

const (
	STATE_NORMAL GuiState = iota
	STATE_FOCUSED
	STATE_PRESSED
	STATE_DISABLED
)

type Alignment int64

const (
	MIDDLE Alignment = iota
	START
	END
)

var (
	BG = rl.White
	FG = rl.Black
)

var guiState GuiState = STATE_NORMAL

func PosRectColliding(pos rl.Vector2, rect rl.Rectangle) bool {
	return pos.X >= rect.X && pos.X < rect.X+rect.Width &&
		pos.Y >= rect.Y && pos.Y < rect.Y+rect.Height
}

func GetTextWidth(text string, textHeight float32) int32 {
	if text == "" {
		return 0
	}

	// Strip optional icon prefix: "#123#Text"
	if len(text) > 0 && text[0] == '#' {
		for i := 1; i < 5 && i < len(text); i++ {
			if text[i] == '#' {
				text = text[i+1:]
				break
			}
		}
	}

	// Stop at first newline
	runes := make([]rune, 0, len(text))
	for _, r := range text {
		if r == '\n' {
			break
		}
		runes = append(runes, r)
	}

	// Assume monospace font:
	// width ≈ character count * character width
	//
	// Typical monospace aspect ratio is ~0.6
	charWidth := textHeight * 0.6

	return int32(float32(len(runes)) * charWidth)
}

func DrawAlignedLabelRect(
	label string,
	bounds rl.Rectangle,
	XAlignment Alignment,
	YAlignment Alignment,
	FontSize int32,
) {
	var yPos float32 = 0

	if YAlignment == START {
		yPos = bounds.Y
	} else if YAlignment == MIDDLE {
		yPos = bounds.Y + bounds.Height/2 - float32(FontSize)/2
	} else if YAlignment == END {
		yPos = bounds.Y + bounds.Height - float32(FontSize)
	}

	var textSize int32 = GetTextWidth(label, float32(FontSize))

	if XAlignment == START {
		rl.DrawText(label, bounds.ToInt32().X, int32(yPos), FontSize, FG)
	} else if XAlignment == MIDDLE {
		rl.DrawText(label, bounds.ToInt32().X+bounds.ToInt32().Width/2-textSize/2, int32(yPos), FontSize, FG)
	} else if XAlignment == END {
		rl.DrawText(label, bounds.ToInt32().X+bounds.ToInt32().Width-textSize, int32(yPos), FontSize, FG)
	}
}

func DrawRect(bounds rl.Rectangle, color ...rl.Color) {
	var col rl.Color = BG

	if len(color) > 0 {
		col = color[0]
	}

	rl.DrawRectangleRec(bounds, col)
}

type ButtonParams struct {
	XAlignment Alignment
	YAlignment Alignment
	Bounds     rl.Rectangle
	Label      string
	FontSize   int32
}

func EnsureButtonParams(params *ButtonParams) {
	params.XAlignment = params.XAlignment
	params.YAlignment = params.YAlignment
	params.Bounds = params.Bounds
	params.Label = params.Label

	if params.FontSize == 0 {
		params.FontSize = 8
	}
}

func Button(params ButtonParams) bool {
	EnsureButtonParams(&params)

	var result bool = false
	var state GuiState = guiState

	if state != STATE_DISABLED {
		var mousePoint rl.Vector2 = rl.GetMousePosition()

		if PosRectColliding(mousePoint, params.Bounds) {
			if rl.IsMouseButtonDown(rl.MouseButtonLeft) {
				state = STATE_PRESSED
			} else {
				state = STATE_FOCUSED
			}

			if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
				result = true
			}
		}
	}

	DrawRect(params.Bounds)
	DrawAlignedLabelRect(params.Label, params.Bounds, params.XAlignment, params.YAlignment, params.FontSize)

	return result
}
