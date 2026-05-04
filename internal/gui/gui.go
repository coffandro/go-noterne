package gui

// Mostly stolen from https://github.com/raysan5/raygui/
// Currently we dont support styling...
// We assume all no borders, backgrounds are raywhite and text is left 16 sized

import rl "github.com/gen2brain/raylib-go/raylib"

type GuiState int64

const (
	STATE_NORMAL GuiState = iota
	STATE_FOCUSED
	STATE_PRESSED
	STATE_DISABLED
)

var guiState GuiState = STATE_NORMAL

func PosRectColliding(pos rl.Vector2, rect rl.Rectangle) bool {
	return pos.X >= rect.X && pos.X < rect.X+rect.Width &&
		pos.Y >= rect.Y && rect.Y < rect.Y+rect.Height
}

func Button(bounds rl.Rectangle, label string) bool {
	var result bool = false
	var state GuiState = guiState

	if state != STATE_DISABLED {
		var mousePoint rl.Vector2 = rl.GetMousePosition()

		if PosRectColliding(mousePoint, bounds) {
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

	rl.DrawRectangleRec(bounds, rl.RayWhite)
	rl.DrawText(label, bounds.ToInt32().X, bounds.ToInt32().Y, 16, rl.Black)

	return result
}
