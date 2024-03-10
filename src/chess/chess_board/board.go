package chess_board

import (
	"github.com/Chanadu/chessAI/src/chess"
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	squareSize            int32
	lightColor, darkColor rl.Color
	xPos, yPos            int32
	Squares               [8][8]*chess.Square
	FirstSquareClicked rl.Vector2
}

func (b *Board) ResetBoard() {
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			var p *chess_pieces.Piece = chess_pieces.NewPiece()

			if j <= 1 || j >= 6 {
				p.Initalized = true
			} else {
				p.Initalized = false
			}
			var isPiece bool

			if j == 0 || j == 7 {
				isPiece = true
			} else if j == 1 || j == 6 {
				isPiece = false
			}

			if j <= 1 {
				p.PieceColor = chess_pieces.Black
			}
			if j >= 6 {
				p.PieceColor = chess_pieces.White
			}

			if isPiece {
				switch i {
				case 0, 7:
					p.PieceType = chess_pieces.Rook
				case 1, 6:
					p.PieceType = chess_pieces.Knight
				case 2, 5:
					p.PieceType = chess_pieces.Bishop
				case 3:
					p.PieceType = chess_pieces.Queen
				case 4:
					p.PieceType = chess_pieces.King
				}
			} else {
				p.PieceType = chess_pieces.Pawn
			}
			b.Squares[i][j] = &chess.Square{
				X:     i,
				Y:     j,
				Piece: p,
			}
		}
	}
}

func NewBoard(squareSize int32, lightColor, darkColor rl.Color, xPos, yPos int32) *Board {
	var b *Board = &Board{}
	b.squareSize = squareSize
	b.lightColor = lightColor
	b.darkColor = darkColor
	b.xPos = xPos
	b.yPos = yPos
	b.ResetBoard()
	return b
}
