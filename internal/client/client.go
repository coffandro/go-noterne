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

	rl.ClearBackground(rl.DarkGray)

	var top = lay.CutTop(&state.DrawRect, 48)
	rl.DrawRectangleRec(top, rl.Green)

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.START,
		YAlignment: gui.START,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.MIDDLE,
		YAlignment: gui.START,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.END,
		YAlignment: gui.START,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.START,
		YAlignment: gui.MIDDLE,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.MIDDLE,
		YAlignment: gui.MIDDLE,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.END,
		YAlignment: gui.MIDDLE,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.START,
		YAlignment: gui.END,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.MIDDLE,
		YAlignment: gui.END,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	if gui.Button(gui.ButtonParams{
		Bounds:     lay.CutLeft(&top, 48),
		Label:      "Hello!",
		XAlignment: gui.END,
		YAlignment: gui.END,
	}) {
		state.Messages = append(state.Messages, "HI!")
	}

	gui.DrawRect(lay.CutLeft(&state.DrawRect, 48), rl.Black)

	lay.Shrink(&state.DrawRect, 8)

	for i := 0; i < len(state.Messages); i++ {
		rl.DrawText(state.Messages[i], state.DrawRect.ToInt32().X, (int32)(state.DrawRect.Y)+(int32)(i*16), 16, rl.Black)
	}

	rl.EndDrawing()
}
