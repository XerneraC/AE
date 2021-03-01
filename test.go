package main
// This is a test file that test the current state of the Engine
// Nothing here is intended to still be around once the engine is finished (though some parts might be lifted into permanent files later on)

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)


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
	for i := 0; i < 20; i++ {

		fmt.Print("\n")
		fmt.Println("thinking...")
		mv, v := bestMove(board, 5)
		fmt.Println(printMove(mv))

		play_move_on(&board, mv)
		printBoard(board)
		fmt.Println(v)
	}
}


func vs_player(human Color) {
	var board State = load_fen("1rbq1knr/pppp1pbp/2n1p1p1/4P3/2BP1B2/2N2N2/PPP1QPPP/2KR3R b Kk - 0 1")
	printBoard(board)
	reader := bufio.NewReader(os.Stdin)
	i := 0
	if human == White { i = 1 }
	for true {
		if i % 2 == 0 {
			fmt.Print("\n")
			fmt.Println("thinking...")
			mv, v := bestMove(board, 5)
			fmt.Println(printMove(mv))

			play_move_on(&board, mv)
			printBoard(board)
			fmt.Println(v)
		} else {
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)
			from := coordinate_to_square(text[:2])
			to := coordinate_to_square(text[2:4])
			mv := Move{from: from, to: to}
			fmt.Println(printMove(mv))

			play_move_on(&board, mv)
			printBoard(board)
		}
		i++
	}
}


func main() {

	//var a State = load_fen("8/8/5K1k/8/6R1/7P/8/8 w - - 5 85")
	//var a State = load_fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	var a State = load_fen("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	printBoard(a)
	//moves := generate_non_pawn_moves(a, Square{3, 3})
	fmt.Println("All:")
	visualize_moves(a, generate_all_possible_moves(a))
	fmt.Print("\n")
	fmt.Println("thinking...")
	start := time.Now()
	mv, v := bestMove(a, 5)
	elapsed := time.Since(start)
	fmt.Println(printMove(mv))
	fmt.Println(v)
	fmt.Printf("Alpha Beta took %s\n", elapsed)
	play_move_on(&a, mv)
	printBoard(a)
	fmt.Println(a)


	//self_play()
	//vs_player(Black)
}
