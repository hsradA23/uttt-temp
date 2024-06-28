package parser

import (
	"fmt"
	"strings"
	"uttt/src/session"
)

func Handle_Message(player_name string, msg string) {
	_ = session.Get_Game_By_Name(player_name)
	if strings.HasPrefix(msg, "move") {
		tokens := strings.Split(msg, " ")
		if len(tokens) < 3 {
			return
		}
		move(tokens[1], tokens[2])
	}

}

func move(board string, cell string) {
	fmt.Printf("%s: %s", board, cell)
}
