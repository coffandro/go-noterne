package main

import (
	"noterne/internal/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	var clientState = client.Init()

	for !rl.WindowShouldClose() {
		clientState.DrawRect = rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(rl.GetScreenWidth()),
			Height: float32(rl.GetScreenHeight()),
		}

		client.Update(&clientState)
	}
}
