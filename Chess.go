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



	////////////////////////////
	// Castling related stuff //
	////////////////////////////

	// All of these could be done using language ternaries
	// IF GO HAD ANY RRRR I HATE THAT GO DOESN'T HAVE TERNARIES THEY'RE SO USEFUL
	// RIGHT NOW I HAVE TO TRUST THE GO COMPILER TO USE BRANCHLESS CODE WHERE IT IMPROVES IT
	// I DON'T THINK GO USES BRANCHLESS CODE AT ALL GRRRRRR I HATE IT MY CODE IS SO SLOW GRRRRR
	kingCastling := WCastleKingside
	queenCastling := WCastleQueenside
	rank := 0
	if state.turnToMove == Black {
		kingCastling  <<= 2
		queenCastling <<= 2
		rank = 7
	}

	// This will only be executed if castlings are still available
	if state.castlings & (kingCastling | queenCastling) != 0 {
		if type_of(state.board[toIndex]) == King {
			state.castlings = state.castlings & ^(kingCastling | queenCastling)
		}

		if move.from == (Square{file: 0, rank: rank}) {
			state.castlings = state.castlings & ^queenCastling
		}
		if move.from == (Square{file: 7, rank: rank})  {
			state.castlings = state.castlings & ^kingCastling
		}


		// This will only be executed when the move is a castling move
		// Notice how it's not checked whether or not the player can still castle that rook
		// Here we just trust that the move is valid, like we do in the entire function
		if isMoveCastling(move) {
			var towerFrom int
			var towerTo int
			if move.to.file > 4 {
				towerFrom = square_index(Square{file: 7, rank: move.to.rank})
				towerTo = square_index(Square{file: 5, rank: move.to.rank})
			} else {
				towerFrom = square_index(Square{file: 0, rank: move.to.rank})
				towerTo = square_index(Square{file: 3, rank: move.to.rank})

			}
			state.board[towerTo] = state.board[towerFrom]
			state.board[towerFrom] = NoPiece
		}
	}

	state.turnToMove = opposite_color(state.turnToMove)
}

func play_move(state State, move Move) State {
	newState := state
	play_move_on(&newState, move)
	return newState
}