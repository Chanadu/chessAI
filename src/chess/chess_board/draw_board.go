package chess_board

import (
	"os"

	"github.com/Chanadu/chessAI/src/chess/chess_pieces"
	"github.com/Chanadu/chessAI/src/extras"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var int32ToLetters [8]string = [8]string{"a", "b", "c", "d", "e", "f", "g", "h"}

func (b *Board) DrawBoard() {
	b.drawBoardTiles()
	b.drawBoardMarkings()
}

func (b *Board) drawBoardTiles() {
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			var color rl.Color
			if (i+j)%2 == 0 {
				color = b.lightColor
			} else {
				color = b.darkColor
			}

			var squareRect rl.Rectangle = rl.NewRectangle(
				float32(i*b.squareSize+b.xPos),
				float32(j*b.squareSize+b.yPos),
				float32(b.squareSize),
				float32(b.squareSize),
			)

			b.squares[i][j].Rect = &squareRect

			rl.DrawRectangleRec(squareRect, color)

			if b.selectedSquare[0] == i && b.selectedSquare[1] == j {
				var selectedSquareRect1 rl.Rectangle = rl.NewRectangle(
					float32(i*b.squareSize+b.xPos),
					float32(j*b.squareSize+b.yPos),
					float32(b.squareSize),
					float32(b.squareSize),
				)
				var selectedSquareRect2 rl.Rectangle = rl.NewRectangle(
					float32(i*b.squareSize+b.xPos+b.squareSelectionInset),
					float32(j*b.squareSize+b.yPos+b.squareSelectionInset),
					float32(b.squareSize-b.squareSelectionInset*2),
					float32(b.squareSize-b.squareSelectionInset*2),
				)
				rl.DrawRectangleRec(selectedSquareRect1, b.selectionColor)
				rl.DrawRectangleRec(selectedSquareRect2, color)
			}
			b.drawBoardPiece(i, j)
		}
	}
}

func (b *Board) drawBoardPiece(i, j int32) {
	if !b.squares[i][j].Piece.Initalized {
		return
	}
	// else {
	// rl.DrawText("K", i*b.squareSize+b.xPos, j*b.squareSize+b.yPos, 36, rl.Black)
	// return
	// }
	var index int32 = -1
	if b.squares[i][j].Piece.PieceColor == chess_pieces.Black {
		index += 6
	}
	switch t := b.squares[i][j].Piece.PieceType; t {
	case chess_pieces.King:
		index += 1
	case chess_pieces.Queen:
		index += 2
	case chess_pieces.Rook:
		index += 3
	case chess_pieces.Bishop:
		index += 4
	case chess_pieces.Knight:
		index += 5
	case chess_pieces.Pawn:
		index += 6
	}
	// rl.DrawText(text, i*b.squareSize+b.xPos+b.squareSize/4, j*b.squareSize+b.yPos+b.squareSize/4, 36, rl.Black)
	rl.DrawTexture(*PieceTextures[index], i*b.squareSize, j*b.squareSize-b.squareSize/64, rl.White)
}

func (b *Board) drawBoardMarkings() {
	for i := int32(0); i < 8; i++ {
		var color1 rl.Color
		var color2 rl.Color
		if i%2 == 0 {
			color1 = b.lightColor
			color2 = b.darkColor
		} else {
			color1 = b.darkColor
			color2 = b.lightColor
		}
		rl.DrawText(extras.FormatInt32ToString(8-i), 8+b.xPos, i*b.squareSize+8+b.yPos, 24, color2)
		rl.DrawText(int32ToLetters[i], (i+1)*b.squareSize-24+b.xPos, b.squareSize*8-24+b.yPos, 24, color1)
	}
}

func (b *Board) LoadPieceImages() {
	b.loadPieceImage(0, "/assets/wk.png")
	b.loadPieceImage(1, "/assets/wq.png")
	b.loadPieceImage(2, "/assets/wr.png")
	b.loadPieceImage(3, "/assets/wb.png")
	b.loadPieceImage(4, "/assets/wn.png")
	b.loadPieceImage(5, "/assets/wp.png")
	b.loadPieceImage(6, "/assets/bk.png")
	b.loadPieceImage(7, "/assets/bq.png")
	b.loadPieceImage(8, "/assets/br.png")
	b.loadPieceImage(9, "/assets/bb.png")
	b.loadPieceImage(10, "/assets/bn.png")
	b.loadPieceImage(11, "/assets/bp.png")

}

func (b *Board) loadPieceImage(index int32, filePath string) {
	path, _ := os.Getwd()
	var tempImage *rl.Image = rl.LoadImage(path + filePath)
	rl.ImageResize(tempImage, b.squareSize, b.squareSize)
	var pieceTextureBB rl.Texture2D = rl.LoadTextureFromImage(tempImage)
	rl.UnloadImage(tempImage)
	PieceTextures[index] = &pieceTextureBB
}
