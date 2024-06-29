package session

import (
	"errors"
	"uttt/src/game"
	redis_handler "uttt/src/redis-handler"
)

func Get_Game_By_Name(name string) string {
	game_id, err := redis_handler.RedisClient.HGet(redis_handler.Ctx, "sessions", name).Result()

	if err != nil {
		return ""
	}

	return game_id
}

func Set_Current_Game(name string, game_id string) error {
	players, _ := redis_handler.RedisClient.LRange(redis_handler.Ctx, game_id+"-players", 0, -1).Result()
	if len(players) >= 2 {
		return errors.New("2 people already in the game.")
	}

	for _, p := range players {
		if p == name {
			return errors.New("Player already connected.")
		}
	}

	// Check if the current user is already in a game
	session_game_id, _ := redis_handler.RedisClient.HGet(redis_handler.Ctx, "sessions", name).Result()
	if session_game_id != "" {
		exists, err := redis_handler.RedisClient.Exists(redis_handler.Ctx, session_game_id).Result()
		if exists == 1 {
			game_id = session_game_id
			if err != nil {
				return err
			}
		} else {
			redis_handler.RedisClient.HDel(redis_handler.Ctx, "sessions", name)

		}
		if err != nil {
			return err
		}
		return nil
	}

	err := redis_handler.RedisClient.RPush(redis_handler.Ctx, game_id+"-players", name).Err()
	if err != nil {
		return err
	}

	err = game.Assign_Player(name, game_id)
	if err != nil {
		return err
	}

	return redis_handler.RedisClient.HSet(redis_handler.Ctx, "sessions", name, game_id).Err()
}

func Unset_Current_Game(name string) error {
	game_id := Get_Game_By_Name(name)

	redis_handler.RedisClient.LRem(redis_handler.Ctx, game_id+"-players", 0, name).Result()
	num_players, _ := redis_handler.RedisClient.LLen(redis_handler.Ctx, game_id+"-players").Result()

	if num_players == 0 {
		players, _ := redis_handler.RedisClient.HMGet(redis_handler.Ctx, game_id, "P1", "P2").Result()
		for _, p := range players {
			if p_str, ok := p.(string); ok {
				redis_handler.RedisClient.HDel(redis_handler.Ctx, "sessions", p_str)
			}
		}

		redis_handler.RedisClient.Del(redis_handler.Ctx, game_id)

	}
	return nil
}
