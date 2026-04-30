package gui

import rl "github.com/gen2brain/raylib-go/raylib"

func CutLeft(rect *rl.Rectangle, n float32) rl.Rectangle {
	var result rl.Rectangle = rl.Rectangle{
		X:      rect.X,
		Y:      rect.Y,
		Width:  n,
		Height: rect.Height,
	}
	rect.X += n
	rect.Width = max(rect.Width-n, 0)
	return result
}

func CutRight(rect *rl.Rectangle, n float32) rl.Rectangle {
	var result rl.Rectangle = rl.Rectangle{
		X:      rect.X + rect.Width - n,
		Y:      rect.Y,
		Width:  n,
		Height: rect.Height,
	}
	rect.Width = max(rect.Width-n, 0)
	return result
}

func CutTop(rect *rl.Rectangle, n float32) rl.Rectangle {
	var result rl.Rectangle = rl.Rectangle{
		X:      rect.X,
		Y:      rect.Y,
		Width:  rect.Width,
		Height: n,
	}
	rect.Y += n
	rect.Height = max(rect.Height-n, 0)
	return result
}

func CutBottom(rect *rl.Rectangle, n float32) rl.Rectangle {
	var result rl.Rectangle = rl.Rectangle{
		X:      rect.X,
		Y:      rect.Y + rect.Height - n,
		Width:  rect.Width,
		Height: n,
	}
	rect.Height = max(rect.Height-n, 0)
	return result
}

func Shrink(rect rl.Rectangle, n float32) rl.Rectangle {
	return rl.Rectangle{
		X:      rect.X + n,
		Y:      rect.Y + n,
		Width:  rect.Width - n*2,
		Height: rect.Height - n*2,
	}
}
