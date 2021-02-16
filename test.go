package main

import "fmt"




var asld = map[Piece]rune{
	NoPiece: '.',
	WPawn: 'P', WKnight: 'N', WBishop: 'B', WRook: 'R', WQueen: 'Q', WKing: 'K',
	BPawn: 'p', BKnight: 'n', BBishop: 'b', BRook: 'r', BQueen: 'q', BKing: 'k',
}
/*
var asld = map[Piece]rune{
	NoPiece: '.',
	WPawn: '♙', WKnight: '♘', WBishop: '♗', WRook: '♖', WQueen: '♕', WKing: '♔',
	BPawn: '♟', BKnight: '♞', BBishop: '♝', BRook: '♜', BQueen: '♛', BKing: '♚',
}*/

func printBoard(state State) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			fmt.Printf("%c ", asld[state.board[square_index(Square{file, rank})]])
		}
		fmt.Print("\n")
	}
}

func visualize_moves(state State, moves []Move) {
	for rank := 7; rank >= 0; rank-- {
		for file := 0; file < 8; file++ {
			this := Square{file, rank}
			printed := false
			for _, m := range moves {
				if m.to == this {
					fmt.Print("X ")
					printed = true
					break
				}
			}
			if !printed {
				fmt.Printf("%c ", asld[state.board[square_index(this)]])
			}

		}
		fmt.Print("\n")
	}
}




func main() {
	var a State = load_fen("8/8/5K1k/8/6R1/7P/8/8 w - - 5 85")
	//var a State = load_fen("rnbqkbnr/pppppppp/8/8/3R4/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	printBoard(a)
	//moves := generate_non_pawn_moves(a, Square{3, 3})
	moves := generate_all_possible_moves(a)
	fmt.Print("\n")
	visualize_moves(a, moves)
	fmt.Println(moves)
}
