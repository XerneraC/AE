package main

// This file includes the functions involved in playing the game of Chess


func play_move_on(state *State, move Move) {
	// This function assumes, that the move is legal

	fromIndex := square_index(move.from)
	toIndex := square_index(move.to)

	state.board[toIndex] = state.board[fromIndex]
	state.board[fromIndex] = NoPiece

	// This line would be the entire thing for turncoat:
	// state.board[fromIndex], state.board[toIndex] = state.board[toIndex], state.board[fromIndex]

	state.turnToMove = opposite_color(state.turnToMove)
}

func play_move(state State, move Move) State {
	newState := state
	play_move_on(&newState, move)
	return newState
}