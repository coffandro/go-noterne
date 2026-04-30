package main

import (
	"fmt"
	"syscall/js"

	"client"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func set_dimensions(this js.Value, args []js.Value) any {
	fmt.Println("Changing dimensions")

	rl.SetWindowSize(args[0].Int(), args[1].Int())
	return map[string]any{"message": nil}
}

func main() {
	js.Global().Set("goSetDimensions", js.FuncOf(set_dimensions))

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	fmt.Println("Starting client")

	// If a global JS function named "onRuntimeInitialized" exists, call it.
	if fn := js.Global().Get("onRuntimeInitialized"); fn.Type() == js.TypeFunction {
		fn.Invoke()
	} else {
		fmt.Println("onRuntimeInitialized not found on global object")
	}

	rl.SetMainLoop(client.Update)
}
