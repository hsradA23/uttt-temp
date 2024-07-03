package game

import (
	"errors"
	"fmt"
	"strings"

	redis_handler "uttt/src/redis-handler"
)

func New_game(name string) {
	// Creates a hash in redis with the key `name`
	// board-{0..8}{0..8} = [0|1|2]
	// board-{0..8} = [0|1|2]
	// turn = [1|2]
	game_data := make(map[string]string)

	for i := 0; i < 9; i++ {
		game_data[fmt.Sprintf("board-%d", i)] = "0"
		for j := 0; j < 9; j++ {
			game_data[fmt.Sprintf("board-%d%d", i, j)] = "0"
		}
	}
	game_data["turn"] = "1"
	game_data["board"] = "-1"
	redis_handler.RedisClient.HSet(redis_handler.Ctx, name, game_data)

	fmt.Println("Created new game.")
}

func Assign_Player(name string, game_id string) error {
	// P1 is X
	// P2 is O
	players, _ := redis_handler.RedisClient.HMGet(redis_handler.Ctx, game_id, "P1", "P2").Result()

	if players[0] == nil || players[0] == name {
		redis_handler.RedisClient.HSet(redis_handler.Ctx, game_id, "P1", name)
		return nil
	} else if players[1] == nil || players[1] == name {
		redis_handler.RedisClient.HSet(redis_handler.Ctx, game_id, "P2", name)
		return nil
	}

	return errors.New("Cannot assign a new player to the game.")
}

func Move(game_id, player_name, movestr string) (string, error) {
	// move board cell
	board, _ := redis_handler.RedisClient.HGetAll(redis_handler.Ctx, game_id).Result()

	var move string
	var next_player string

	if board["P1"] == player_name {
		move = "1"
		next_player = "2"
	} else if board["P2"] == player_name {
		move = "2"
		next_player = "1"
	}

	// Check if the current player is supposed to move
	if board["turn"] != move {
		return "", errors.New("The player is not supposed to move")
	}

	tokens := strings.Split(movestr, " ")
	target_board := tokens[1]
	target_cell := tokens[2]
	cell_id := fmt.Sprintf("board-%s%s", target_board, target_cell)

	// Check if the board is valid
	if board["board"] != "-1" && board["board"] != target_board {
		return "", errors.New("The player is not supposed to play in the board.")
	}

	// Check if cell is already full
	if board[cell_id] != "0" {
		return "", errors.New("Cell already full")
	}

	redis_handler.RedisClient.HSet(redis_handler.Ctx, game_id, cell_id, move, "board", target_cell, "turn", next_player)

	// If everything goes well
	// we return "move board cell P1"
	return fmt.Sprintf("%s P%s", movestr, move), nil
}
