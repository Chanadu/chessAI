package main

import (
	chess "github.com/Chanadu/chessAI/src/chess"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	var ScreenWidth int32 = 800
	var ScreenHeight int32 = 800
	var SquareSize float64 = float64(min(ScreenWidth, ScreenHeight) / 8.0)

	var LightBoardColor = rl.Color{234, 233, 210, 255}
	var DarkBoardColor = rl.Color{75, 115, 153, 255}

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
