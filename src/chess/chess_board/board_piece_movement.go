package chess_board

import (
	"errors"
	"log"
	"math"

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
	// println(oldPiece.PieceType)
	switch oldPiece.PieceType {
	case chess_pieces.Pawn:
		//println("PAWN")
		return b.canPawnMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Bishop:
		//println("BISHOP")
		return b.canBishopMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, false)
	case chess_pieces.Knight:
		//println("KNIGHT")
		return b.canKnightMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Rook:
		//println("ROOK")
		return b.canRookMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, false)
	case chess_pieces.Queen:
		//println("QUEEN")
		return b.canQueenMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.King:
		//println("KING")
		return b.canKingMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	}
	//println("27")
	log.Fatal(errors.New("CanPieceMoveTo ERROR, NOT IN PIECE LIST"))
	return false
}

func (b *Board) canPieceSee(initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	var oldPiece *chess_pieces.Piece = b.squares[initalXPos][initalYPos].Piece
	// println(oldPiece.PieceType)
	switch oldPiece.PieceType {
	case chess_pieces.Pawn:
		// println("PAWN")
		return b.canPawnMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Bishop:
		// println("BISHOP")
		return b.canBishopMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, false)
	case chess_pieces.Knight:
		// println("KNIGHT")
		return b.canKnightMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.Rook:
		// println("ROOK")
		return b.canRookMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, false)
	case chess_pieces.Queen:
		// println("QUEEN")
		return b.canQueenMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	case chess_pieces.King:
		// println("KING")
		return b.canKingMoveTo(initalXPos, initalYPos, finalXPos, finalYPos)
	}
	//println("27")
	log.Fatal(errors.New("CanPieceMoveTo ERROR, NOT IN PIECE LIST"))
	return false
}

func (b *Board) canPawnMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.Pawn {
		log.Fatal(errors.New("canPawnMoveTo ERROR, NOT PAWN"))
		return false
	}
	if initalXPos == finalXPos {
		if b.squares[finalXPos][finalYPos].Piece.Initalized {
			return false
		}
		if b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.White {
			if initalYPos == 6 && finalYPos == 4 {
				return true
			} else if finalYPos+1 == initalYPos {
				return true
			}
		} else {
			if initalYPos == 1 && finalYPos == 3 {
				return true
			} else if finalYPos-1 == initalYPos {
				return true
			}
		}
	} else if math.Abs(float64(initalXPos-finalXPos)) == 1 && math.Abs(float64(finalYPos-initalYPos)) == 1 {
		//Captures
		if b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.White {
			if finalYPos-initalYPos == -1 && b.squares[finalXPos][finalYPos].Piece.Initalized {
				return true
			}
		}
		if b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.Black {
			if finalYPos-initalYPos == 1 && b.squares[finalXPos][finalYPos].Piece.Initalized {
				return true
			}
		}

		//En Passant
		if initalYPos == 3 &&
			b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.White &&
			b.squares[finalXPos][initalYPos].Piece.PieceType == chess_pieces.Pawn &&
			b.squares[finalXPos][initalYPos].Piece.PieceColor != b.squares[initalXPos][initalYPos].Piece.PieceColor &&
			b.squares[finalXPos][1].OldPiece.PieceType == chess_pieces.Pawn {
			b.squares[finalXPos][initalYPos].Piece = chess_pieces.NewPiece()
			b.squares[finalXPos][initalYPos].Piece.Initalized = false
			return true
		}
		if initalYPos == 4 &&
			b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.Black &&
			b.squares[finalXPos][initalYPos].Piece.PieceType == chess_pieces.Pawn &&
			b.squares[finalXPos][initalYPos].Piece.PieceColor != b.squares[initalXPos][initalYPos].Piece.PieceColor &&
			b.squares[finalXPos][6].OldPiece.PieceType == chess_pieces.Pawn {

			b.squares[finalXPos][initalYPos].Piece = chess_pieces.NewPiece()
			b.squares[finalXPos][initalYPos].Piece.Initalized = false
			return true
		}
	}

	return false
}

