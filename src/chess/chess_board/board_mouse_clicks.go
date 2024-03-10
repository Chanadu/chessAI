package chess_board

import rl "github.com/gen2brain/raylib-go/raylib"

func (b *Board) CheckMouseClicks() {
	for i := int32(0); i < 8; i++ {
		for j := int32(0); j < 8; j++ {
			if rl.GetMousePosition().X == 0 || rl.GetMousePosition().Y == 0 {
				continue
			}
			var collision bool = rl.CheckCollisionPointRec(rl.GetMousePosition(), *b.Squares[i][j].Rect)
			if collision {
				b.FirstSquareClicked = rl.Vector2{float32(i), float32(j)}
				println(i, " ", j)
			}

		}
	}
}
