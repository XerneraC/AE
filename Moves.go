package main




var NonPawnMoveOffsets = []Square{
	// Rook
	// Queen
	// King
	{ 1,  0},
	{ 0,  1},
	{-1,  0},
	{ 0, -1},

	// Bishop
	// Queen
	// King
	{ 1,  1},
	{ 1, -1},
	{-1, -1},
	{-1,  1},

	// Knight
	{ 2,  1},
	{ 2, -1},
	{-2, -1},
	{-2,  1},
	{ 1,  2},
	{ 1, -2},
	{-1, -2},
	{-1,  2},
}

// TODO: change these functions, so that they append the moves to a passed slice instead of creating their own (since concatenating 2 slices is expensive)
// TODO: Implement castling
// TODO: Implement En Passant
// TODO: Implement Promotion
// TODO: Always check for checks (only allow moves that would not result in a check)

func generate_non_pawn_moves(state State, from Square) []Move {
	piece := type_of(state.board[square_index(from)])
	var startingOffset int
	var endingOffset int
	var maxDist int
	moves := make([]Move, 0)
	switch piece {
		case Bishop:
			startingOffset = 4
			endingOffset = 8
			maxDist = 8
		case Rook:
			startingOffset = 0
			endingOffset = 4
			maxDist = 8
		case Queen:
			startingOffset = 0
			endingOffset = 8
			maxDist = 8
		case King:
			startingOffset = 0
			endingOffset = 8
			maxDist = 2
		case Knight:
			startingOffset = 8
			endingOffset = 16
			maxDist = 2
		default:
			return moves
	}

	for i := startingOffset; i < endingOffset; i++ {
		offset := NonPawnMoveOffsets[i]
		for dist := 1; dist < maxDist; dist++ {
			target := Square{
				file: from.file + offset.file*dist,
				rank: from.rank + offset.rank*dist,
			}
			if !square_legal(target) { break }
			pieceThatsThere := state.board[square_index(target)]
			if color_of(pieceThatsThere) == state.turnToMove { break }
			if color_of(pieceThatsThere) == opposite_color(state.turnToMove) { dist = 8 }

			move := Move{
				from: from,
				to: target,
			}
			moves = append(moves, move)
		}
	}
	return moves
}



func generate_pawn_moves(state State, from Square) []Move {
	moves := make([]Move, 0)
	if type_of(state.board[square_index(from)]) != Pawn { return moves }
	step := 0
	switch state.turnToMove {
		case White:
			step = 1
		case Black:
			step = -1
	}
	walks := [2]Square{
		{from.file, from.rank + step},
		{from.file, from.rank + step + step},
	}
	captures := [2]Square{
		{from.file + 1, from.rank + step},
		{from.file - 1, from.rank + step},
	}

	for _, sq := range walks {
		if state.board[square_index(sq)] != NoPiece { break }
		moves = append(moves, Move{from, sq})
		if !((from.rank == 1) || (from.rank == 6)) { break }
	}
	for _, sq := range captures {
		if color_of(state.board[square_index(sq)]) != opposite_color(state.turnToMove) { break }
		moves = append(moves, Move{from, sq})
	}
	return moves
}



func generate_moves(state State, from Square) []Move {
	var moves []Move
	moves = append(moves, generate_non_pawn_moves(state, from)...)
	moves = append(moves, generate_pawn_moves(state, from)...)
	return moves
}



// TODO: Improve this function (the way it's written right now is pretty bad. There's alot of improvements that could be made)
func generate_all_possible_moves(state State) []Move {
	moves := make([]Move, 0)
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			from := Square{file, rank}
			if color_of(state.board[square_index(from)]) != state.turnToMove { continue }
			var newMoves []Move
			newMoves = generate_moves(state, from)
			moves = append(moves, newMoves...)
		}
	}
	return moves
}