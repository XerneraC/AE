package main

import "fmt"

func get_file_number(letter byte) int {
	return int(letter - 'a')
}

func get_file_letter(number int) byte {
	return byte(number) + 'a'
}

func get_rank_number(letter byte) int {
	return int(letter - '1')
}

func get_rank_letter(number int) byte {
	return byte(number) + '1'
}

func coordinate_to_square_advanced(coord string, index *int) Square {
	var sq Square
	sq.file = get_file_number(coord[*index])
	(*index)++
	sq.rank = get_rank_number(coord[*index])
	(*index)++
	return sq
}

func coordinate_to_square(coord string) Square {
	var index = 0
	return coordinate_to_square_advanced(coord, &index)
}

func square_to_coordinate(sq Square) string {
	return fmt.Sprintf("%c%c", get_file_letter(sq.file), get_rank_letter(sq.rank))
}