func (b *Board) canBishopMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32, isQueenMove bool) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.Bishop && !isQueenMove {
		log.Fatal(errors.New("canBishopMoveTo ERROR, NOT BISHOP"))
		return false
	}
	if initalXPos-finalXPos == initalYPos-finalYPos {
		if finalXPos > initalXPos {
			for i := int32(1); i < finalXPos-initalXPos; i++ {
				if initalYPos+i > 7 {
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

func (b *Board) canKnightMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.Knight {
		log.Fatal(errors.New("canKnightMoveTo ERROR, NOT Knight"))
		return false
	}

	if math.Abs(float64(finalXPos-initalXPos)) == 1 {
		return math.Abs(float64(finalYPos-initalYPos)) == 2
	}
	if math.Abs(float64(finalXPos-initalXPos)) == 2 {
		return math.Abs(float64(finalYPos-initalYPos)) == 1
	} else {
		return false
	}
}

func (b *Board) canRookMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32, isQueenMove bool) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.Rook && !isQueenMove {
		log.Fatal(errors.New("canRookMoveTo ERROR, NOT ROOK"))
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
				return false
			}
		}
	} else {
		return false
	}

	if initalXPos == 0 && initalYPos == 0 && b.canBlackKingCastleLeft {
		b.canBlackKingCastleLeft = false
	} else if initalXPos == 7 && initalYPos == 0 && b.canBlackKingCastleRight {
		b.canBlackKingCastleRight = false
	} else if initalXPos == 0 && initalYPos == 7 && b.canWhiteKingCastleLeft {
		b.canWhiteKingCastleLeft = false
	} else if initalXPos == 7 && initalYPos == 7 && b.canWhiteKingCastleRight {
		b.canWhiteKingCastleRight = false
	}

	return true
}

func (b *Board) canQueenMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.Queen {
		log.Fatal(errors.New("canQueenMoveTo ERROR, NOT QUEEN"))
		return false
	}
	return b.canRookMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, true) || b.canBishopMoveTo(initalXPos, initalYPos, finalXPos, finalYPos, true)
}

func (b *Board) canKingMoveTo(initalXPos, initalYPos, finalXPos, finalYPos int32) bool {
	if b.squares[initalXPos][initalYPos].Piece.PieceType != chess_pieces.King {
		log.Fatal(errors.New("canKingMoveTo ERROR, NOT KING"))
		return false
	}

	if math.Abs(float64(finalXPos-initalXPos)) <= 1 && math.Abs(float64(finalYPos-initalYPos)) <= 1 {
		for i := int32(-1); i <= 1; i++ {
			for j := int32(-1); j <= 1; j++ {
				var xPos, yPos int32 = finalXPos + i, finalYPos + j
				if xPos > 7 || yPos > 7 || xPos < 0 || yPos < 0 || (xPos == initalXPos && yPos == initalXPos) {
					continue
				}
				if b.squares[xPos][yPos].Piece.Initalized && b.squares[xPos][yPos].Piece.PieceColor != b.squares[initalXPos][initalYPos].Piece.PieceColor && b.squares[xPos][yPos].Piece.PieceType == chess_pieces.King {
					// println(1)
					return false
				}
			}
		}
		if b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.White {
			b.canWhiteKingCastleLeft = false
			b.canWhiteKingCastleRight = false
		} else {
			b.canBlackKingCastleLeft = false
			b.canBlackKingCastleRight = false
		}
		return true
	} else {
		// println(2)
		// CASTLING
		if initalYPos == 7 && b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.White {
			if (finalXPos-initalXPos == 2) &&
				b.canWhiteKingCastleRight &&
				!b.squares[initalXPos+1][7].Piece.Initalized &&
				!b.squares[initalXPos+2][7].Piece.Initalized {
				//Castle Right (Move Rook)
				b.forceMovePiece(b.squares[7][7], b.squares[5][7])
				return true
			} else if (initalXPos-finalXPos == 3) &&
				b.canWhiteKingCastleLeft &&
				!b.squares[initalXPos-1][7].Piece.Initalized &&
				!b.squares[initalXPos-2][7].Piece.Initalized &&
				!b.squares[initalXPos-3][7].Piece.Initalized {
				//Castle Left (Move Rook)
				b.forceMovePiece(b.squares[0][7], b.squares[3][7])
				return true
			}
		} else if initalYPos == 0 && b.squares[initalXPos][initalYPos].Piece.PieceColor == chess_pieces.Black {
			if (finalXPos-initalXPos == 2) && b.canBlackKingCastleRight &&
				!b.squares[initalXPos+1][0].Piece.Initalized &&
				!b.squares[initalXPos+2][0].Piece.Initalized {
				//Castle Right (Move Rook)
				b.forceMovePiece(b.squares[7][0], b.squares[5][0])
				return true
			} else if (initalXPos-finalXPos == 3) && b.canBlackKingCastleLeft &&
				!b.squares[initalXPos-1][0].Piece.Initalized &&
				!b.squares[initalXPos-2][0].Piece.Initalized &&
				!b.squares[initalXPos-3][0].Piece.Initalized {
				//Castle Left (Move Rook)
				b.forceMovePiece(b.squares[0][0], b.squares[3][0])
				return true
			}
		}
		return false
	}
}

