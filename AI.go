// the Æ chess engine
package main
// This file includes all the functions that make up the Engine itself

import (
	"math"
)

//import "fmt"

/*
var pieceValue = map[PieceType]float32{
	Pawn: 1, Knight: 3, Bishop: 3, Rook: 5, Queen: 9, King: 0, // The King has a value of 0, since kings are handled in the evaluate function separately
}*/

// My tuned Piece Values. Mainly to give the Engine a personality and to make it feel a bit different than other Engines
// My thinking is, that my Engine will not be a good Engine anyway, so might aswell sacrifice some skill to make it more interesting.
var pieceValue = map[PieceType]float64{
	Pawn: 0.8, Knight: 3, Bishop: 3.1, Rook: 5, Queen: 9, King: 0, // The King has a value of 0, since kings are handled in the evaluate function separately
}



// Piece-Square Tables Incoming
// These are copy-pasted off the internet and don't really represent the ones that I'd want to use
// They're not used right now and were only really intended to debug the alpha beta pruning
// I left them here, because I have a creeping suspicion, that this'll come in handy later on
var pieceSquares = map[PieceType][64]float64{
	Pawn: {
		0,  0,  0,  0,  0,  0,  0,  0,
		50, 50, 50, 50, 50, 50, 50, 50,
		10, 10, 20, 30, 30, 20, 10, 10,
		5,  5, 10, 27, 27, 10,  5,  5,
		0,  0,  0, 25, 25,  0,  0,  0,
		5, -5,-10,  0,  0,-10, -5,  5,
		5, 10, 10,-25,-25, 10, 10,  5,
		0,  0,  0,  0,  0,  0,  0,  0,
	},
	Knight: {
		-50,-40,-30,-30,-30,-30,-40,-50,
		-40,-20,  0,  0,  0,  0,-20,-40,
		-30,  0, 10, 15, 15, 10,  0,-30,
		-30,  5, 15, 20, 20, 15,  5,-30,
		-30,  0, 15, 20, 20, 15,  0,-30,
		-30,  5, 10, 15, 15, 10,  5,-30,
		-40,-20,  0,  5,  5,  0,-20,-40,
		-50,-40,-20,-30,-30,-20,-40,-50,
	},
	Bishop: {
		-20,-10,-10,-10,-10,-10,-10,-20,
		-10,  0,  0,  0,  0,  0,  0,-10,
		-10,  0,  5, 10, 10,  5,  0,-10,
		-10,  5,  5, 10, 10,  5,  5,-10,
		-10,  0, 10, 10, 10, 10,  0,-10,
		-10, 10, 10, 10, 10, 10, 10,-10,
		-10,  5,  0,  0,  0,  0,  5,-10,
		-20,-10,-40,-10,-10,-40,-10,-20,
	},
	King: {
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-30, -40, -40, -50, -50, -40, -40, -30,
		-20, -30, -30, -40, -40, -30, -30, -20,
		-10, -20, -20, -20, -20, -20, -20, -10,
		20,  20,   0,   0,   0,   0,  20,  20,
		20,  30,  10,   0,   0,  10,  30,  20,
	},
	Queen: {
		-20,-10,-10,-10,-10,-10,-10,-20,
		-10,  0,  0,  0,  0,  0,  0,-10,
		-10,  0,  5, 10, 10,  5,  0,-10,
		-10,  5,  5, 10, 10,  5,  5,-10,
		-10,  0, 10, 10, 10, 10,  0,-10,
		-10, 10, 10, 10, 10, 10, 10,-10,
		-10,  5,  0,  0,  0,  0,  5,-10,
		-20,-10,-40,-10,-10,-40,-10,-20,
	},
}

var placeMultiplier = [64]float64{
	0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8,
	0.8, 0.7, 0.7, 0.7, 0.7, 0.7, 0.7, 0.8,
	0.8, 0.7, 0.9, 0.9, 0.9, 0.9, 0.7, 0.8,
	0.8, 0.7, 0.9, 1.0, 1.0, 0.9, 0.7, 0.8,
	0.8, 0.7, 0.9, 1.0, 1.0, 0.9, 0.7, 0.8,
	0.8, 0.7, 0.9, 0.9, 0.9, 0.9, 0.7, 0.8,
	0.8, 0.7, 0.7, 0.7, 0.7, 0.7, 0.7, 0.8,
	0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8, 0.8,
}


func isOver(node State) bool {
	whiteKingPresent := false
	blackKingPresent := false
	for i := 0; i < 64; i++ {
		switch node.board[i] {
			case WKing:
				whiteKingPresent = true
				break
			case BKing:
				blackKingPresent = false
				break
		}
	}
	return !(whiteKingPresent || blackKingPresent)
}


