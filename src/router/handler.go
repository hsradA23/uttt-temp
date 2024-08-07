package router

import (
	"fmt"
	"log"
	"math/rand/v2"

	"uttt/src/game"
	"uttt/src/session"
)

func New_connection(player_name string, game_id string) error {
	if game_id == "" {
		game_id = session.Get_Game_By_Name(player_name)
	}
	if game_id == "" {
		game_id = fmt.Sprintf("%d", 1000+rand.IntN(8999))
		game.New_game(game_id)
		err := session.Set_Current_Game(player_name, game_id)
		if err != nil {
			log.Println(err.Error())
		} else {
			log.Printf("[INFO] New game for user=%s created\n", player_name)
		}

	} else {
		err := session.Set_Current_Game(player_name, game_id)
		if err != nil {
			return err
		}
		log.Printf("User already connected to %s\n", game_id)
	}
	return nil
}

func Disconnect(player_name string) {
	session.Unset_Current_Game(player_name)
}
