package session

import (
	"errors"
	"fmt"
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
			err := redis_handler.RedisClient.RPush(redis_handler.Ctx, session_game_id+"-players", name).Err()
			if err != nil {
				return err
			}

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

	return redis_handler.RedisClient.HSet(redis_handler.Ctx, "sessions", name, game_id).Err()
}

func Unset_Current_Game(name string) error {
	game_id := Get_Game_By_Name(name)

	redis_handler.RedisClient.LRem(redis_handler.Ctx, game_id+"-players", 0, name).Result()
	num_players, _ := redis_handler.RedisClient.LLen(redis_handler.Ctx, game_id+"-players").Result()

	if num_players == 0 {
		redis_handler.RedisClient.Del(redis_handler.Ctx, game_id)
		redis_handler.RedisClient.HDel(redis_handler.Ctx, "sessions", name)
	}
	fmt.Println("User removed.")
	return nil
}
