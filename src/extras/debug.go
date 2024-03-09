package extras

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Debug(value string) {
	rl.DrawText(value, 50, 50, 36, rl.Green)
}
