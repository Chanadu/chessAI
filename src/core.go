package main

import (
	"github.com/Chanadu/chessAI/src/chess/chess_board"
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const screenWidth int32 = 800
const screenHeight int32 = 800

const squareSize int32 = min(screenWidth, screenHeight) / 8
const squareSelectionInset int32 = squareSize / 10

var lightBoardColor rl.Color = rl.Color{234, 233, 210, 255}
var darkBoardColor rl.Color = rl.Color{75, 115, 153, 255}
var selectionBoardColor rl.Color = rl.Color{226, 245, 56, 255}

var board *chess_board.Board

func RaylibCreateWindow() {
	rl.InitWindow(screenWidth, screenHeight, "ChessAI")
	rl.SetTargetFPS(170)
}

func RaylibWindowShouldClose() bool {
	return rl.WindowShouldClose()
}

func RaylibCloseWindow() {
	rl.CloseWindow()
}

func PreLoop() {
	board = chess_board.NewBoard(squareSize, squareSelectionInset, lightBoardColor, darkBoardColor, selectionBoardColor, 0, 0, chess_pieces.White, true)
	board.LoadPieceImages()
}

func MainGameLoop() {
	rl.BeginDrawing()
	{
		rl.ClearBackground(rl.RayWhite)
		board.DrawBoard()
		board.CheckMouseClicks()
	}
	rl.EndDrawing()
}
