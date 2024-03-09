package chess_pieces

type PieceColor int
type PieceType int

const (
	White PieceColor = iota
	Black
)

const (
	King PieceType = iota
	Queen
	Rook
	Bishop
	Knight
	Pawn
)
