package main
// This is a test file that test the current state of the Engine
// Nothing here is intended to still be around once the engine is finished (though some parts might be lifted into permanent files later on)

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

func printMove(mv Move) string {
	return square_to_coordinate(mv.from) + square_to_coordinate(mv.to)
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


func self_play() {
	var board State = load_fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	printBoard(board)
	for i := 0; i <20; i++ {

		fmt.Print("\n")
		fmt.Println("thinking...")
		mv, v := bestMove(board, 1)
		fmt.Println(printMove(mv))

		play_move_on(&board, mv)
		printBoard(board)
		fmt.Println(v)
	}
}



func main() {
	/*
	//var a State = load_fen("8/8/5K1k/8/6R1/7P/8/8 w - - 5 85")
	var a State = load_fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	printBoard(a)
	//moves := generate_non_pawn_moves(a, Square{3, 3})
	fmt.Print("\n")
	fmt.Println("thinking...")
	mv, v := bestMove(a, 6)
	fmt.Println(printMove(mv))
	fmt.Println(v)

	 */
	self_play()
}
