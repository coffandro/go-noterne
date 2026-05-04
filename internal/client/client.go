package client

import (
	gui "noterne/internal/gui"
	lay "noterne/internal/lay"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ClientState struct {
	DrawRect rl.Rectangle
	Messages []string
}

func Init() ClientState {
	var state ClientState = ClientState{
		DrawRect: rl.Rectangle{},
		Messages: []string{},
	}

	return state
}

func Update(state *ClientState) {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	var top = lay.CutTop(&state.DrawRect, 48)
	rl.DrawRectangleRec(top, rl.Green)

	if gui.Button(lay.CutLeft(&top, 48), "Wow!") {
		state.Messages = append(state.Messages, "HI!")
	}

	rl.DrawRectangleRec(lay.CutLeft(&state.DrawRect, 48*2), rl.Black)

	lay.Shrink(&state.DrawRect, 8)

	for i := 0; i < len(state.Messages); i++ {
		rl.DrawText(state.Messages[i], state.DrawRect.ToInt32().X, (int32)(state.DrawRect.Y)+(int32)(i*16), 16, rl.Black)
	}

	rl.EndDrawing()
}
