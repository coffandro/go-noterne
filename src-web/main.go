package main

import (
	"fmt"
	"syscall/js"

	"noterne/internal/client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func set_dimensions(this js.Value, args []js.Value) any {
	rl.SetWindowSize(args[0].Int(), args[1].Int())
	return map[string]any{"message": nil}
}

func main() {
	// Expose function to JS to allow for resizing
	js.Global().Set("goSetDimensions", js.FuncOf(set_dimensions))

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	rl.SetWindowState(rl.FlagWindowResizable)
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)

	fmt.Println("Starting client")

	client.Init()

	// If a global JS function named "onRuntimeInitialized" exists, call it.
	if fn := js.Global().Get("onRuntimeInitialized"); fn.Type() == js.TypeFunction {
		fn.Invoke()
	} else {
		fmt.Println("onRuntimeInitialized not found on global object")
	}

	var update = func() {
		var drawRect rl.Rectangle = rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(rl.GetScreenWidth()),
			Height: float32(rl.GetScreenHeight()),
		}

		client.Update(&drawRect)
	}

	rl.SetMainLoop(update)
}
