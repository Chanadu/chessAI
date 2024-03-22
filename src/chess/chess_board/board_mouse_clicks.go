package chess_board

import (
	"github.com/Chanadu/chessAI/src/chess"
	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (b *Board) CheckMouseClicks() {
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			if rl.GetMousePosition().X == 0 || rl.GetMousePosition().Y == 0 {
				continue
			}
			var collision bool = rl.CheckCollisionPointRec(rl.GetMousePosition(), *b.squares[i][j].Rect)

			if collision && rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
				b.squareClicked(i, j)
			}

		}
	}
}

func (b *Board) squareClicked(i, j int32) {
	// println(b.SelectedSquare[0], b.SelectedSquare[1])

	var newSquare *chess.Square = b.squares[i][j]

	if b.selectedSquare[0] == -1 || b.selectedSquare[1] == -1 {
		if newSquare.Piece.Initalized && newSquare.Piece.PieceColor == b.currentTurnColor {
			b.selectedSquare = [2]int32{i, j}
		} else {
			b.selectedSquare = [2]int32{-1, -1}
		}
		return
	}

	if newSquare.Piece.Initalized && newSquare.Piece.PieceColor == b.currentTurnColor {
		b.selectedSquare = [2]int32{i, j}
		return
	}

	var oldSquare *chess.Square = b.squares[b.selectedSquare[0]][b.selectedSquare[1]]

	if oldSquare.X == i && oldSquare.Y == j {
		return
	}

	//EARLY CHECKS DONE
	b.selectedSquare = [2]int32{-1, -1}
	if b.CanPieceMoveTo(oldSquare, newSquare) {
		newSquare.Piece = oldSquare.Piece
		oldSquare.Piece = chess_pieces.NewPiece()
		oldSquare.Piece.Initalized = false
		b.changeTurnColor()
		return
	} else {
		// println("21")
	}
}
