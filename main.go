package main

import (
	"fmt"
	"net/http"

	"uttt/src/parser"
	"uttt/src/router"
	"uttt/src/session"
	"uttt/src/sessionManager"

	"github.com/olahol/melody"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/get_game", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		player_name := r.URL.Query().Get("Name")
		game_id := session.Get_Game_By_Name(player_name)
		fmt.Fprintf(w, game_id)
		fmt.Println("Sent response for ", game_id)
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		game_id := r.URL.Query().Get("Game")
		player_name := r.URL.Query().Get("Name")
		if player_name == "" {
			return
		}
		err := router.New_connection(player_name, game_id)
		if err != nil {
			return
		}
		sessionManager.MelodySession.HandleRequest(w, r)
	})

	sessionManager.MelodySession.HandleMessage(func(s *melody.Session, msg []byte) {
		player_name := s.Request.URL.Query().Get("Name")
		resp, err := parser.Handle_Message(player_name, string(msg))
		if err != nil {
			sessionManager.MelodySession.BroadcastFilter([]byte(err.Error()),
				func(sfilter *melody.Session) bool { return sfilter == s })

		} else {
			game_id := session.Get_Game_By_Name(player_name)
			players := session.Get_Players_By_Game(game_id)
			send_to := s.Request.URL.Query().Get("Name")
			sessionManager.MelodySession.BroadcastFilter([]byte(resp),
				func(s *melody.Session) bool { return send_to == players[0] || send_to == players[1] })
		}
	})

	sessionManager.MelodySession.HandleDisconnect(func(s *melody.Session) {
		player_name := s.Request.URL.Query().Get("Name")
		router.Disconnect(player_name)
	})

	fmt.Println("Starting server on port 5000")
	http.ListenAndServe(":5000", nil)
}
