package chess

import (
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Square struct {
	X, Y            int32
	Piece, OldPiece *chess_pieces.Piece
	Rect            *rl.Rectangle
}
