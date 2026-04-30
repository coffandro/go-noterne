module noterne

go 1.25.9

replace github.com/gen2brain/raylib-go/raylib => ../Raylib-Go-Wasm/raylib

replace client => ../internal/client

require github.com/gen2brain/raylib-go/raylib v0.55.1

require client v0.0.0-00010101000000-000000000000

require github.com/BrownNPC/wasm-ffi-go v1.2.0 // indirect