func (b *Board) movePiece(oldSquare, newSquare *chess.Square) bool {

	if b.CanPieceMoveTo(oldSquare, newSquare) {
		b.forceMovePiece(oldSquare, newSquare)
		if b.currentTurnColor == chess_pieces.White &&
			b.isKingInCheck(b.findPiecePosition(chess_pieces.King, chess_pieces.White)[0], chess_pieces.White) {

			println("WHITE CHECK")
			b.forceMovePiece(newSquare, oldSquare)
			return false
		}
		if b.currentTurnColor == chess_pieces.Black &&
			b.isKingInCheck(b.findPiecePosition(chess_pieces.King, chess_pieces.Black)[0], chess_pieces.Black) {
			b.forceMovePiece(newSquare, oldSquare)
			println("BLACK CHECK")
			return false
		}
		b.changeTurnColor()
		return true
	}
	return false
}

func (b *Board) findPiecePosition(pType chess_pieces.PieceType, pColor chess_pieces.PieceColor) [][]int32 {
	var foundPositions [][]int32 = make([][]int32, 0)
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			if b.squares[i][j].Piece.Initalized && b.squares[i][j].Piece.PieceType == pType && b.squares[i][j].Piece.PieceColor == pColor {
				foundPositions = append(foundPositions, []int32{i, j})
			}
		}
	}
	return foundPositions
}

func (b *Board) forceMovePiece(oldSquare, newSquare *chess.Square) {
	oldSquare.OldPiece = chess_pieces.NewPiece()
	oldSquare.OldPiece.Initalized = true
	oldSquare.OldPiece.PieceType = oldSquare.Piece.PieceType
	oldSquare.OldPiece.PieceColor = oldSquare.Piece.PieceColor
	newSquare.Piece = oldSquare.Piece
	oldSquare.Piece = chess_pieces.NewPiece()

	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			if i == oldSquare.X && j == oldSquare.Y {
				continue
			}
			if b.squares[i][j].OldPiece.PieceColor == oldSquare.OldPiece.PieceColor {
				b.squares[i][j].OldPiece = chess_pieces.NewPiece()
				b.squares[i][j].OldPiece.Initalized = false
			}
		}
	}
	oldSquare.Piece.Initalized = false
}

func (b *Board) isKingInCheck(kingPos []int32, color chess_pieces.PieceColor) bool {
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			if i == kingPos[0] && j == kingPos[1] {
				continue
			}
			if b.squares[i][j].Piece.Initalized && b.squares[i][j].Piece.PieceColor != color {
				if b.canPieceSee(i, j, kingPos[0], kingPos[1]) {
					println(i, j, kingPos[0], kingPos[1])
					println("!!!!!!!!!!!!!")
					return true
				}
			}
		}
	}
	return false
}
