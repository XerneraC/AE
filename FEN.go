package main
// This file includes all the functions involved in loading and generating FEN strings


var fenPieces = map[byte]Piece{
	'P': WPawn, 'N': WKnight, 'B': WBishop, 'R': WRook, 'Q': WQueen, 'K': WKing,
	'p': BPawn, 'n': BKnight, 'b': BBishop, 'r': BRook, 'q': BQueen, 'k': BKing,
}

var fenColors = map[byte]Color{
	'w': White,
	'b': Black,
}


func load_fen_advanced(fen string, index *int) State {
	var loaded State

	// Loading the board
	for rank := 7; rank >= 0; rank-- {
		file := 0
		for file < 8 {
			letter := fen[*index]
			(*index)++
			if letter > '0' && letter <= '8' {
				file += int(letter - '0')
			} else {
				piece := fenPieces[letter]
				loaded.board[square_index(Square{file, rank})] = piece
				file++
			}
		}
		(*index)++
	}

	// Loading who's turn it is
	loaded.turnToMove = fenColors[fen[*index]]
	(*index)++
	(*index)++



	return loaded
}


func load_fen(fen string) State {
	index := 0
	return load_fen_advanced(fen, &index)
}
