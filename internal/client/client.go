package client

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func init() {

}

func Update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)
	rl.DrawText("Congrats! You created your first window!", 10, 10, 20, rl.LightGray)

	rl.EndDrawing()
}
