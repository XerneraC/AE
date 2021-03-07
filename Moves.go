// the Æ chess engine
package main

// This file includes the functions involved in generating moves



var NonPawnMoveOffsets = []Square{
	// Rook
	// Queen
	// King
	{+1, +0},
	{+0, +1},
	{-1, -0},
	{-0, -1},

	// Bishop
	// Queen
	// King
	{+1, +1},
	{+1, -1},
	{-1, -1},
	{-1, +1},

	// Knight
	{+2, +1},
	{+2, -1},
	{-2, -1},
	{-2, +1},
	{+1, +2},
	{+1, -2},
	{-1, -2},
	{-1, +2},
}

// TODO: Change these functions, so that they append the moves to a passed slice instead of creating their own (since concatenating 2 slices is expensive)
// TODO: Implement castling
// TODO: Implement En Passant
// TODO: Implement Promotion
// TODO: Always check for checks!!! (only allow moves that would not result in a check)

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
		if !square_legal(sq) { break }
		if state.board[square_index(sq)] != NoPiece { break }
		moves = append(moves, Move{from: from, to: sq})
		if !((from.rank == 1) || (from.rank == 6)) { break }
	}
	for _, sq := range captures {
		if !square_legal(sq) { continue }
		if color_of(state.board[square_index(sq)]) != opposite_color(state.turnToMove) { continue }
		moves = append(moves, Move{from: from, to: sq})
	}
	return moves
}

// TODO: figure out how UCI handles castling moves
func generate_castling_moves(state State) []Move {
	var moves []Move
	castling := state.castlings
	if state.turnToMove == Black { castling = state.castlings >> 2 }
	rank := 0
	if state.turnToMove == Black { rank = 7 }

	kingSquare := Square{file: 4, rank: rank}

	if (castling & WCastleKingside) != 0 {
		firstInTheWay  := state.board[square_index(Square{file: 5, rank: rank})]
		secondInTheWay := state.board[square_index(Square{file: 6, rank: rank})]
		if (firstInTheWay == NoPiece) && (secondInTheWay == NoPiece) {
			newMove := Move{from: kingSquare, to: Square{file: 6, rank: rank}, additionalFlags: CastlingMove}
			moves = append(moves, newMove)
		}
	}

	if (castling & WCastleQueenside) != 0 {
		firstInTheWay  := state.board[square_index(Square{file: 3, rank: rank})]
		secondInTheWay := state.board[square_index(Square{file: 2, rank: rank})]
		thirdInTheWay  := state.board[square_index(Square{file: 1, rank: rank})]
		if (firstInTheWay == NoPiece) && (secondInTheWay == NoPiece) && (thirdInTheWay == NoPiece) {
			newMove := Move{from: kingSquare, to: Square{file: 2, rank: rank}, additionalFlags: CastlingMove}
			moves = append(moves, newMove)
		}
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
	moves = append(moves, generate_castling_moves(state)...)
	return moves
}

