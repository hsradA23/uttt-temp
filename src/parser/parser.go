package parser

import (
	"log"
	"strings"
	"uttt/src/game"
	"uttt/src/session"
)

func Handle_Message(player_name string, msg string) (string, error) {
	game_id := session.Get_Game_By_Name(player_name)
	var resp string = ""
	var err error
	if strings.HasPrefix(msg, "move") {
		resp, err = game.Move(game_id, player_name, msg)
		if err != nil {
			log.Println(err)
			return "", err
		}
	}
	return resp, nil
}
