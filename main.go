package main

import (
	"net/http"

	"uttt/src/melody_session"
	"uttt/src/parser"
	"uttt/src/router"

	"github.com/olahol/melody"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		player_name := r.Header.Get("Name")
		game_id := r.URL.Query().Get("Game")
		if player_name == "" {
			return
		}
		err := router.New_connection(player_name, game_id)
		if err != nil {
			return
		}
		melodysession.MelodySession.HandleRequest(w, r)
	})

	melodysession.MelodySession.HandleMessage(func(s *melody.Session, msg []byte) {
		player_name := s.Request.Header.Get("Name")
		parser.Handle_Message(player_name, string(msg))
		melodysession.MelodySession.Broadcast(msg)
	})

	melodysession.MelodySession.HandleDisconnect(func(s *melody.Session) {
		player_name := s.Request.Header.Get("Name")
		router.Disconnect(player_name)
	})

	http.ListenAndServe(":5000", nil)
}
