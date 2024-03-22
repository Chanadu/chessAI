package chess_board

import (
	"errors"
	"log"

	"github.com/Chanadu/chessAI/src/chess"
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
)

func (b *Board) CanPieceMoveTo(oldSquare, newSquare *chess.Square) bool {
	//println("RUN")
	var oldPiece *chess_pieces.Piece = oldSquare.Piece
	var newPiece *chess_pieces.Piece = newSquare.Piece

	if newPiece.Initalized {
		if oldPiece.PieceColor == newPiece.PieceColor {
			//println("25")
			return false
		}
	}
	var initalXPos, initalYPos, finalXPos, finalYPos int32 = oldSquare.X, oldSquare.Y, newSquare.X, newSquare.Y
	if initalXPos == finalXPos && finalYPos == initalYPos {
		//println("26")
		return false
	}

	switch oldPiece.PieceType {
	case chess_pieces.Pawn:
		//println("PAWN")
		return b.canPawnMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Bishop:
		//println("BISHOP")
		return b.canBishopMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Knight:
		//println("KNIGHT")
		return b.canKnightMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Rook:
		//println("ROOK")
		return b.canRookMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Queen:
		//println("QUEEN")
		return b.canQueenMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.King:
		//println("KING")
		return b.canKingMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
	}
	//println("27")
	log.Fatal(errors.New("CanPieceMoveTo ERROR, NOT IN PIECE LIST"))
	return false
}

func (b *Board) canPawnMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.Pawn {
		log.Fatal(errors.New("canPawnMoveTo ERROR, NOT PAWN"))
		return false
	}
	return true
}

func (b *Board) canBishopMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.Bishop {
		log.Fatal(errors.New("canBishopMoveTo ERROR, NOT BISHOP"))
		return false
	}
	if initalXPos-finalXPos == initalYPos-finalYPos {
		if finalXPos > initalXPos {
			for i := int32(1); i < finalXPos-initalXPos; i++ {
				if initalYPos+i < 7 {
					continue
				}
				if b.squares[initalXPos+i][initalYPos+i].Piece.Initalized {
					return false
				}
			}
		} else {
			for i := int32(1); i < initalXPos-finalXPos; i++ {
				if initalYPos-i < 0 {
					continue
				}
				if b.squares[initalXPos-i][initalYPos-i].Piece.Initalized {
					return false
				}
			}
		}
	} else if -(initalXPos - finalXPos) == initalYPos-finalYPos {
		if finalXPos > initalXPos {
			for i := int32(1); i < finalXPos-initalXPos; i++ {
				if initalYPos-i < 0 {
					continue
				}
				if b.squares[initalXPos+i][initalYPos-i].Piece.Initalized {
					return false
				}
			}
		} else {
			for i := int32(1); i < initalXPos-finalXPos; i++ {
				if initalYPos+i > 7 {
					continue
				}
				if b.squares[initalXPos-i][initalYPos+i].Piece.Initalized {
					return false
				}
			}
		}
	} else {
		return false
	}

	return true
}

func (b *Board) canKnightMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.Knight {
		log.Fatal(errors.New("canKnightMoveTo ERROR, NOT Knight"))
		return false
	}

	return true
}

func (b *Board) canRookMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.Rook {
		log.Fatal(errors.New("canPawnRookTo ERROR, NOT ROOK"))
		return false
	}

	if initalXPos == finalXPos {
		var dx int32
		if initalYPos < finalYPos {
			dx = 1
		} else {
			dx = -1
		}

		for i := initalYPos + dx; i != finalYPos; i += dx {
			if b.squares[initalXPos][i].Piece.Initalized {
				//println("1", initalXPos, finalXPos, j, dx, finalYPos)
				return false
			}
		}
	} else if initalYPos == finalYPos {
		var dy int32
		if initalXPos < finalXPos {
			dy = 1
		} else {
			dy = -1
		}

		for i := initalXPos + dy; i != finalXPos; i += dy {
			if b.squares[i][initalYPos].Piece.Initalized {
				//println("2")
				return false
			}
		}
	} else {
		//println("3")
		return false
	}
	return true
}

func (b *Board) canQueenMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.Queen {
		log.Fatal(errors.New("canQueenMoveTo ERROR, NOT QUEEN"))
		return false
	}
	return b.canRookMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos) || b.canBishopMoveTo(oldSquare, newSquare, initalXPos, initalYPos, finalXPos, finalYPos)
}

func (b *Board) canKingMoveTo(oldSquare, newSquare *chess.Square, initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if oldSquare.Piece.PieceType != chess_pieces.King {
		log.Fatal(errors.New("canKingMoveTo ERROR, NOT KING"))
		return false
	}
	return true
}
