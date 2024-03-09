package chess_pieces

type Piece struct {
	Initalized bool
	PieceType  PieceType
	PieceColor PieceColor
}

func NewPiece() *Piece {
	var p *Piece = &Piece{}
	return p
}
