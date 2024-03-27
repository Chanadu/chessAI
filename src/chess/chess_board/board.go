package chess_board

import (
	"github.com/Chanadu/chessAI/src/chess"
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Board struct {
	squareSize, squareSelectionInset                                                                 int32
	lightColor, darkColor, selectionColor                                                            rl.Color
	xPos, yPos                                                                                       int32
	squares                                                                                          [8][8]*chess.Square
	selectedSquare                                                                                   [2]int32
	currentTurnColor                                                                                 chess_pieces.PieceColor
	canWhiteKingCastleLeft, canWhiteKingCastleRight, canBlackKingCastleLeft, canBlackKingCastleRight bool
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

			var op *chess_pieces.Piece = chess_pieces.NewPiece()
			op.Initalized = false

			b.squares[i][j] = &chess.Square{
				X:        i,
				Y:        j,
				Piece:    p,
				OldPiece: op,
			}
		}
	}
}

func NewBoard(squareSize, squareSelectionInset int32, lightColor, darkColor, selectionColor rl.Color, xPos, yPos int32, turnColor chess_pieces.PieceColor, canKingsCastle bool) *Board {
	var b *Board = &Board{}
	b.squareSize = squareSize
	b.squareSelectionInset = squareSelectionInset
	b.lightColor = lightColor
	b.darkColor = darkColor
	b.selectionColor = selectionColor
	b.yPos = yPos
	b.xPos = xPos
	b.currentTurnColor = turnColor
	b.canWhiteKingCastleLeft, b.canWhiteKingCastleRight, b.canBlackKingCastleLeft, b.canBlackKingCastleRight = canKingsCastle, canKingsCastle, canKingsCastle, canKingsCastle

	b.selectedSquare = [2]int32{-1, -1}
	b.ResetBoard()
	return b
}

func (b *Board) changeTurnColor() {
	if b.currentTurnColor == chess_pieces.White {
		b.currentTurnColor = chess_pieces.Black
	} else {
		b.currentTurnColor = chess_pieces.White
	}
}