// TODO: Completely redo this function.
// I made it originally, to debug a bug I encountered (which took me 2h and the bug was not even in the part of the code that i was debugging fml).
// The way it's written right now is complete dogshit. I'm not even using snake case for the function name. I just didn't care when I wrote this function
func bestMove(state State, depth int) (Move, float64) {
	maximizingPlayer := state.turnToMove == White
	var best Move
	var bestValue = math.Inf(+1)
	if maximizingPlayer { bestValue = math.Inf(-1) }

	for _, move := range generate_all_possible_moves(state) {
		newState := play_move(state, move)
		value := α_β_pruning(newState, depth - 1, math.Inf(-1), math.Inf(+1))
		condition := value <= bestValue
		if maximizingPlayer {
			condition = value >= bestValue
		}
		if condition {
			best = move
			bestValue = value
		}
	}
	return best, bestValue
}

func get_child_nodes(node State) []State {
	states := make([]State, 0)
	for _, move := range generate_all_possible_moves(node) {
		states = append(states, play_move(node, move))
	}
	return states
}

// Dear Go developers
//
// I love you!!!
// You support identifiers that include non ascii symbols.
// I love you!!!
//
// Best wishes
// Xern
type cache_entry struct {
	exploredDepth int
	value         float64
}
var cache = make(map[State]cache_entry)
func α_β_pruning(node State, depth int, α float64, β float64) float64 {
	if depth == 0 { return evaluate(node) }
	if isOver(node) { return evaluate(node) }

	// Check cache
	cacheEntry := cache[node]
	if cacheEntry.exploredDepth >= depth { return cacheEntry.value }

	maximizingPlayer := node.turnToMove == White
	value := math.Inf(+1)
	if maximizingPlayer {
		value = math.Inf(-1)
	}
	for _, child := range get_child_nodes(node) {
		// There should be a way to improve this branch.
		// I could pull it out of the for loop, which would probably improve performance, but would make the code more repetitive
		if maximizingPlayer {
			value = math.Max(value, α_β_pruning(child, depth - 1, α, β))
			α = math.Max(α, value)
			if α >= β { break }
		} else {
			value = math.Min(value, α_β_pruning(child, depth - 1, α, β))
			β = math.Min(β, value)
			if β <= α { break }
		}
	}

	// Update cache
	cache[node] = cache_entry{exploredDepth: depth, value: value}

	return value
}

// This function is not really used. It's mainly here to test if my alpha beta pruning implementation works
// I'm keeping it here just in case as a fallback
func minmax(node State, depth int) float64 {
	maximizingPlayer := node.turnToMove == White
	if depth == 0 {
		return evaluate(node)
	}
	sign := +1
	if maximizingPlayer {
		sign = -1
	}
	value := math.Inf(sign)
	for _, child := range get_child_nodes(node) {
		res := minmax(child, depth-1)

		if maximizingPlayer {
			if res == math.Inf(+1) { return res }
			value = math.Max(value, res)
		} else {
			if res == math.Inf(-1) { return res }
			value = math.Min(value, res)
		}
	}
	return value
}

func evaluate(state State) float64 {
	var value float64 = 0
	whiteHasKing := false
	blackHasKing := false
	for i := 0; i < 64; i++ {
		piece := state.board[i]
		pieceType, pieceColor := seperate(piece)
		switch pieceColor {
			case White:
				whiteHasKing = whiteHasKing || (pieceType == King)
				value += pieceValue[pieceType]
			case Black:
				blackHasKing = blackHasKing || (pieceType == King)
				value -= pieceValue[pieceType]
			default:
				continue
		}
	}
	if whiteHasKing && !blackHasKing { return math.Inf(+1) } else if !whiteHasKing && blackHasKing { return math.Inf(-1) }
	return value
}


// Evaluates the board using the piece-squares tables from above.
// Not really used right now and only here since it might come in handy to keep it for later
func evaluate_PieceSquares(state State) float64 {
	var value float64 = 0
	whiteHasKing := false
	blackHasKing := false
	for i := 0; i < 64; i++ {
		piece := state.board[i]
		pieceType, pieceColor := seperate(piece)
		switch pieceColor {
		case White:
			whiteHasKing = whiteHasKing || (pieceType == King)
 			// This doesn't work really good since it doesn't mirror the maps for Black
			value += pieceSquares[pieceType][i] * placeMultiplier[i]
		case Black:
			blackHasKing = blackHasKing || (pieceType == King)
  			// This doesn't work really good since it doesn't mirror the maps for Black
			value -= pieceSquares[pieceType][i] * placeMultiplier[i]
		default:
			continue
		}
	}
	if whiteHasKing && !blackHasKing {
		return math.Inf(+1)
	} else if !whiteHasKing && blackHasKing {
		return math.Inf(-1)
	}
	return value
}

