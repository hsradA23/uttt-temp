package game

import (
	"context"
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

func Move(client redis.Client, ctx context.Context, movestr string) bool {

	return true
}
