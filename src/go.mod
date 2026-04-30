module noterne

go 1.25.9

require github.com/gen2brain/raylib-go/raylib v0.55.1

require (
	github.com/ebitengine/purego v0.7.1 // indirect
	golang.org/x/exp v0.0.0-20240506185415-9bf2ced13842 // indirect
	golang.org/x/sys v0.20.0 // indirect
)

replace client => ../internal/client

require client v0.0.0-00010101000000-000000000000
