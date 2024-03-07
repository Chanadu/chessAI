package main

import (
	chess "github.com/Chanadu/chessAI/chess"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const ScreenWidth = 800
const ScreenHeight = 800
const SquareSize = min(ScreenWidth, ScreenHeight) / 8

var LightBoardColor = rl.Color{234, 233, 210, 255}
var DarkBoardColor = rl.Color{75, 115, 153, 255}

func main() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "chessAI")
	defer rl.CloseWindow()

	rl.SetTargetFPS(170)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		{
			rl.ClearBackground(rl.RayWhite)
			// rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
			chess.Board(SquareSize, LightBoardColor, DarkBoardColor)
		}
		rl.EndDrawing()
	}
}
