package client

import (
	"fmt"
	"noterne/internal/gui"

	rgui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Init() {

}

func Update(drawRect *rl.Rectangle) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	var top = gui.CutTop(drawRect, 48)
	rl.DrawRectangleRec(top, rl.Green)

	if rgui.Button(gui.CutLeft(&top, 48), "Wow!") {
		fmt.Println("hi!")
	}

	rl.DrawRectangleRec(gui.CutLeft(drawRect, 128), rl.Black)

	rl.EndDrawing()
}
