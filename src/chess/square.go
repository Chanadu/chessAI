package chess

import "github.com/Chanadu/chessAI/src/chess/chess_pieces"

type Square struct {
	X, Y  int32
	Piece *chess_pieces.Piece
}
