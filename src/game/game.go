package game

import (
	"context"
	"errors"
	"fmt"

	redis_handler "uttt/src/redis-handler"

	"github.com/redis/go-redis/v9"
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

func Move(client redis.Client, ctx context.Context, movestr string) bool {

	return true
}
