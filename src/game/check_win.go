package game

import (
	"fmt"
)

func get_cell(board string, cell int) string {
	// TODO: Find out if this is more expensive than cached strings
	return fmt.Sprintf("board-%s%d", board, cell)
}

func Check_Board_Win(board_id string, board map[string]string) string {
	// Checks if a player has won in a small board
	// Returns "1" or "2" if won
	// otherwise returns empty string ""

	for i := 0; i < 7; i += 3 {
		// Check horizontal win
		if board[get_cell(board_id, i)] == board[get_cell(board_id, i+1)] &&
			board[get_cell(board_id, i+1)] == board[get_cell(board_id, i+2)] && board[get_cell(board_id, i)] != "0" {
			return board[get_cell(board_id, i)]
		}
	}

	for i := 0; i < 3; i++ {
		// Check vertical win
		if board[get_cell(board_id, i)] == board[get_cell(board_id, i+3)] &&
			board[get_cell(board_id, i+3)] == board[get_cell(board_id, i+6)] && board[get_cell(board_id, i)] != "0" {
			return board[get_cell(board_id, i)]
		}
	}
	//First diagonal
	if board[get_cell(board_id, 0)] == board[get_cell(board_id, 4)] &&
		board[get_cell(board_id, 4)] == board[get_cell(board_id, 8)] && board[get_cell(board_id, 0)] != "0" {
		return board[get_cell(board_id, 0)]
	}

	//Second diagonal
	if board[get_cell(board_id, 2)] == board[get_cell(board_id, 4)] &&
		board[get_cell(board_id, 4)] == board[get_cell(board_id, 6)] && board[get_cell(board_id, 2)] != "0" {
		return board[get_cell(board_id, 2)]
	}

	return ""
}

func Check_Win() {
	// Check if a player has won the game
}
