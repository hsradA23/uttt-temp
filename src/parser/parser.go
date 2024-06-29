package parser

import (
	"fmt"
	"strings"
	"uttt/src/game"
	"uttt/src/session"
)

func Handle_Message(player_name string, msg string) {
	game_id := session.Get_Game_By_Name(player_name)
	if strings.HasPrefix(msg, "move") {
		err := game.Move(game_id, player_name, msg)
		if err != nil {
			fmt.Println(err)
		}
	}
}
